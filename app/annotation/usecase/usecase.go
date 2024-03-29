package usecase

import (
	"errors"
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/app/annotation/domain/response"
	"metroanno-api/app/annotation/repo"
	"metroanno-api/infrastructure/config"

	"github.com/labstack/echo/v4"
)

type Impl interface {
	AddDocument(ctx echo.Context, param request.ReqAddDocument) error
	UpdateDocumentsById(ctx echo.Context, param request.ReqEditDocument) (*models.Document, error)
	GetAllDocuments(ctx echo.Context) ([]models.Document, error)
	GetDocumentsById(ctx echo.Context, documentId int64) (*models.Document, error)
	GetAllQuestionTypes(ctx echo.Context) (*[]models.QuestionType, error)
	CreateQuestionTypes(ctx echo.Context, param request.ReqAddQuestionType) (*models.QuestionType, error)
	DeleteQuestionTypes(ctx echo.Context, id int64) (*models.QuestionType, error)
	RandomDocuments(ctx echo.Context) (int64, error)
	GetAllDocumentsAdmin(ctx echo.Context, pageNumber int64) (response.Pagination, error)
	UpdateIsAprrovedDocument(ctx echo.Context, documentID int64, isApproved bool) error
	UpdateIsCheckedAdminQuestionAnnotations(ctx echo.Context, id int64, isChecked bool) error
	GetDocumentsByCreatedBy(ctx echo.Context) ([]models.Document, error)
	GetAllSubject(ctx echo.Context) ([]models.Subject, error)
	MarkQuestionAnnotations(ctx echo.Context, ids []int64, mark bool) error
	GetAllQAuser(ctx echo.Context, pageNumber int64) (response.PaginationQA, error)
}

type AnnotationsApp struct {
	AnnotationsRepo repo.AnnotationsRepo
	Cfg             config.Config
}

func New(a AnnotationsApp) AnnotationsApp {
	return a
}

var (
	ErrNotHaveDocuments    = errors.New("dont have document to assign")
	ErrDifferentDocumentID = errors.New("please submit for your document")
)
