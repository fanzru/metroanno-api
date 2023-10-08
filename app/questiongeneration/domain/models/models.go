package models

import (
	"metroanno-api/infrastructure/config"
	"time"
)

type QuestionsHistory struct {
	ID              int        `json:"id"`
	Difficulty      string     `json:"difficulty"`
	ReadingMaterial string     `json:"reading_material"`
	Topic           string     `json:"topic"`
	Random          string     `json:"random"`
	Bloom           string     `json:"bloom"`
	Graesser        string     `json:"graesser"`
	CreatedAt       time.Time  `json:"created_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	UserID          int64      `json:"user_id"`
}

type QuestionType struct {
	Random   []config.QuestionType `json:"random"`
	Bloom    []config.QuestionType `json:"bloom"`
	Graesser []config.QuestionType `json:"graesser"`
}
