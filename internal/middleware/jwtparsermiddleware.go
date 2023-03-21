package middleware

import (
	"backend/internal/utils/jwt"
	"net/http"
)

type JWTParserMiddleware struct {
	Secret string
	Claims **jwt.Claims
}

func NewJWTParserMiddleware(secret string, claims **jwt.Claims) *JWTParserMiddleware {
	return &JWTParserMiddleware{Secret: secret, Claims: claims}
}

func (m *JWTParserMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := jwt.Parse(r.Header.Get("Authorization"), m.Secret)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}

		m.Claims = &claims
		next(w, r)
	}
}
