package err

import (
	"errors"
	"net/http"
)

type ErrParam struct {
	Err      error
	Msg      string
	Status   int
	LogLevel string
}

type AppErr struct {
	E map[string]*ErrParam
}

var Err = &AppErr{E: map[string]*ErrParam{}}

// Add adds a new error to the global registry
// Add a new error locally (instance method)
func (e *AppErr) Add(key, err,
	msg string, status int, logLevel ...string) *ErrParam {
	var level string
	if len(logLevel) > 0 {
		level = logLevel[0]
	}
	param := &ErrParam{
		Err:      errors.New(err),
		Msg:      msg,
		Status:   status,
		LogLevel: level,
	}
	e.E[key] = param
	return param
}

// global Add (shortcut)
func Add(key, err, msg string, status int, logLevel ...string) *ErrParam {
	return Err.Add(key, err, msg, status, logLevel...)
}

// Get retrieves an error by key
func (e *AppErr) Get(key string) *ErrParam {
	return e.E[key]
}

// All predefined errors
func Init() {
	Add("BlogNotFound",     "blog not found", "blog not found", http.StatusNotFound)
	Add("DbUnavailable",    "database unavailable", "service unavailable", http.StatusServiceUnavailable, "err")

	Add("EmailExists",      "email already exists", "email already exists", http.StatusConflict)
	Add("UserNameExists",   "username already exists", "username already exists", http.StatusConflict)

	Add("InvalidCredentials",   "invalid credentials", "invalid credentials", http.StatusUnauthorized)
	Add("InvalidUID",           "invalid UID", "invalid UID", http.StatusBadRequest)
	Add("InvalidUser",          "invalid user", "invalid user", http.StatusBadRequest)
	Add("InvalidUserName",      "invalid username", "invalid username", http.StatusBadRequest)
	Add("InvalidSlug",          "invalid slug", "invalid slug", http.StatusBadRequest)

	Add("RefreshTokenNotFound", "refresh token not found", "refresh token not found", http.StatusUnauthorized)
	Add("RefreshTokenRevoked",  "refresh token revoked", "refresh token revoked", http.StatusUnauthorized)
	Add("RefreshTokenExpired",  "refresh token expired", "refresh token expired", http.StatusUnauthorized)
}
