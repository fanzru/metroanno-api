package request

type ReqAddDocument struct {
	SubjectId            int64  `json:"subject_id"`
	LearningOutcome      string `json:"learning_outcome"`
	TextDocument         string `json:"text_document"`
	MinNumberOfQuestions int64  `json:"min_number_of_questions"`
}

type ReqEditDocument struct {
	SubjectId            int64  `json:"subject_id"`
	LearningOutcome      string `json:"learning_outcome"`
	TextDocument         string `json:"text_document"`
	MinNumberOfQuestions int64  `json:"min_number_of_questions"`
}
