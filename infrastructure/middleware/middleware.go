package middleware

import (
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"
	"metroanno-api/pkg/jwt"
	"metroanno-api/pkg/response"
	"strings"

	"github.com/labstack/echo/v4"
)

// Service Authorizer
type MiddlewareAuth struct {
	DB  database.Connection
	Cfg config.Config
}

func NewServiceAuthorizer(db database.Connection, cfg config.Config) MiddlewareAuth {
	return MiddlewareAuth{
		DB:  db,
		Cfg: cfg,
	}
}

func (m MiddlewareAuth) BearerTokenMiddlewareAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId, typeUser := m.getUserIdAndTypeFromJWT(ctx)
		if userId == 0 {
			return response.ResponseErrorUnauthorized(ctx)
		}
		if typeUser != 2 {
			return response.ResponseErrorUnauthorized(ctx)
		}
		return next(ctx)
	}
}
func (m MiddlewareAuth) BearerTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId, _ := m.getUserIdAndTypeFromJWT(ctx)
		if userId == 0 {
			return response.ResponseErrorUnauthorized(ctx)
		}
		return next(ctx)
	}
}

func (m MiddlewareAuth) getUserIdAndTypeFromJWT(ctx echo.Context) (int, uint64) {
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0, 0
	}

	headerSplit := strings.Split(authHeader, " ")
	if len(headerSplit) < 1 {
		return 0, 0
	}

	// Get Token
	token := headerSplit[1]

	// Get JWT Claims
	claims, err := jwt.DecodeToken(token, m.Cfg.JWTTokenSecret)
	if err != nil {
		return 0, 0
	}

	// Bind UserID to context
	ctx.Set("user_id", claims.UserId)
	ctx.Set("usernamel", claims.Username)
	return 1, claims.Type
}
