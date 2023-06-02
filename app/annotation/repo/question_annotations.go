package repo

import (
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (a *AnnotationsRepo) CountQuestionAnnotationByDocumentID(ctx echo.Context, documentID int64) (int64, error) {
	var count int64
	err := a.MySQL.DB.Table(TableQuestionAnnotations).Where("document_id = ?", documentID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *AnnotationsRepo) GetQuestionAnnotationByID(ctx echo.Context, ID int64) (*models.QuestionAnnotation, error) {
	models := &models.QuestionAnnotation{}
	err := a.MySQL.DB.Table(TableQuestionAnnotations).Where("id = ?", ID).First(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (a *AnnotationsRepo) GetQuestionAnnotationByDocumentID(ctx echo.Context, documentID int64) (*[]models.QuestionAnnotation, error) {
	models := &[]models.QuestionAnnotation{}
	err := a.MySQL.DB.Table(TableQuestionAnnotations).Where("document_id = ?", documentID).Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (a *AnnotationsRepo) BulkInsertQuestionAnnotations(ctx echo.Context, arrQuestionAnnotations []models.QuestionAnnotation) (*[]models.QuestionAnnotation, *gorm.DB, error) {
	tx := a.MySQL.DB.Begin()

	err := tx.Table(TableQuestionAnnotations).CreateInBatches(&arrQuestionAnnotations, 10).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	return &arrQuestionAnnotations, tx, nil
}

func (i *AnnotationsRepo) UpdateIsCheckedAdminQuestionAnnotations(ctx echo.Context, id int64, isChecked bool) error {
	err := i.MySQL.DB.Table("question_annotations").Where("id = ?", id).Update("is_checked_admin", isChecked).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *AnnotationsRepo) MarkQuestionAnnotations(ctx echo.Context, ids []int64, mark bool) error {
	err := i.MySQL.DB.Table("question_annotations").Where("id IN ?", ids).Update("mark", mark).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *AnnotationsRepo) GetAllQuestionAnnotationsUser(ctx echo.Context, userId, pageNumber int64) (response.PaginationQA, error) {
	resp := response.PaginationQA{}
	var totalRows int64

	err := i.MySQL.DB.Model(&models.QuestionAnnotation{}).Where("user_id=?", userId).Count(&totalRows).Error
	if err != nil {
		return resp, err
	}

	// calculate page dan offset
	// contoh halaman yang diminta
	pageSize := 10 // contoh ukuran halaman
	offset := (int(pageNumber) - 1) * pageSize

	qa := []*models.QuestionAnnotation{}
	result := i.MySQL.DB.Set("gorm:auto_preload", true).Model(&models.QuestionAnnotation{}).Where("user_id = ?  ", userId).Order("id desc").Offset(offset).Limit(pageSize).Find(&qa)
	if result.Error != nil {
		return resp, result.Error
	}

	var prevPage, nextPage int64
	if pageNumber > 1 {
		prevPage = pageNumber - 1
	}
	if int(offset)+len(qa) < int(totalRows) {
		nextPage = pageNumber + 1
	}

	var totalPages int64
	if totalRows%int64(pageSize) == 0 {
		totalPages = totalRows / int64(pageSize)
	} else {
		totalPages = (totalRows / int64(pageSize)) + 1
	}

	var start int64 = 1
	if totalRows == 0 {
		start = 0
	}
	resp = response.PaginationQA{
		Page:  pageNumber,
		Limit: int64(pageSize),
		Prev:  prevPage,
		Next:  nextPage,
		Start: start,
		End:   totalPages,
		Data:  qa,
	}

	return resp, nil
}
