package services

import (
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"
	"metroanno-api/infrastructure/middleware"
)

func RegisterMiddleware(db database.Connection, cfg config.Config) middleware.MiddlewareAuth {
	middlewareAuth := middleware.NewServiceAuthorizer(db, cfg)
	return middlewareAuth
}
