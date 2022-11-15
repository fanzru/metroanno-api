package models

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type Document struct {
	Id                                     int64     `json:"id" gorm:"id"`
	SubjectId                              int64     `json:"subject_id" gorm:"subject_id"`
	LearningOutcome                        string    `json:"learning_outcome" gorm:"learning_outcome"`
	TextDocument                           string    `json:"text_document" gorm:"text_document"`
	MinNumberOfAnnotators                  int64     `json:"min_number_of_annotators" gorm:"min_number_of_annotators"`
	CurrentNumberOfAnnotatorsAssigned      int64     `json:"current_number_of_annotators_assigned" gorm:"current_number_of_annotators_assigned"`
	MinNumberOfQuestionsPerAnnotator       int64     `json:"min_number_of_questions_per_annotator" gorm:"min_number_of_questions_per_annotator"`
	CurrentTotalNumberOfQuestionsAnnotated int64     `json:"current_total_number_of_questions_annotated" gorm:"current_total_number_of_questions_annotated"`
	CreatedAt                              time.Time `json:"created_at" gorm:"created_at"`
	DeletedAt                              null.Time `json:"deleted_at" gorm:"deleted_at"`
}

type QuestionType struct {
	Id           int64     `json:"id" gorm:"id"`
	QuestionType string    `json:"question_type" gorm:"question_type"`
	Description  string    `json:"description" gorm:"question_type"`
	CreatedAt    time.Time `json:"created_at" gorm:"created_at"`
	DeletedAt    null.Time `json:"deleted_at" gorm:"deleted_at"`
}
