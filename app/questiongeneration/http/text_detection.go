package http

import (
	"fmt"
	"metroanno-api/app/questiongeneration/domain/request"
	"metroanno-api/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *QuestionGeneratioHandler) TextDetection(ctx echo.Context) error {

	requestBody := &request.ReqTextDetection{}
	err := ctx.Bind(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	fmt.Println()

	topic, err := h.App.TextDetection(ctx, *requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	//
	return response.ResponseSuccessOK(ctx, topic)
}
