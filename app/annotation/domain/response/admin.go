package response

import "metroanno-api/app/annotation/domain/models"

type Pagination struct {
	Page  int64              `json:"page"`
	Limit int64              `json:"limit"`
	Prev  int64              `json:"prev"`
	Next  int64              `json:"next"`
	Start int64              `json:"start"`
	End   int64              `json:"end"`
	Data  []*models.Document `json:"data"`
}

type PaginationQA struct {
	Page  int64                        `json:"page"`
	Limit int64                        `json:"limit"`
	Prev  int64                        `json:"prev"`
	Next  int64                        `json:"next"`
	Start int64                        `json:"start"`
	End   int64                        `json:"end"`
	Data  []*models.QuestionAnnotation `json:"data"`
}
