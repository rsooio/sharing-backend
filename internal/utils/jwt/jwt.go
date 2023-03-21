package jwt

import (
	"backend/internal/pkg/res"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID    uint             `json:"uid"`
	Resources res.ResourceFlag `json:"res"`
}

func (c *Claims) Validate() error {
	return nil
}

func (c *Claims) Has(flags ...res.ResourceFlag) bool {
	for _, flag := range flags {
		if c.Resources&flag == 0 {
			return false
		}
	}
	return true
}

func Sign(uid uint, resources res.ResourceFlag, secret string, expireSeconds int64) (token string, err error) {
	claims := Claims{
		UserID:           uid,
		Resources:        resources,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	claims.IssuedAt = &jwt.NumericDate{Time: time.Now()}
	claims.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(time.Second * time.Duration(expireSeconds))}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

func Parse(token, secret string) (claims *Claims, err error) {
	claims = new(Claims)
	t, err := jwt.ParseWithClaims(token, *claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, errors.New("claims parse error")
	}
}
