package http

import (
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *AnnotationHandler) CreateFeedback(ctx echo.Context) error {
	requestBody := request.ReqCreateFeedback{}
	err := ctx.Bind(&requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	result, err := h.App.CreateFeeback(ctx, requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, result)
}
