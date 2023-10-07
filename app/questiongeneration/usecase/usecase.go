package usecase

import (
	"metroanno-api/app/questiongeneration/domain/models"
	"metroanno-api/app/questiongeneration/domain/params"
	"metroanno-api/app/questiongeneration/domain/request"
	"metroanno-api/app/questiongeneration/domain/response"
	"metroanno-api/app/questiongeneration/repo"
	"metroanno-api/infrastructure/config"

	"github.com/labstack/echo/v4"
)

type Impl interface {
	GenerateQuestion(ctx echo.Context, params request.ReqGenerateQuestion) ([]response.JSONResponse, error)
	GetHistoryQuestionUser(ctx echo.Context, params params.FilterQuestions) ([]models.QuestionsHistory, error)
	SaveQuestions(ctx echo.Context, params request.ReqSaveQuestions) error
}

type QuestionGenerationApp struct {
	QuestionGenerationRepo repo.QuestionGenerationRepo
	Cfg                    config.Config
}

func New(a QuestionGenerationApp) QuestionGenerationApp {
	return a
}
