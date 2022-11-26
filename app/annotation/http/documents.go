package http

import (
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/pkg/response"
	"strconv"

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
	requestBody := &request.ReqEditDocument{}
	err := ctx.Bind(requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	err = validator.New().Struct(requestBody)
	if err != nil {
		return err
	}
	document, err := h.App.UpdateDocumentsById(ctx, *requestBody)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, document)
}

func (h *AnnotationHandler) GetDocumentById(ctx echo.Context) error {
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	document, err := h.App.GetDocumentsById(ctx, int64(id))
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, document)
}

func (h *AnnotationHandler) GetAllDocuments(ctx echo.Context) error {
	document, err := h.App.GetAllDocuments(ctx)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, document)
}

func (h *AnnotationHandler) DeleteDocumentsByID(ctx echo.Context) error {
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	result, err := h.App.DeleteDocumentsByID(ctx, int64(id))
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, result)
}
