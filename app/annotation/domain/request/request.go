package request

type ReqAddDocument struct {
	SubjectId             int64  `json:"subject_id" validate:"required"`
	LearningOutcome       string `json:"learning_outcome"`
	TextDocument          string `json:"text_document" validate:"required"`
	MinNumberOfQuestions  uint64 `json:"min_number_of_questions" validate:"required"`
	MinNumberOfAnnotators uint64 `json:"min_number_of_annotators" validate:"required"`
}

type ReqEditDocument struct {
	DocumentId            int64  `json:"document_id" validate:"required"`
	SubjectId             int64  `json:"subject_id" validate:"required"`
	LearningOutcome       string `json:"learning_outcome"`
	TextDocument          string `json:"text_document" validate:"required"`
	MinNumberOfQuestions  int64  `json:"min_number_of_questions" validate:"required"`
	MinNumberOfAnnotators int64  `json:"min_number_of_annotators" validate:"required"`
}

type ReqAddQuestionType struct {
	QuestionType string `json:"question_type" validate:"required"`
	Description  string `json:"description" validate:"required"`
}

type ReqCreateFeedback struct {
	DocumentID   int64  `json:"document_id" validate:"required"`
	FeedbackText string `json:"feedback_text" validate:"required"`
}

type ReqCreateQuestionAnnoatations struct {
	DocumentId             int64                 `json:"document_id" validate:"required"`
	IsLearningOutcomeShown bool                  `json:"is_learning_outcome_shown" validate:"required"`
	TimeDuration           int64                 `json:"time_duration" validate:"required"`
	QuestionAnnoatations   []QuestionAnnoatation `json:"question_annotations" validate:"required"`
}

type QuestionAnnoatation struct {
	QuestionTypeId int64  `json:"question_type_id" validate:"required"`
	Keywords       string `json:"keywords" validate:"required"`
	QuestionText   string `json:"question_text" validate:"required"`
	AnswerText     string `json:"answer_text" validate:"required"`
}
