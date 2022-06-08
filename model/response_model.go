package model

import (
	"encoding/json"
	"net/http"
	"fidtest_golang/entity"
)
func ModelResponseText(w http.ResponseWriter, code int,text string) {
	var responseText entity.ResponseText
	responseText.StatusCode=code
	responseText.Message=text
	ResponseWithJson(w, code, responseText)
}
func ModelResponseData(w http.ResponseWriter, code int,text string,payload interface{}) {
	var responseData entity.ResponseData
	responseData.StatusCode=code
	responseData.Message= text
	responseData.Data=payload
	ResponseWithJson(w, code, responseData)
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response,_:=json.Marshal(payload)
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(code)
	w.Write(response)
}