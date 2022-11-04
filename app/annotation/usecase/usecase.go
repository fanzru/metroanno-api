package usecase

import (
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/app/annotation/repo"
	"metroanno-api/infrastructure/config"

	"github.com/labstack/echo/v4"
)

type Impl interface {
	AddDocument(ctx echo.Context, param request.ReqAddDocument) error
}

type AnnotationsApp struct {
	AnnotationsRepo repo.AnnotationsRepo
	Cfg             config.Config
}

func New(a AnnotationsApp) AnnotationsApp {
	return a
}
