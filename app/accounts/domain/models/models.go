package models

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type User struct {
	Id                        int64     `json:"id" gorm:"id"`
	Type                      uint64    `json:"type" gorm:"type"`
	IsDocumentAnnotator       bool      `json:"is_document_annotator" gorm:"is_document_annotator"`
	IsQuestionAnnotator       bool      `json:"is_question_annotator" gorm:"is_question_annotator"`
	Username                  string    `json:"username" gorm:"username"`
	Password                  string    `json:"password" gorm:"password"`
	Contact                   string    `json:"contact" gorm:"contact"`
	Age                       int64     `json:"age" gorm:"age"`
	NumberOfDocumentAdded     int64     `json:"number_of_document_added" gorm:"number_of_document_added"`
	NumberOfQuestionAnnotated int64     `json:"number_of_question_annotated" gorm:"number_of_question_annotated"`
	Status                    string    `json:"status" gorm:"status"`
	CurrentDocumentID         int64     `json:"current_document_id" gorm:"current_document_id"`
	CreatedAt                 time.Time `json:"created_at" gorm:"created_at"`
	DeletedAt                 null.Time `json:"deleted_at" gorm:"deleted_at"`
}

type UsersSubjects struct {
	UserId    int64 `json:"user_id" gorm:"user_id"`
	SubjectId int64 `json:"subject_id" gorm:"subject_id"`
}

type UserWithoutPassword struct {
	Id                        int64     `json:"id" gorm:"id"`
	Type                      uint64    `json:"type" gorm:"type"`
	IsDocumentAnnotator       bool      `json:"is_document_annotator" gorm:"is_document_annotator"`
	IsQuestionAnnotator       bool      `json:"is_question_annotator" gorm:"is_question_annotator"`
	Username                  string    `json:"username" gorm:"username"`
	Contact                   string    `json:"contact" gorm:"contact"`
	Age                       int64     `json:"age" gorm:"age"`
	NumberOfDocumentAdded     int64     `json:"number_of_document_added" gorm:"number_of_document_added"`
	NumberOfQuestionAnnotated int64     `json:"number_of_question_annotated" gorm:"number_of_question_annotated"`
	Status                    string    `json:"status" gorm:"status"`
	CurrentDocumentID         int64     `json:"current_document_id" gorm:"current_document_id"`
	CreatedAt                 time.Time `json:"created_at" gorm:"created_at"`
	DeletedAt                 null.Time `json:"deleted_at" gorm:"deleted_at"`
	TotalQuestionCheckedAdmin int64     `json:"total_question_checked_admin"`
	TotalQuestion             int64     `json:"total_question_non_checked_admin"`
}
