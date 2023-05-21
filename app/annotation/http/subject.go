package http

import (
	"metroanno-api/pkg/response"

	"github.com/labstack/echo/v4"
)

func (h *AnnotationHandler) GetAllSubjects(ctx echo.Context) error {
	s, err := h.App.GetAllSubject(ctx)
	if err != nil {
		return response.ResponseErrorBadRequest(ctx, err)
	}
	return response.ResponseSuccessOK(ctx, s)
}
