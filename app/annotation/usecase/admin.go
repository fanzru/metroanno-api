package usecase

import (
	"metroanno-api/app/annotation/domain/response"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) GetAllDocumentsAdmin(ctx echo.Context, pageNumber int64, limit int64) (response.Pagination, error) {
	documents, err := a.AnnotationsRepo.GetAllDocumentsAdmin(ctx, pageNumber, limit)
	if err != nil {
		return documents, err
	}
	return documents, nil
}

func (a *AnnotationsApp) UpdateIsAprrovedDocument(ctx echo.Context, documentID int64, isApproved bool) error {
	err := a.AnnotationsRepo.UpdateIsAprrovedDocument(ctx, documentID, isApproved)
	if err != nil {
		return err
	}
	return nil
}

func (a *AnnotationsApp) UpdateIsCheckedAdminQuestionAnnotations(ctx echo.Context, id int64, isChecked bool) error {
	err := a.AnnotationsRepo.UpdateIsCheckedAdminQuestionAnnotations(ctx, id, isChecked)
	if err != nil {
		return err
	}
	return nil
}

func (a *AnnotationsApp) MarkQuestionAnnotations(ctx echo.Context, ids []int64, mark bool) error {
	err := a.AnnotationsRepo.MarkQuestionAnnotations(ctx, ids, mark)
	if err != nil {
		return err
	}
	return nil
}
