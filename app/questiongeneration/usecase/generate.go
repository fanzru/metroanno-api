package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"metroanno-api/app/questiongeneration/domain/request"
	"metroanno-api/app/questiongeneration/domain/response"
	"metroanno-api/pkg/debugger"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *QuestionGenerationApp) GenerateQuestion(ctx echo.Context, params request.ReqGenerateQuestion) ([]response.JSONResponse, error) {

	// construct command to chat gpt
	contextMessage := params.BuildContextForChatGpt()

	// call chat gpt api endpoint
	functionContext := params.BuildFunctionsCustomForChatGpt()

	// Data untuk dikirimkan dalam permintaan POST
	data := map[string]interface{}{
		"model": a.Cfg.ModelName,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": contextMessage,
			},
		},
		"function_call":     "auto",
		"functions":         functionContext,
		"temperature":       0,
		"max_tokens":        256,
		"top_p":             1,
		"frequency_penalty": 0,
		"presence_penalty":  0,
	}

	// Mengubah data menjadi format JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return []response.JSONResponse{}, err
	}

	// Membuat permintaan POST
	req, err := http.NewRequest("POST", a.Cfg.APIUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return []response.JSONResponse{}, err
	}

	// Menambahkan header ke permintaan
	req.Header.Set("Content-Type", "application/json")

	// Ganti "YOUR_API_KEY" dengan kunci API OpenAI Anda
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", a.Cfg.APIKey))

	// Mengirimkan permintaan
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return []response.JSONResponse{}, err
	}
	defer resp.Body.Close()

	// construct response
	responseData, err := a.constructResponse(ctx, resp, params)
	if err != nil {
		fmt.Println("Error construct response:", err)
		return []response.JSONResponse{}, err
	}
	// save to history database

	// return response
	return responseData, nil
}

func (a *QuestionGenerationApp) constructResponse(ctx echo.Context, resp *http.Response, params request.ReqGenerateQuestion) ([]response.JSONResponse, error) {
	// Membaca dan mencetak respons
	var result response.ChatGPTResponse
	err := json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return []response.JSONResponse{}, nil
	}
	debugger.PrintJson(result, "ChatGPTResponse")
	// unmarshal response from chat gpt
	var responseUser []response.JSONResponse

	stringResponse := result.Choices[0].Message.FunctionCall.Arguments
	err = json.Unmarshal([]byte(stringResponse), &responseUser)
	if err != nil {

		var ru response.JSONResponse
		stringResponse := result.Choices[0].Message.FunctionCall.Arguments
		err = json.Unmarshal([]byte(stringResponse), &ru)
		if err != nil {
			fmt.Println("Error unmarshaling JSON message chat gpt:", err)
			return []response.JSONResponse{}, errors.New("error unmarshaling JSON message chat gpt")
		}
		responseUser = append(responseUser, ru)

		fmt.Println("Error unmarshaling JSON message chat gpt:", err)
		return responseUser, nil
	}
	return responseUser, nil

	// if params.QuestionCount > 1 {
	// 	stringResponse := result.Choices[0].Message.FunctionCall.Arguments
	// 	err = json.Unmarshal([]byte(stringResponse), &responseUser)
	// 	if err != nil {
	// 		fmt.Println("Error unmarshaling JSON message chat gpt:", err)
	// 		return []response.JSONResponse{}, nil
	// 	}
	// } else {
	// 	var ru response.JSONResponse
	// 	stringResponse := result.Choices[0].Message.FunctionCall.Arguments
	// 	err = json.Unmarshal([]byte(stringResponse), &ru)
	// 	if err != nil {
	// 		fmt.Println("Error unmarshaling JSON message chat gpt:", err)
	// 		return []response.JSONResponse{}, nil
	// 	}
	// 	responseUser = append(responseUser, ru)
	// }

	// return responseUser, nil
}
