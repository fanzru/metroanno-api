package usecase

import (
	"fmt"
	"metroanno-api/app/annotation/domain/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) GetAllQAuser(ctx echo.Context, pageNumber int64) (response.PaginationQA, error) {
	userID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return response.PaginationQA{}, err
	}

	r, err := a.AnnotationsRepo.GetAllQuestionAnnotationsUser(ctx, userID, pageNumber)
	if err != nil {
		return r, err
	}
	return r, nil
}
