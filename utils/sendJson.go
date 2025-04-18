package utils

import (
	"cadUser/model"
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, resp model.Response, statusCode int) {
	logger := StructLogger()

	data, err := json.Marshal(resp)
	if err != nil {

		logger.Error("error to marshal response", "error", err)
		SendJSON(w, model.Response{Error: "internal server error"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	if _, err := w.Write(data); err != nil {
		logger.Error("error to write response", "error", err)
		return
	}
}
