package main

import (
	"log"
	"metroanno-api/cmd/services"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"
	"metroanno-api/infrastructure/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("Start Services....")

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// middleware
	middlewareAuth := services.RegisterMiddleware(db, cfg)

	// services
	accountsHandler := services.RegisterServiceAccounts(db, cfg)
	annotationsHandler := services.RegisterServiceAnnotations(db, cfg)
	qgHandler := services.RegisterServiceQuestionGeneration(db, cfg)

	// register routes
	mHandler := routes.ModuleHandler{
		AccountHandler:           accountsHandler,
		MiddlewareAuth:           middlewareAuth,
		AnnotationsHandler:       annotationsHandler,
		QuestionGeneratioHandler: qgHandler,
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e = routes.NewRoutes(mHandler, e)

	log.Fatal(e.Start(":8889"))
}
