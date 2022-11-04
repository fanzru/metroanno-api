package http

import (
	"metroanno-api/app/annotation/usecase"
	"metroanno-api/infrastructure/config"
)

type AnnotationHandler struct {
	App usecase.AnnotationsApp
	Cfg config.Config
}
