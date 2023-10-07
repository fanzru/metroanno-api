package request

import (
	"fmt"
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
	Topic           string `json:"topic" validate:"required"`
	Random          string `json:"random"`
	Bloom           string `json:"bloom"`
	Graesser        string `json:"graesser"`
}

type ReqSaveQuestions struct {
	SaveQuestions []ReqSaveQuestion `json:"save_questions" validate:"required"`
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
					"referensi_text": map[string]interface{}{
						"type":        "string",
						"description": "text reference from the given text material",
					},
					"soal": map[string]interface{}{
						"type":        "string",
						"description": "Make questions from the material provided",
					},
					"jawaban": map[string]interface{}{
						"type":        "string",
						"description": "Make an answer from the material provided",
					},
				},
				"required": []string{
					"id",
					"referensi_text",
					"soal",
					"jawaban",
				},
			},
		},
	}
	return result
}
