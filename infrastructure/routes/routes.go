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
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    "good connection",
			"message": "metroanno api 200 OK",
		})
	})

	//accounts
	accountsgateway := app.Group("/accounts")
	accountsgateway.POST("/register/annotator", h.AccountHandler.RegisterUser)
	accountsgateway.POST("/register/admin", h.AccountHandler.RegisterAdmin)
	accountsgateway.POST("/login/annotator", h.AccountHandler.Login)
	accountsgateway.GET("/user", h.MiddlewareAuth.BearerTokenMiddleware(h.AccountHandler.UserProfile))

	//documents
	documentsgateway := app.Group("/documents")
	documentsgateway.POST("/add", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.AddTheory))
	documentsgateway.PUT("/edit", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.EditTheory))
	documentsgateway.GET("/", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.GetAllDocuments))
	documentsgateway.GET("/:id", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.GetDocumentById))
	documentsgateway.DELETE("/delete/:id", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.DeleteDocumentsByID))
	documentsgateway.POST("/random-user", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.RandomDocuments))

	//question-type
	questiongateway := app.Group("/question-type")
	questiongateway.POST("/create", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.AddQuestionsTypes))
	questiongateway.GET("/", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.GetAllQuestionsTypes))
	questiongateway.DELETE("/delete/:id", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.DeleteQuestionsTypes))

	// feedback
	feedbackgateway := app.Group("/feedback")
	feedbackgateway.POST("/create", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.CreateFeedback))

	// question annotations
	questionannotationsgateway := app.Group("/question-annotations")
	questionannotationsgateway.POST("/bulk", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.BulkInsertQuestion))

	return app
}
