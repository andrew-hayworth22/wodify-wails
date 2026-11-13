package app

import (
	"errors"

	"github.com/andrew-hayworth22/wodify-wails/internal/config"
	"github.com/andrew-hayworth22/wodify-wails/internal/session"
)

// GetConfig returns the application configuration
func (a *App) GetConfig() config.Config {
	if a.config == nil {
		panic("config not loaded: startup must be called before accessing config")
	}
	return *a.config
}

// SetConfig sets the application configuration.
func (a *App) SetConfig(cfg config.Config) error {
	if a.session != nil {
		err := a.session.SetClientOptions(session.ClientOptions{
			BaseURL: cfg.BaseURL,
			Timeout: cfg.Timeout,
		})
		if err != nil {
			return err
		}
	}
	a.config = &cfg
	return nil
}

// SaveConfig saves the application configuration to disk
func (a *App) SaveConfig() error {
	return a.config.Write()
}

// HasAPIToken reports whether an API token is configured
func (a *App) HasAPIToken() bool {
	return a.session != nil && a.session.HasToken()
}

// SetAPIToken validates the token against the Wodify API, then saves it
// to the OS keyring and uses it for later API calls
func (a *App) SetAPIToken(token string) error {
	if a.session == nil {
		return errors.New("API session not initialized")
	}
	return a.session.SetToken(a.ctx, token)
}

// SignOut clears the API token from memory and the OS keyring
func (a *App) SignOut() error {
	if a.session == nil {
		return errors.New("API session not initialized")
	}
	return a.session.Clear()
}
