package http

import (
	"metroanno-api/infrastructure/config"
	"metroanno-api/pkg/response"

	"github.com/labstack/echo/v4"
)

type AnnotationHandler struct {
	Cfg config.Config
}

func (h *AnnotationHandler) AddTheory(ctx echo.Context) error {

	return response.ResponseSuccessOK(ctx, nil)
}

func (h *AnnotationHandler) EditTheory(ctx echo.Context) error {

	return response.ResponseSuccessOK(ctx, nil)
}
