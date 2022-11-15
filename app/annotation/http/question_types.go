package http

import (
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/pkg/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *AnnotationHandler) AddQuestionsTypes(ctx echo.Context) error {
	requestBody := request.ReqAddQuestionType{}
	err := ctx.Bind(&requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	result, err := h.App.CreateQuestionTypes(ctx, requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, result)
}

func (h *AnnotationHandler) DeleteQuestionsTypes(ctx echo.Context) error {
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	result, err := h.App.DeleteQuestionTypes(ctx, int64(id))
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, result)
}

func (h *AnnotationHandler) GetAllQuestionsTypes(ctx echo.Context) error {
	result, err := h.App.GetAllQuestionTypes(ctx)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, result)
}
