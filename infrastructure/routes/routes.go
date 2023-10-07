package routes

import (
	"log"
	accounts "metroanno-api/app/accounts/http"
	annotations "metroanno-api/app/annotation/http"
	questiongeneration "metroanno-api/app/questiongeneration/http"
	"metroanno-api/infrastructure/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ModuleHandler struct {
	MiddlewareAuth           middleware.MiddlewareAuth
	AccountHandler           accounts.AccountHandler
	AnnotationsHandler       annotations.AnnotationHandler
	QuestionGeneratioHandler questiongeneration.QuestionGeneratioHandler
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

	// admin
	admingateway := app.Group("/admin")
	admingateway.POST("/register", h.AccountHandler.RegisterAdmin)
	admingateway.GET("/users", h.MiddlewareAuth.BearerTokenMiddlewareAdmin(h.AccountHandler.GetAllUserNonAdmin))
	admingateway.PATCH("/users", h.MiddlewareAuth.BearerTokenMiddlewareAdmin(h.AccountHandler.UpdateStatusUsers))
	admingateway.GET("/documents", h.MiddlewareAuth.BearerTokenMiddlewareAdmin(h.AnnotationsHandler.GetAllDocumentsAdmin))
	admingateway.PATCH("/documents", h.MiddlewareAuth.BearerTokenMiddlewareAdmin(h.AnnotationsHandler.UpdateIsAprrovedDocument))
	admingateway.PATCH("/question", h.MiddlewareAuth.BearerTokenMiddlewareAdmin(h.AnnotationsHandler.UpdateIsCheckedAdminQuestionAnnotations))

	// accounts
	accountsgateway := app.Group("/accounts")
	accountsgateway.POST("/register", h.AccountHandler.RegisterUser)
	accountsgateway.POST("/login/annotator", h.AccountHandler.Login)
	accountsgateway.GET("/user", h.MiddlewareAuth.BearerTokenMiddleware(h.AccountHandler.UserProfile))

	// subject
	subjectgateway := app.Group("/subject")
	subjectgateway.GET("/", h.AnnotationsHandler.GetAllSubjects)

	// documents
	documentsgateway := app.Group("/documents")
	documentsgateway.POST("/add", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.AddTheory))
	documentsgateway.PUT("/edit", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.EditTheory))
	documentsgateway.GET("/", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.GetAllDocuments))
	documentsgateway.GET("/:id", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.GetDocumentById))
	documentsgateway.DELETE("/delete/:id", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.DeleteDocumentsByID))
	documentsgateway.GET("/random-user", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.RandomDocuments))
	documentsgateway.GET("/added", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.GetAllDocumentsByCreatedBy))

	// question-type
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
	questionannotationsgateway.POST("/mark", h.MiddlewareAuth.BearerTokenMiddlewareAdmin(h.AnnotationsHandler.MarkQuestionAnnotations))
	questionannotationsgateway.GET("/user", h.MiddlewareAuth.BearerTokenMiddleware(h.AnnotationsHandler.GetAllQAuser))

	// question generation
	questiongenerationgateway := app.Group("/question-generation")
	questiongenerationgateway.POST("/chat-gpt", h.MiddlewareAuth.BearerTokenMiddleware(h.QuestionGeneratioHandler.GenerateQuestion))
	questiongenerationgateway.POST("/save", h.MiddlewareAuth.BearerTokenMiddleware(h.QuestionGeneratioHandler.SaveQuestions))
	questiongenerationgateway.GET("/histories", h.MiddlewareAuth.BearerTokenMiddleware(h.QuestionGeneratioHandler.FindQuestions))
	questiongenerationgateway.GET("/question-type", h.QuestionGeneratioHandler.FindConfig)
	return app
}
