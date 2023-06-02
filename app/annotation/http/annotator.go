package http

import (
	"fmt"
	errs "metroanno-api/app/annotation/domain/errors"
	"metroanno-api/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *AnnotationHandler) GetAllQAuser(ctx echo.Context) error {
	PN := ctx.QueryParam("pageNumber")

	if PN == "" {
		return response.ResponseErrorBadRequest(ctx, errs.ErrPageNumber)
	}

	pageNumber, err := strconv.ParseInt(fmt.Sprintf("%v", PN), 10, 64)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	r, err := h.App.GetAllQAuser(ctx, pageNumber)

	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}

	return response.ResponseSuccessOK(ctx, r)
}
