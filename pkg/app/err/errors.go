package err

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type ErrParam struct {
	Err      error
	Msg      string
	HttpStatus   int
	LogLevel string
}

type AppErr struct {
	E map[string]*ErrParam
}

var Err = &AppErr{
	E: map[string]*ErrParam{},
}

// Add adds a new error to the global registry
// Add a new error locally (instance method)
func (e *AppErr) Add(key, err, msg string, HttpStatus int, logLevel ...string) *ErrParam {
	var level string
	if len(logLevel) > 0 {
		level = logLevel[0]
	}
	param := &ErrParam{
		Err:      errors.New(err),
		Msg:      msg,
		HttpStatus:   HttpStatus,
		LogLevel: level,
	}
	// store err param to 
	e.E[err] = param
	return param
}

// global Add (shortcut)
func Add(key, err, msg string, HttpStatus int, logLevel ...string) *ErrParam {
	return Err.Add(key, err, msg, HttpStatus, logLevel...)
}

// Get retrieves an error by key
func (e *AppErr) Get(key string) *ErrParam { return e.E[key] }

// All predefined errors
var (
	EmailExists    = Add("EmailExists", "email already exists", "email already exists", http.StatusConflict)
	UserNameExists = Add("UserNameExists", "username already exists", "username already exists", http.StatusConflict)

	InvalidCredentials = Add("InvalidCredentials", "invalid credentials", "invalid credentials", 401)
	InvalidUID         = Add("InvalidUID", "invalid UID", "invalid UID", http.StatusBadRequest)
	InvalidUser        = Add("InvalidUser", "invalid user", "invalid user", http.StatusBadRequest)
	InvalidUserName    = Add("InvalidUserName", "invalid username", "invalid username", http.StatusBadRequest)
	InvalidSlug        = Add("InvalidSlug", "invalid slug", "invalid slug", http.StatusBadRequest)

	RefreshTokenNotFound = Add("RefreshTokenNotFound", "refresh token not found", "refresh token not found", 401)
	RefreshTokenRevoked  = Add("RefreshTokenRevoked", "refresh token revoked", "refresh token revoked", 401)
	RefreshTokenExpired  = Add("RefreshTokenExpired", "refresh token expired", "refresh token expired", 401)
	
	DbUnavailable = Add("DbUnavailable", "database unavailable", "service unavailable", http.StatusServiceUnavailable)
	DbRecordNotFound = Add(gorm.ErrRecordNotFound.Error(), gorm.ErrRecordNotFound.Error(), "resource not found", 404)

	BlogNotFound  = Add("BlogNotFound", "blog not found", "blog not found", http.StatusNotFound)
)
