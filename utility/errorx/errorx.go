package errorx

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	// ErrMissingSecretKey indicates Secret key is required
	ErrMissingSecretKey = gerror.New("secret key is required")

	// ErrForbidden when HTTP status 403 is given
	ErrForbidden = gerror.New("you don't have permission to access this resource")

	// ErrMissingAuthenticatorFunc indicates Authenticator is required
	ErrMissingAuthenticatorFunc = gerror.New("ginJWTMiddleware.Authenticator func is undefined")

	// ErrMissingLoginValues indicates a user tried to authenticate without username or password
	ErrMissingLoginValues = gerror.New("missing Username or Password")

	// ErrFailedAuthentication indicates authentication failed, could be faulty username or password
	ErrFailedAuthentication = gerror.New("incorrect Username or Password")

	// ErrFailedTokenCreation indicates JWT Token failed to create, reason unknown
	ErrFailedTokenCreation = gerror.New("failed to create JWT Token")

	// ErrExpiredToken indicates JWT token has expired. Can't refresh.
	ErrExpiredToken = gerror.New("token is expired") // in practice, this is generated from the jwt library not by us

	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = gerror.New("auth header is empty")

	// ErrMissingExpField missing exp field in token
	ErrMissingExpField = gerror.New("missing exp field")

	// ErrWrongFormatOfExp field must be float64 format
	ErrWrongFormatOfExp = gerror.New("exp must be float64 format")

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = gerror.New("auth header is invalid")

	// ErrEmptyQueryToken can be thrown if authing with URL Query, the query token variable is empty
	ErrEmptyQueryToken = gerror.New("query token is empty")

	// ErrEmptyCookieToken can be thrown if authing with a cookie, the token cookie is empty
	ErrEmptyCookieToken = gerror.New("cookie token is empty")

	// ErrEmptyParamToken can be thrown if authing with parameter in path, the parameter in path is empty
	ErrEmptyParamToken = gerror.New("parameter token is empty")

	// ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512, RS256, RS384 or RS512
	ErrInvalidSigningAlgorithm = gerror.New("invalid signing algorithm")

	// ErrNoPrivKeyFile indicates that the given private key is unreadable
	ErrNoPrivKeyFile = gerror.New("private key file unreadable")

	// ErrNoPubKeyFile indicates that the given public key is unreadable
	ErrNoPubKeyFile = gerror.New("public key file unreadable")

	// ErrInvalidPrivKey indicates that the given private key is invalid
	ErrInvalidPrivKey = gerror.New("private key invalid")

	// ErrInvalidPubKey indicates the the given public key is invalid
	ErrInvalidPubKey = gerror.New("public key invalid")

	// ErrMissingIdentity identity key and identity value is null
	ErrMissingIdentity = gerror.New("payload don't have identity key and identity value")

	// ErrMissingContext indicates Context is required
	ErrMissingContext = gerror.New("context is required")

	// ErrInvalidToken indicates JWT token has invalid. Can't refresh.
	ErrInvalidToken = gerror.New("token is invalid")
)

var (
	// ErrRecordNotFound record not found
	ErrRecordNotFound = gerror.NewCode(gcode.New(10001, "record not found", nil))
)
