package http

import (
	"fmt"
	"metroanno-api/app/questiongeneration/domain/params"
	"metroanno-api/app/questiongeneration/domain/request"
	"metroanno-api/pkg/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *QuestionGeneratioHandler) SaveQuestions(ctx echo.Context) error {

	requestBody := &request.ReqSaveQuestions{}
	err := ctx.Bind(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = h.App.SaveQuestions(ctx, *requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, nil)
}

func (h *QuestionGeneratioHandler) FindQuestions(ctx echo.Context) error {
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return err
	}
	historyIDStr := ctx.QueryParam("history_id")
	if historyIDStr == "" {
		historyIDStr = "0"
	}
	historyID, err := strconv.ParseInt(historyIDStr, 10, 64)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	questions, err := h.App.GetHistoryQuestionUser(ctx, params.FilterQuestions{
		UserID:    userID,
		HistoryID: historyID,
	})
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, questions)
}
