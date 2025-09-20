package err

import "errors"

type AppErr struct {
    BlogNotFound    error
    DbUnavailable   error
    EmailExists     error
    InvalidUID      error
    InvalidUserName error
    InvalidUser     error
    InvalidSlug     error
    UserExists      error
}

// singleton instance
var Err = &AppErr{
    BlogNotFound:    errors.New("blog not found"),
    DbUnavailable:   errors.New("database unavailable"),
    EmailExists:     errors.New("email already registered"),
    InvalidUID:      errors.New("invalid UID"),
    InvalidUserName: errors.New("invalid username"),
    InvalidUser:     errors.New("invalid user"),
    InvalidSlug:     errors.New("invalid slug"),
    UserExists:      errors.New("user already exists"),
}
