// Package session represents a session with the Wodify API.
package session

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/andrew-hayworth22/wodify-go"
	"github.com/andrew-hayworth22/wodify-go/utils"
)

// ClientOptions configures the underlying Wodify HTTP client.
type ClientOptions struct {
	BaseURL string
	Timeout time.Duration
}

// Options configures the Wodify session.
type Options struct {
	Client         ClientOptions
	OnUnauthorized func()
}

// Session represents a session with the Wodify API.
type Session struct {
	mu             sync.RWMutex
	opts           ClientOptions
	token          string
	wc             *wodify.Client
	onUnauthorized func()
}

// New builds a new Wodify session using the given options.
func New(opts Options) (*Session, error) {
	s := &Session{
		opts:           opts.Client,
		onUnauthorized: opts.OnUnauthorized,
	}

	tok, err := getToken()
	if err != nil {
		return nil, err
	}
	if tok != "" {
		wc, err := newClient(tok, opts.Client)
		if err != nil {
			return nil, err
		}
		s.token, s.wc = tok, wc
	}
	return s, nil
}

// HasToken returns true if the session has a valid token.
func (s *Session) HasToken() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.token != ""
}

func (s *Session) SetToken(ctx context.Context, token string) error {
	token = strings.TrimSpace(token)
	if token == "" {
		return &Error{Type: ErrValidation, Message: "API key is required."}
	}

	s.mu.RLock()
	opts := s.opts
	s.mu.RUnlock()

	wc, err := newClient(token, opts)
	if err != nil {
		return toError(err)
	}
	if err := validateClient(ctx, wc); err != nil {
		e := toError(err)
		if e.Type == ErrUnauthorized || e.Type == ErrForbidden {
			e.Message = "That API key was not accepted."
		}
		return e
	}

	if err := setToken(token); err != nil {
		return toError(err)
	}

	s.mu.Lock()
	s.token, s.wc = token, wc
	s.mu.Unlock()
	return nil
}

// SetClientOptions updates the client options and re-creates the Wodify API client with the new options.
func (s *Session) SetClientOptions(opts ClientOptions) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.opts = opts
	if s.token == "" {
		return nil
	}

	wc, err := newClient(s.token, opts)
	if err != nil {
		return toError(err)
	}
	s.wc = wc
	return nil
}

// Clear removes the token from memory and the keyring.
func (s *Session) Clear() error {
	s.mu.Lock()
	s.token, s.wc = "", nil
	s.mu.Unlock()
	return deleteToken()
}

// current returns the active Wodify API client.
func (s *Session) current() *wodify.Client {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.wc
}

// expire clears session after an unauthorized response if the client is the same.
func (s *Session) expire(wc *wodify.Client) {
	if s.wc != wc {
		return
	}

	_ = s.Clear()
	if s.onUnauthorized != nil {
		s.onUnauthorized()
	}
}

// newClient creates a Wodify API client using the given token and config
func newClient(token string, opts ClientOptions) (*wodify.Client, error) {
	o := []wodify.Option{
		wodify.WithAPIKey(token),
		wodify.WithBaseURL(opts.BaseURL),
	}
	if opts.Timeout > 0 {
		o = append(o, wodify.WithTimeout(opts.Timeout))
	}
	return wodify.New(o...)
}

// validateClient validates that the Wodify token is working properly with a cheap read-only call to the API
func validateClient(ctx context.Context, wc *wodify.Client) error {
	_, err := wc.Utils.ListGenders(ctx, utils.NewGenderListRequest(
		wodify.NewPaginationRequest(1, 1),
		wodify.SortAscending(utils.GenderFieldID),
	))
	return err
}
