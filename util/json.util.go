package util

import (
	"encoding/json"
	"io"
	"log"
	"my-go-app/dto"
	"my-go-app/model/monad/result"
	"net/http"
)

// GetRequestはジェネリックな関数で、リクエストボディをデコードします
func GetRequest[T any](r *http.Request) result.Result[T] {
	// リクエストボディを読み取る
	bodyBytes, readRequestBodyError := io.ReadAll(r.Body)
	if readRequestBodyError != nil {
		// log.Println("failed to read body", readRequestBodyError)
		return result.OfErr[T](readRequestBodyError)
	}

	// JSONを指定された型にデコード
	var req T
	jsonUnmarshalError := json.Unmarshal(bodyBytes, &req)
	if jsonUnmarshalError != nil {
		log.Println("failed to decode from json", jsonUnmarshalError)
		// T型のゼロ値を返し、エラーを返す
		return result.OfErr[T](jsonUnmarshalError)
	}

	// 成功した場合、デコードされた構造体を返す
	return result.OfOk[T](req)
}

func GetResponse[T any](r *http.Response) result.Result[T] {
	// リクエストボディを読み取る
	bodyBytes, readRequestBodyError := io.ReadAll(r.Body)
	if readRequestBodyError != nil {
		return result.OfErr[T](readRequestBodyError)
	}

	// JSONを指定された型にデコード
	var req T
	jsonUnmarshalError := json.Unmarshal(bodyBytes, &req)
	if jsonUnmarshalError != nil {
		log.Println("failed to decode from json", jsonUnmarshalError)
		// T型のゼロ値を返し、エラーを返す
		return result.OfErr[T](jsonUnmarshalError)
	}

	// 成功した場合、デコードされた構造体を返す
	return result.OfOk[T](req)
}

func JsonDecode[T any](req dto.RequestTest) T {
	var request T
	err := json.Unmarshal([]byte(req.Json), &request)
	if err != nil {
		log.Println("ERR:", err)
	}
	return request
}

func JsonUnmarshal[T any](byte []byte) result.Result[T] {
	var req T
	err := json.Unmarshal(byte, &req)
	if err == nil {
		return result.OfOk(req)
	}
	return result.OfErr[T](err)
}
