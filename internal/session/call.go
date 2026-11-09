package session

import (
	"context"

	"github.com/andrew-hayworth22/wodify-go"
)

// Call runs fn with the session's client and returns errors as *Error.
// A 401 response clears the session and fires the OnUnauthorized callback.
func Call[T any](ctx context.Context, s *Session, fn func(context.Context, *wodify.Client) (T, error)) (T, error) {
	var zero T

	wc := s.current()
	if wc == nil {
		return zero, &Error{Type: ErrUnauthorized, Message: defaultMessage(ErrUnauthorized)}
	}

	res, err := fn(ctx, wc)
	if err != nil {
		e := toError(err)
		if e.Type == ErrUnauthorized {
			s.expire(wc)
		}
		return zero, e
	}
	return res, nil
}
