package repo

import (
	accountsModel "metroanno-api/app/accounts/domain/models"
	"metroanno-api/app/annotation/domain/models"
	"metroanno-api/app/annotation/domain/request"
	"metroanno-api/app/annotation/domain/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (a *AnnotationsRepo) CreateTheory(ctx echo.Context, document models.Document) error {
	err := a.MySQL.DB.Table("documents").Create(&document).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *AnnotationsRepo) GetDocumentsById(ctx echo.Context, documentId int64) (*models.Document, error) {
	document := &models.Document{}
	err := a.MySQL.DB.Table("documents").Where("id = ?", documentId).First(&document).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (a *AnnotationsRepo) GetAllDocuments(ctx echo.Context) ([]models.Document, error) {
	document := []models.Document{}
	err := a.MySQL.DB.Table("documents").Find(&document).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (a *AnnotationsRepo) GetAllDocumentsWithWhere(ctx echo.Context) ([]models.Document, error) {
	document := []models.Document{}
	err := a.MySQL.DB.Table("documents").Where("is_approved = ?", true).Find(&document).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (a *AnnotationsRepo) UpdateDocumentsById(ctx echo.Context, param request.ReqEditDocument) (*models.Document, error) {
	documentfind := &models.Document{}
	err := a.MySQL.DB.Table("documents").Where("id=?", param.DocumentId).First(documentfind).Error
	if err != nil {
		return nil, err
	}

	document := &models.Document{
		Id:                               param.DocumentId,
		SubjectId:                        param.SubjectId,
		LearningOutcome:                  param.LearningOutcome,
		TextDocument:                     param.TextDocument,
		MinNumberOfQuestionsPerAnnotator: param.MinNumberOfQuestions,
		MinNumberOfAnnotators:            param.MinNumberOfAnnotators,
	}
	err = a.MySQL.DB.Table("documents").Where("id = ?", param.DocumentId).Updates(document).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (a *AnnotationsRepo) UpdateAddDoneDocumentsById(ctx echo.Context, param request.ReqEditDocument, tx *gorm.DB) (*gorm.DB, error) {
	documentfind := &models.Document{}
	err := tx.Table("documents").Where("id=?", param.DocumentId).First(documentfind).Error
	if err != nil {
		return nil, err
	}

	document := &models.Document{
		Id:                     param.DocumentId,
		DoneNumberOfAnnotators: documentfind.DoneNumberOfAnnotators + 1,
	}

	err = tx.Table("documents").Where("id = ?", param.DocumentId).Updates(document).Error
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (a *AnnotationsRepo) GetDocumentIdByUserId(ctx echo.Context) (int64, error) {
	modelAccount := &accountsModel.User{}
	err := a.MySQL.DB.Table(TableUsers).Where("id = ?", ctx.Get("user_id")).Find(&modelAccount).Error
	if err != nil {
		return 0, err
	}
	if modelAccount.CurrentDocumentID == 0 {
		return 0, ErrDocumentNotFound
	}

	modelDocument := &models.Document{}
	err = a.MySQL.DB.Table(TableDocuments).Where("id = ?", modelAccount.CurrentDocumentID).First(&modelDocument).Error
	if err != nil {
		if err.Error() == "record not found" {
			return 0, ErrDocumentIsDeleted
		}
		return 0, err
	}

	return modelAccount.CurrentDocumentID, nil
}

func (a *AnnotationsRepo) DeleteDocumentsByID(ctx echo.Context, id int64) (*models.Document, error) {
	document := &models.Document{
		Id: id,
	}

	documentGet := &models.Document{}
	err := a.MySQL.DB.Table("documents").First(documentGet, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	err = a.MySQL.DB.Table("documents").Delete(&document, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return document, nil
}

func (a *AnnotationsRepo) GetAllDocumentDoneUser(ctx echo.Context, userID int64) (*[]models.DoneDocumentUser, error) {
	arrObj := &[]models.DoneDocumentUser{}
	err := a.MySQL.DB.Table("done_document_user").Where("user_id = ?", userID).Find(arrObj).Error
	if err != nil {
		return nil, err
	}

	return arrObj, nil
}

func (i *AnnotationsRepo) GetAllDocumentsAdmin(ctx echo.Context, pageNumber int64) (response.Pagination, error) {
	resp := response.Pagination{}
	var totalRows int64

	err := i.MySQL.DB.Model(&models.Document{}).Count(&totalRows).Error
	if err != nil {
		return resp, err
	}

	// calculate page dan offset
	// contoh halaman yang diminta
	pageSize := 10 // contoh ukuran halaman
	offset := (int(pageNumber) - 1) * pageSize

	documents := []models.Document{}
	result := i.MySQL.DB.Set("gorm:auto_preload", true).Model(&models.Document{}).Offset(offset).Limit(pageSize).Preload("QuestionAnnotations").Find(&documents)
	if result.Error != nil {
		return resp, result.Error
	}

	var prevPage, nextPage int64
	if pageNumber > 1 {
		prevPage = pageNumber - 1
	}
	if int(offset)+len(documents) < int(totalRows) {
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
	resp = response.Pagination{
		Page:  pageNumber,
		Limit: int64(pageSize),
		Prev:  prevPage,
		Next:  nextPage,
		Start: start,
		End:   totalPages,
		Data:  documents,
	}

	return resp, nil
}

func (i *AnnotationsRepo) UpdateIsAprrovedDocument(ctx echo.Context, documentID int64, isApproved bool) error {
	err := i.MySQL.DB.Table("documents").Where("id = ?", documentID).Update("is_approved", isApproved).Error
	if err != nil {
		return err
	}
	return nil
}
