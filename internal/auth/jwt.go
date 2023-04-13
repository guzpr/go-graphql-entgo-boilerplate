package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

func Decode(token string) (*Claims, error) {
	claims := Claims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		signingKey := os.Getenv("JWT_SIGNING_KEY_SECRET")
		if signingKey == "" {
			return "", errors.New("JWT signing key is empty")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func Encode(claims Claims) (string, error) {
	if err := claims.Valid(); err != nil {
		return "", err
	}
	signingKey := os.Getenv("JWT_SIGNING_KEY_SECRET")
	if signingKey == "" {
		return "", errors.New("JWT signing key is empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func NewClaims(opt ...func(*Claims)) Claims {
	claims := &Claims{}
	claims.Issuer = os.Getenv("SERVER")
	claims.ID = uuid.New().String()
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	for _, o := range opt {
		o(claims)
	}
	return *claims
}

func (c Claims) Valid() error {
	vErr := new(jwt.ValidationError)
	now := time.Now()

	// Taken from jwt-v4 go RegisteredClaims.Valid()
	if !c.VerifyExpiresAt(now, false) {
		delta := now.Sub(c.ExpiresAt.Time)
		vErr.Inner = fmt.Errorf("%s by %s", jwt.ErrTokenExpired, delta)
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	if !c.VerifyIssuedAt(now, false) {
		vErr.Inner = jwt.ErrTokenUsedBeforeIssued
		vErr.Errors |= jwt.ValidationErrorIssuedAt
	}

	if !c.VerifyNotBefore(now, false) {
		vErr.Inner = jwt.ErrTokenNotValidYet
		vErr.Errors |= jwt.ValidationErrorNotValidYet
	}

	// Our custom validation
	if c.UserID == "" {
		vErr.Inner = jwt.ErrTokenInvalidClaims
		vErr.Errors |= jwt.ValidationErrorClaimsInvalid
	}

	if vErr.Errors != 0 {
		return vErr
	}

	return nil
}

func WithAudience(aud []string) func(*Claims) {
	return func(c *Claims) {
		c.Audience = aud
	}
}

func WithExpiresAt(time time.Time) func(*Claims) {
	return func(c *Claims) {
		c.ExpiresAt = jwt.NewNumericDate(time)
	}
}

func WithNotBefore(time time.Time) func(*Claims) {
	return func(c *Claims) {
		c.NotBefore = jwt.NewNumericDate(time)
	}
}

func WithUserID(userID string) func(*Claims) {
	return func(c *Claims) {
		c.UserID = userID
	}
}
