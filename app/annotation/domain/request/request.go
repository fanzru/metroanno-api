package request

type ReqAddDocument struct {
	SubjectId            int64  `json:"subject_id" validate:"required"`
	LearningOutcome      string `json:"learning_outcome"`
	TextDocument         string `json:"text_document" validate:"required"`
	MinNumberOfQuestions uint64 `json:"min_number_of_questions" validate:"required"`
}

type ReqEditDocument struct {
	SubjectId            int64  `json:"subject_id" validate:"required"`
	LearningOutcome      string `json:"learning_outcome"`
	TextDocument         string `json:"text_document" validate:"required"`
	MinNumberOfQuestions int64  `json:"min_number_of_questions" validate:"required"`
}
