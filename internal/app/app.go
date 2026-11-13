package app

import (
	"context"
	"fmt"
	"time"

	"github.com/andrew-hayworth22/wodify-wails/internal/config"
	"github.com/andrew-hayworth22/wodify-wails/internal/session"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// windowShowFallback is how long Startup waits before showing the window in event
// that the frontend fails to show the window.
const windowShowFallback = 3 * time.Second

// EventAuthExpired is emitted to the frontend when an API call is rejected as unauthorized and the session is cleared.
// Keep in sync with AUTH_EXPIRED_EVENT in frontend/src/lib/auth.ts.
const EventAuthExpired = "auth:expired"

// App struct
type App struct {
	ctx     context.Context
	config  *config.Config
	session *session.Session
}

// New creates a new App application struct
func New() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	time.AfterFunc(windowShowFallback, func() { runtime.WindowShow(ctx) })

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("reading config: %v\n", err)
		return
	}
	a.config = cfg

	s, err := session.New(session.Options{
		Client: session.ClientOptions{
			BaseURL: cfg.BaseURL,
			Timeout: cfg.Timeout,
		},
		OnUnauthorized: func() {
			runtime.EventsEmit(ctx, EventAuthExpired)
		},
	})
	if err != nil {
		fmt.Printf("starting API session: %v\n", err)
		return
	}
	a.session = s
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string, err bool) (string, error) {
	if err {
		return "", fmt.Errorf("error")
	}
	return fmt.Sprintf("Hello %s, It's show time!", name), nil
}
