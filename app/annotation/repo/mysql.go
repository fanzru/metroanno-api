package repo

import (
	"errors"
	usersmodel "metroanno-api/app/accounts/domain/models"
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Impl interface {
	CreateTheory(ctx echo.Context, document models.Document) error
	GetDocumentsById(ctx echo.Context, documentId int64) (*models.Document, error)
	GetAllDocuments(ctx echo.Context) (*[]models.Document, error)
	UpdateDocumentsById(ctx echo.Context, param request.ReqEditDocument) (*models.Document, error)
	GetAllQuestionTypes(ctx echo.Context) ([]models.QuestionType, error)
	CreateQuestionTypes(ctx echo.Context, param request.ReqAddQuestionType) (*models.QuestionType, error)
	DeleteQuestionTypes(ctx echo.Context, id int64) (*models.QuestionType, error)
	CreateFeedback(ctx echo.Context, feedback models.Feedback) (*models.Feedback, error)
	GetDocumentIdByUserId(ctx echo.Context) (int64, error)
	DeleteDocumentsByID(ctx echo.Context, id int64) (*models.Document, error)
	BulkInsertQuestionAnnotations(ctx echo.Context, arrQuestionAnnotations []models.QuestionAnnotation) (*[]models.QuestionAnnotation, *gorm.DB, error)
	UpdateUsersByIdWithTX(ctx echo.Context, docId int64, userID int64, tx *gorm.DB) (*usersmodel.User, *gorm.DB, error)
	CreateDoneDocumentUser(ctx echo.Context, docId int64, userID int64) error
}

type AnnotationsRepo struct {
	MySQL database.Connection
	Cfg   config.Config
}

func New(a AnnotationsRepo) AnnotationsRepo {
	return a
}

const TableQuestionAnnotations = "question_annotations"
const TableUsers = "users"
const TableDocuments = "documents"

var (
	ErrDocumentNotFound  = errors.New("you are not included in any document")
	ErrDocumentIsDeleted = errors.New("your document assigned is deleted")
	ErrInstanceNotFound  = errors.New("user not found")
)
