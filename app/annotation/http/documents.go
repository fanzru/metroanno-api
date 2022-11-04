package http

import (
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *AnnotationHandler) AddTheory(ctx echo.Context) error {
	requestBody := &request.ReqAddDocument{}
	err := ctx.Bind(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = h.App.AddDocument(ctx, *requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, nil)
}

func (h *AnnotationHandler) EditTheory(ctx echo.Context) error {

	return response.ResponseSuccessOK(ctx, nil)
}
