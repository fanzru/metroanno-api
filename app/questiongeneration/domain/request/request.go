package request

import (
	"fmt"
	"metroanno-api/infrastructure/config"
)

type ReqGenerateQuestion struct {
	ReadingMaterial string   `json:"reading_material" validate:"required"`
	Topic           string   `json:"topic"`
	Random          string   `json:"random"`
	Bloom           []string `json:"bloom"`
	Graesser        []string `json:"graesser"`
	QuestionCount   uint64   `json:"question_count" validate:"required"`
}

type ReqTextDetection struct {
	ReadingMaterial string `json:"reading_material" validate:"required"`
}

type ReqSaveQuestion struct {
	Difficulty string `json:"difficulty" validate:"required"`
	SourceText string `json:"source_text" validate:"required"`
	Topic      string `json:"topic" validate:"required" `
	Random     string `json:"random"`
	Bloom      string `json:"bloom"`
	Graesser   string `json:"graesser"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	Type       string `json:"type"`
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
	topic := ""
	if r.Topic == "" {
		topic = fmt.Sprintf("- topic dari text yang saya berikan adalah %s", r.Topic)
	}
	instructions := "- buatkan juga tipe dari setiap pertanyaan yang di generated berdasarkan "

	bloom := ""
	if len(r.Bloom) > 0 {
		bloom = fmt.Sprintf("- buatkan response pada kolom tipe untuk setiap pertanyaan yang kamu generate, tipe bloom yang boleh digenerate adalah  %v", r.Bloom)
		instructions = fmt.Sprintf("%v, bloom : (%v)", instructions, r.Bloom)
	}

	graesser := ""
	if len(r.Graesser) > 0 {
		graesser = fmt.Sprintf("- buatkan response pada kolom tipe untuk setiap pertanyaan yang kamu generate, tipe graesser yang boleh digenerate adalah  %v", r.Graesser)
		instructions = fmt.Sprintf("%v, graesser : (%v)", instructions, r.Graesser)
	}

	random := ""
	if r.Random != "" {
		random = fmt.Sprintf("- buatkan pada kolom tipe setiap pertanyaan yang kamu generate ini berlevel berikut %v", r.Random)
		instructions = fmt.Sprintf("%v, level : (%v)", instructions, r.Random)
	}

	return fmt.Sprintf(`
		PLEASE RETURN RESPONSE ARRAY of OBJECTS!!

		context:
		- saya memiliki sebuah material text dibawah ini
		- buatkan pertanyaan berjumlah %v

		material text:
		"""
		%v
		"""
		
		instruksi: 
		 - buatkan pertanyaan dengan jawabannya dan berikan juga referensi text yang kamu gunakan untuk membuat pertanyaan
		 - referensi text merupakan kalimat yang menjadi tempat kamu membuat soal dan jawaban
		 %v
		 %v
		 %v
		 %v

		`,
		r.QuestionCount,
		r.ReadingMaterial,
		topic,
		random,
		graesser,
		bloom,
		// instructions,
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
					"type": map[string]interface{}{
						"type":        "string",
						"description": "Make type from random / graesser and bloom enum",
					},
				},
				"required": []string{
					"id",
					"source_text",
					"question",
					"answer",
					"type",
				},
			},
		},
	}
	return result
}

func (r *ReqTextDetection) BuildContextForChatGpt() string {
	return fmt.Sprintf(`
		text:
		"""
		%v
		"""

		notes:
		- tolong deteksi topik dari text yang saya kirimkan!
		- responsenya hanya boleh panjangnyaa 1-10 kata!
		
		apa itu topik:
		- Dalam konteks penulisan atau pembicaraan, "topik" merujuk pada pokok atau subjek utama dari suatu teks atau pembicaraan. Topik menentukan tentang apa yang sedang dibahas atau disampaikan. Dalam sebuah teks, topik biasanya diungkapkan melalui kalimat utama atau tesis, sementara dalam pembicaraan, topik dapat diidentifikasi dari konteks dan fokus percakapan.
		- Contoh, jika Anda membaca artikel tentang perubahan iklim, topiknya adalah perubahan iklim. Jika Anda sedang berbicara dengan seseorang tentang rencana liburan Anda, topiknya adalah liburan Anda.
		
		`,
		r.ReadingMaterial,
	)
}

func (r *ReqTextDetection) BuildFunctionsCustomForChatGpt() []map[string]interface{} {
	result := []map[string]interface{}{
		{
			"name":        "text_detection",
			"description": `deteksi topik dari text yang saya kirimkan!!`,
			"parameters": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"topic": map[string]interface{}{
						"type":        "string",
						"description": "hasil deteksi text yang saya kirim topicnya tentang apa",
					},
				},
				"required": []string{
					"topic",
				},
			},
		},
	}
	return result
}
