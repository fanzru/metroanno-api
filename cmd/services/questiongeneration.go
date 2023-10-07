package services

import (
	qghandler "metroanno-api/app/questiongeneration/http"
	qgrepo "metroanno-api/app/questiongeneration/repo"
	qgusecase "metroanno-api/app/questiongeneration/usecase"
	"metroanno-api/infrastructure/config"
	"metroanno-api/infrastructure/database"
)

func RegisterServiceQuestionGeneration(db database.Connection, cfg config.Config) qghandler.QuestionGeneratioHandler {
	DB := qgrepo.New(qgrepo.QuestionGenerationRepo{
		MySQL: db,
		Cfg:   cfg,
	})

	App := qgusecase.New(qgusecase.QuestionGenerationApp{
		QuestionGenerationRepo: DB,
		Cfg:                    cfg,
	})

	Handler := qghandler.QuestionGeneratioHandler{
		App: App,
		Cfg: cfg,
	}

	return Handler
}
