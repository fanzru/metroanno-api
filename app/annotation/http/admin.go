package http

import (
	"fmt"
	errs "metroanno-api/app/annotation/domain/errors"
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/pkg/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (h *AnnotationHandler) GetAllDocumentsAdmin(ctx echo.Context) error {
	PN := ctx.QueryParam("pageNumber")

	if PN == "" {
		return response.ResponseErrorBadRequest(ctx, errs.ErrPageNumber)
	}

	pageNumber, err := strconv.ParseInt(fmt.Sprintf("%v", PN), 10, 64)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	lim := ctx.QueryParam("limit")

	if lim == "" {
		lim = "10"
	}

	limit, err := strconv.ParseInt(fmt.Sprintf("%v", lim), 10, 64)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	r, err := h.App.GetAllDocumentsAdmin(ctx, pageNumber, limit)

	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, r)
}

func (h *AnnotationHandler) UpdateIsAprrovedDocument(ctx echo.Context) error {
	requestBody := request.ReqUpdateDocument{}
	err := ctx.Bind(&requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = h.App.UpdateIsAprrovedDocument(ctx, requestBody.DocumentID, requestBody.IsApproved)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, nil)
}

func (h *AnnotationHandler) UpdateIsCheckedAdminQuestionAnnotations(ctx echo.Context) error {
	requestBody := request.ReqCheckedQuestion{}
	err := ctx.Bind(&requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = h.App.UpdateIsCheckedAdminQuestionAnnotations(ctx, requestBody.QuestionID, requestBody.IsChecked)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, nil)
}

func (h *AnnotationHandler) MarkQuestionAnnotations(ctx echo.Context) error {
	requestBody := request.ReqMarkQuestion{}
	err := ctx.Bind(&requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = h.App.MarkQuestionAnnotations(ctx, requestBody.QuestionIDs, requestBody.Mark)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, nil)
}
