package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/mritd/ginmvc/conf"

	"github.com/dgrijalva/jwt-go"
)

const (
	JWTSigningMethodHS256 = "HS256"
	JWTSigningMethodHS384 = "HS384"
	JWTSigningMethodHS512 = "HS512"
)

const JWTClaimsKey = "JWT_CLAIMS"

var (
	JWTTokenMalformed   = errors.New("jwt token malformed")
	JWTTokenExpired     = errors.New("jwt token expired")
	JWTTokenNotValidYet = errors.New("jwt token not valid yet")
	JWTTokenInvalid     = errors.New("jwt token invalid")
)

type JWT struct {
	SigningKey []byte
}

type JWTClaims struct {
	jwt.StandardClaims
}

// create jwt token
func (j *JWT) CreateToken(claims JWTClaims) (string, error) {
	token := jwt.NewWithClaims(getSigningMethod(), claims)
	return token.SignedString(j.SigningKey)
}

// parse token
func (j *JWT) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, JWTTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, JWTTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, JWTTokenNotValidYet
			} else {
				return nil, JWTTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, JWTTokenInvalid
}

// update token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	return j.RefreshTokenWithTime(tokenString, 1*time.Hour)
}

// update token with time
func (j *JWT) RefreshTokenWithTime(tokenString string, t time.Duration) (string, error) {
	// prevent tokens from failing
	jwt.TimeFunc = func() time.Time { return time.Unix(0, 0) }
	defer func() { jwt.TimeFunc = time.Now }()

	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	jwt.TimeFunc = time.Now
	claims.ExpiresAt = time.Now().Add(t).Unix()
	return j.CreateToken(*claims)
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(conf.Basic.JWT.Secret),
	}
}

func getSigningMethod() *jwt.SigningMethodHMAC {
	var signingMethod *jwt.SigningMethodHMAC

	switch strings.ToUpper(conf.Basic.JWT.SigningMethod) {
	case JWTSigningMethodHS256:
		signingMethod = jwt.SigningMethodHS256
	case JWTSigningMethodHS384:
		signingMethod = jwt.SigningMethodHS384
	case JWTSigningMethodHS512:
		signingMethod = jwt.SigningMethodHS512
	default:
		signingMethod = jwt.SigningMethodHS256
	}
	return signingMethod
}
