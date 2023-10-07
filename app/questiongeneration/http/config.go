package http

import (
	"metroanno-api/app/questiongeneration/domain/models"
	"metroanno-api/pkg/response"

	"github.com/labstack/echo/v4"
)

func (h *QuestionGeneratioHandler) FindConfig(ctx echo.Context) error {
	return response.ResponseSuccessOK(ctx, models.QuestionType{
		Random:   h.Cfg.MyMapRandom,
		Bloom:    h.Cfg.MyMapBloom,
		Graesser: h.Cfg.MyMapGraesser,
	})
}
