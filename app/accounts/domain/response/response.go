package response

import (
	"metroanno-api/app/accounts/domain/models"
	modelsannotation "metroanno-api/app/annotation/domain/models"
)

type UserLoginRes struct {
	AccessToken string `json:"access_token"`
}

type ProfileRes struct {
	User     UserRes                    `json:"user"`
	Document *modelsannotation.Document `json:"document"`
}
type UserRes struct {
	Id                        int64  `json:"id" gorm:"id"`
	Type                      uint64 `json:"type" gorm:"type"`
	IsDocumentAnnotator       bool   `json:"is_document_annotator" gorm:"is_document_annotator"`
	IsQuestionAnnotator       bool   `json:"is_question_annotator" gorm:"is_question_annotator"`
	Username                  string `json:"username" gorm:"username"`
	Contact                   string `json:"contact" gorm:"contact"`
	Age                       int64  `json:"age" gorm:"age"`
	NumberOfDocumentAdded     int64  `json:"number_of_document_added" gorm:"number_of_document_added"`
	NumberOfQuestionAnnotated int64  `json:"number_of_question_annotated" gorm:"number_of_question_annotated"`
	Status                    string `json:"status" gorm:"status"`
	CurrentDocumentID         int64  `json:"current_document_id" gorm:"current_document_id"`
}

type Pagination struct {
	Page  int64                         `json:"page"`
	Limit int64                         `json:"limit"`
	Prev  int64                         `json:"prev"`
	Next  int64                         `json:"next"`
	Start int64                         `json:"start"`
	End   int64                         `json:"end"`
	Data  []*models.UserWithoutPassword `json:"data"`
}
