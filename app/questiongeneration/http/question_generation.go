package http

import (
	"metroanno-api/app/questiongeneration/domain/request"
	"metroanno-api/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *QuestionGeneratioHandler) GenerateQuestion(ctx echo.Context) error {

	requestBody := &request.ReqGenerateQuestion{}
	err := ctx.Bind(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	responseChatGPT, err := h.App.GenerateQuestion(ctx, *requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	//
	return response.ResponseSuccessOK(ctx, responseChatGPT)
}
