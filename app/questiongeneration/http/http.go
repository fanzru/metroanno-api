package http

import (
	"metroanno-api/app/questiongeneration/usecase"
	"metroanno-api/infrastructure/config"
)

type QuestionGeneratioHandler struct {
	App usecase.QuestionGenerationApp
	Cfg config.Config
}
