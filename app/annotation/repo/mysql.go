package repo

import (
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"

	"github.com/labstack/echo/v4"
)

type Impl interface {
	CreateTheory(ctx echo.Context, document models.Document) error
	GetDocumentsById(ctx echo.Context, documentId int64) (*models.Document, error)
	GetAllDocuments(ctx echo.Context) (*[]models.Document, error)
	UpdateDocumentsById(ctx echo.Context, param request.ReqEditDocument) (*models.Document, error)
	GetAllQuestionTypes(ctx echo.Context) ([]models.QuestionType, error)
	CreateQuestionTypes(ctx echo.Context, param request.ReqAddQuestionType) (*models.QuestionType, error)
	DeleteQuestionTypes(ctx echo.Context, id int64) (*models.QuestionType, error)
}

type AnnotationsRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(a AnnotationsRepo) AnnotationsRepo {
	return a
}
