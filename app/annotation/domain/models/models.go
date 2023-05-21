package models

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	Id                                     int64          `json:"id" gorm:"id"`
	SubjectId                              int64          `json:"subject_id" gorm:"subject_id"`
	LearningOutcome                        string         `json:"learning_outcome" gorm:"learning_outcome"`
	TextDocument                           string         `json:"text_document" gorm:"text_document"`
	MinNumberOfAnnotators                  int64          `json:"min_number_of_annotators" gorm:"min_number_of_annotators"`
	CurrentNumberOfAnnotatorsAssigned      int64          `json:"current_number_of_annotators_assigned" gorm:"current_number_of_annotators_assigned"`
	MinNumberOfQuestionsPerAnnotator       int64          `json:"min_number_of_questions_per_annotator" gorm:"min_number_of_questions_per_annotator"`
	CurrentTotalNumberOfQuestionsAnnotated int64          `json:"current_total_number_of_questions_annotated" gorm:"current_total_number_of_questions_annotated"`
	DoneNumberOfAnnotators                 int64          `json:"done_number_of_annotators"`
	IsApproved                             bool           `json:"is_approved" gorm:"is_approved"`
	CreatedAt                              time.Time      `json:"created_at" gorm:"created_at"`
	DeletedAt                              gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
	CreatedByUserId                        int64          `json:"created_by_user_id" gorm:"created_by_user_id"`

	QuestionAnnotations []*QuestionAnnotation `json:"question_annotations" gorm:"foreignKey:document_id;"`
}

type QuestionType struct {
	Id           int64          `json:"id" gorm:"id"`
	QuestionType string         `json:"question_type" gorm:"question_type"`
	Description  string         `json:"description" gorm:"question_type"`
	CreatedAt    time.Time      `json:"created_at" gorm:"created_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
}

type Feedback struct {
	Id           int64          `json:"id" gorm:"id"`
	UserID       int64          `json:"user_id" gorm:"user_id"`
	DocumentID   int64          `json:"document_id" gorm:"document_id"`
	FeedbackText string         `json:"feedback_text" gorm:"feedback_text"`
	CreatedAt    time.Time      `json:"created_at" gorm:"created_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
}

type QuestionAnnotation struct {
	Id                     int64          `json:"id" gorm:"id"`
	UserID                 int64          `json:"user_id" gorm:"user_id"`
	DocumentID             int64          `json:"document_id" gorm:"document_id"`
	IsLearningOutcomeShown bool           `json:"is_learning_outcome_shown" gorm:"is_learning_outcome_shown"`
	QuestionOrder          int64          `json:"question_order" gorm:"question_order"`
	QuestionTypeId         int64          `json:"question_type_id" gorm:"question_type_id"`
	Keywords               string         `json:"keywords" gorm:"keywords"`
	QuestionText           string         `json:"question_text" gorm:"question_text"`
	AnswerText             string         `json:"answer_text" gorm:"answer_text"`
	TimeDuration           int64          `json:"time_duration" gorm:"time_duration"`
	IsCheckedAdmin         bool           `json:"is_checked_admin" gorm:"is_checked_admin"`
	CreatedAt              time.Time      `json:"created_at" gorm:"created_at"`
	DeletedAt              gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`

	//Document Document `json:"-"`
}

type DoneDocumentUser struct {
	Id         int64 `json:"id" gorm:"id"`
	UserID     int64 `json:"user_id" gorm:"user_id"`
	DocumentID int64 `json:"document_id" gorm:"document_id"`
	Done       bool  `json:"done" gorm:"done"`
}

type Subject struct {
	Id          int64  `json:"id" gorm:"id"`
	SubjectText string `json:"subject_text"`
}
