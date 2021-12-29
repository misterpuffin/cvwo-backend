package utils

import (
	"net/http"
	"encoding/json"
	"server/packages/config"
)

type ErrorResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
}

/*
HTTP Response handling for errors,
Returns valid JSON with error type and response code
*/

// func Message(status bool, message string) (map[string]interface{}) {
//     return map[string]interface{} {"status" : status, "message" : message}
// }

// func Respond(w http.ResponseWriter, data map[string] interface{})  {
//     w.Header().Add("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(data)
// }

func NewErrorResponse(w http.ResponseWriter, statusCode int, response string){
	error := ErrorResponse{
		true,
		response,
	}
	w.Header().Set("Access-Control-Allow-Origin", config.Config[config.CLIENT_URL])
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&error)
	return
}

func NewJSONResponse(w http.ResponseWriter, data interface{}){
	w.Header().Set("Access-Control-Allow-Origin", config.Config[config.CLIENT_URL])
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	return
}