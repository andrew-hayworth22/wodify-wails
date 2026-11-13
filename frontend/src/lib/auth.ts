// Dependency key for the auth guard.
export const AUTH_DEP = 'app:auth';

// Event emitted when the API token is rejected and cleared.
// Keep in sync with EventAuthExpired in internal/app/app.go.
export const AUTH_EXPIRED_EVENT = 'auth:expired';