package router

import (
	"app/domain/model"
	"app/presenter/http/handler"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"net/http"
)

// https://pkg.go.dev/github.com/labstack/echo/v5/middleware

// CreateJWTGoParseTokenFunc creates JWTGo implementation for ParseTokenFunc
//
// signingKey is signing key to validate token.
// This is one of the options to provide a token validation key.
// The order of precedence is a user-defined SigningKeys and SigningKey.
// Required if signingKeys is not provided.
//
// signingKeys is Map of signing keys to validate token with kid field usage.
// This is one of the options to provide a token validation key.
// The order of precedence is a user-defined SigningKeys and SigningKey.
// Required if signingKey is not provided
func CreateJWTGoParseTokenFunc(signingKey interface{}, signingKeys map[string]interface{}) func(c echo.Context, auth string) (interface{}, error) {
	// keyFunc defines a user-defined function that supplies the public key for a token validation.
	// The function shall take care of verifying the signing algorithm and selecting the proper key.
	// A user-defined KeyFunc can be useful if tokens are issued by an external party.
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != middleware.AlgorithmHS256 {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		if len(signingKeys) == 0 {
			return signingKey, nil
		}

		if kid, ok := t.Header["kid"].(string); ok {
			if key, ok := signingKeys[kid]; ok {
				return key, nil
			}
		}
		return nil, fmt.Errorf("unexpected jwt key id=%v", t.Header["kid"])
	}

	return func(c echo.Context, auth string) (interface{}, error) {
		// token, err := jwt.ParseWithClaims(auth, jwt.MapClaims{}, keyFunc) // you could add your default claims here
		token, err := jwt.ParseWithClaims(auth, &model.JwtCustomClaims{}, keyFunc)
		if err != nil {
			return nil, err
		}
		if !token.Valid {
			return nil, errors.New("invalid token")
		}
		return token, nil
	}
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Golang Frame Work Echo V5-alpha is Running ...")
}

// NewRouter Routerの設定を行います.
func NewRouter(e *echo.Echo, h handler.AppHandler) {
	e.GET("/", hello)

	api := e.Group("/api")
	api.POST("/signup", h.CreateUser)
	api.POST("/login", h.Login)

	// https://qiita.com/x-color/items/24ff2491751f55e866cf
	// ログイン後のtoken認証
	// Configure middleware with the custom claims type
	/**
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}*/

	apiAuth := e.Group("/api/auth")
	// api/authの下のルートはJWTの認証が必要
	apiAuth.Use(middleware.JWTWithConfig(
		middleware.JWTConfig{
			ParseTokenFunc: CreateJWTGoParseTokenFunc([]byte("secret"), nil),
		}))
	apiAuth.GET("/refresh", h.Refresh)
	apiAuth.POST("/logout", h.Logout)
}
