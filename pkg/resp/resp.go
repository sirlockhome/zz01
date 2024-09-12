package resp

import (
	"encoding/json"
	"net/http"
)

type Response struct {
}

func WriteJSONData(w http.ResponseWriter, data interface{}, status int) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataBytes)
}

func WriteJSONMessage(w http.ResponseWriter, mess string, status int) {
	resp := struct {
		Message string `json:"message"`
	}{
		Message: mess,
	}

	dataBytes, err := json.Marshal(resp)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataBytes)
}
