package repo

import (
	"metroanno-api/app/annotation/domain/models"

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
