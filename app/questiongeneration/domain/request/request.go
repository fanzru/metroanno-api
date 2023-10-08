package request

import (
	"fmt"
	"metroanno-api/infrastructure/config"
)

type ReqGenerateQuestion struct {
	ReadingMaterial string `json:"reading_material" validate:"required"`
	Topic           string `json:"topic" validate:"required"`
	Random          string `json:"random"`
	Bloom           string `json:"bloom"`
	Graesser        string `json:"graesser"`
	QuestionCount   uint64 `json:"question_count" validate:"required"`
}

type ReqSaveQuestion struct {
	Difficulty      string `json:"difficulty" validate:"required"`
	ReadingMaterial string `json:"reading_material" validate:"required"`
	Topic           string `json:"topic" validate:"required" `
	Random          string `json:"random"`
	Bloom           string `json:"bloom"`
	Graesser        string `json:"graesser"`
	Question        string `json:"question"`
	Answer          string `json:"answer"`
}

type ReqSaveQuestions struct {
	ReadingMaterial string            `json:"reading_material" validate:"required"`
	SaveQuestions   []ReqSaveQuestion `json:"save_questions" validate:"required"`
}

func (r *ReqSaveQuestion) IsEnumValueRandom(value map[string]bool) bool {
	return value[r.Random]
}

func (r *ReqSaveQuestion) IsEnumValueBloom(value map[string]bool) bool {
	return value[r.Bloom]
}

func (r *ReqSaveQuestion) IsEnumValueGraesser(value map[string]bool) bool {
	return value[r.Graesser]
}

func (r *ReqSaveQuestions) IsValidEnum(cfg config.Config) bool {
	randomValue := cfg.MapTrueValueRandom()
	bloomValue := cfg.MapTrueValueBloom()
	graesserValue := cfg.MapTrueValueGraesser()

	for _, q := range r.SaveQuestions {
		if !q.IsEnumValueBloom(bloomValue) {
			return false
		}
		if !q.IsEnumValueGraesser(graesserValue) {
			return false
		}
		if !q.IsEnumValueRandom(randomValue) {
			return false
		}
	}
	return true
}

func (r *ReqGenerateQuestion) BuildContextForChatGpt() string {
	return fmt.Sprintf(`
		context:
		- Saya memiliki sebuah material text dibawah ini
		
		topic:
		"""
		%v
		"""
		material text:
		"""
		%v
		"""

		instruksi: 
		 - buatkan pertanyaan dengan jawabannya dan berikan juga referensi text yang kamu gunakan untuk membuat pertanyaan
		 - referensi text merupakan kalimat yang menjadi tempat kamu membuat soal dan jawaban
		 - pertanyaan akan berjumlah %v
		 - tolong berikan response dalam array!
		`,
		r.Topic,
		r.ReadingMaterial,
		r.QuestionCount,
	)
}

func (r *ReqGenerateQuestion) BuildFunctionsCustomForChatGpt() []map[string]interface{} {
	result := []map[string]interface{}{
		{
			"name":        "list_question",
			"description": fmt.Sprintf(`buatkan %v pertanyaan dengan jawabannya dan berikan juga referensi text yang kamu gunakan untuk membuat pertanyaan, referensi text merupakan kalimat yang menjadi tempat kamu membuat soal dan jawaban`, r.QuestionCount),
			"parameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id": map[string]interface{}{
						"type":        "string",
						"description": "soal number",
					},
					"source_text": map[string]interface{}{
						"type":        "string",
						"description": "text reference from the given text material",
					},
					"question": map[string]interface{}{
						"type":        "string",
						"description": "Make questions from the material provided",
					},
					"answer": map[string]interface{}{
						"type":        "string",
						"description": "Make an answer from the material provided",
					},
				},
				"required": []string{
					"id",
					"source_text",
					"question",
					"answer",
				},
			},
		},
	}
	return result
}
