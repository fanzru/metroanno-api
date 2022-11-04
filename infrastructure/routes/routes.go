package routes

import (
	"log"
	accounts "metroanno-api/app/accounts/http"
	annotations "metroanno-api/app/annotation/http"
	"metroanno-api/infrastructure/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ModuleHandler struct {
	MiddlewareAuth     middleware.MiddlewareAuth
	AccountHandler     accounts.AccountHandler
	AnnotationsHandler annotations.AnnotationHandler
}

func NewRoutes(h ModuleHandler, app *echo.Echo) *echo.Echo {

	log.Println("Starting to create new routing...")

	// test api connect or not
	app.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "FANZRU PASTI LULUS S1 INFORMATIKA 200 OK")
	})

	//accounts
	accountsgateway := app.Group("/accounts")
	accountsgateway.POST("/register/annotator", h.AccountHandler.RegisterUser)
	accountsgateway.POST("/register/admin", h.AccountHandler.RegisterAdmin)
	accountsgateway.POST("/login/annotator", h.AccountHandler.Login)

	//documents
	documentsgateway := app.Group("/documents")
	documentsgateway.POST("/add", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.AddTheory))
	documentsgateway.PUT("/edit", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.EditTheory))
	documentsgateway.GET("/", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.GetAllDocuments))
	documentsgateway.GET("/:id", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.GetDocumentById))
	return app
}
