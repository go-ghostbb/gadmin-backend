package errorx

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	// ErrFailedAuthentication indicates authentication failed, could be faulty username or password
	ErrFailedAuthentication = gerror.New("incorrect Username or Password")

	// ErrFailedTokenCreation indicates JWT Token failed to create, reason unknown
	ErrFailedTokenCreation = gerror.New("failed to create JWT Token")

	// ErrExpiredToken indicates JWT token has expired. Can't refresh.
	ErrExpiredToken = gerror.New("token is expired") // in practice, this is generated from the jwt library not by us

	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = gerror.New("auth header is empty")

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = gerror.New("auth header is invalid")

	// ErrInvalidToken indicates JWT token has invalid. Can't refresh.
	ErrInvalidToken = gerror.New("token is invalid")
)

var (
	// ErrRecordNotFound record not found
	ErrRecordNotFound = gerror.NewCode(gcode.New(10001, "record not found", nil))
)
