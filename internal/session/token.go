package session

import (
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
)

const (
	keyringService = "wodify"
	keyringUser    = "session-token"
)

func getToken() (string, error) {
	tok, err := keyring.Get(keyringService, keyringUser)
	if errors.Is(err, keyring.ErrNotFound) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("reading token from keyring: %w", err)
	}
	return tok, err
}

func setToken(token string) error {
	if err := keyring.Set(keyringService, keyringUser, token); err != nil {
		return fmt.Errorf("writing token to keyring: %w", err)
	}
	return nil
}

func deleteToken() error {
	// A missing entry is already the desired state.
	err := keyring.Delete(keyringService, keyringUser)
	if err != nil && !errors.Is(err, keyring.ErrNotFound) {
		return fmt.Errorf("deleting token from keyring: %w", err)
	}
	return nil
}
