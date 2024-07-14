package respond

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AxelTahmid/golang-starter/pkg/message"
)

func (rw *Writer) WithJson(payload interface{}) {
	rw.w.Header().Set("Content-Type", "application/json")

	if payload == nil {
		return
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		jsonErr(rw.w, http.StatusInternalServerError, message.ErrInternalError)
		return
	}

	if string(data) == "null" {
		_, _ = rw.w.Write([]byte("[]"))
		return
	}

	_, err = rw.w.Write(data)
	if err != nil {
		log.Println(err)
		jsonErr(rw.w, http.StatusInternalServerError, message.ErrInternalError)
		return
	}
}

func jsonErr(w http.ResponseWriter, statusCode int, message error) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(statusCode)

	var p map[string]string
	if message == nil {
		write(w, nil)
		return
	}

	p = map[string]string{
		"message": message.Error(),
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}

	if string(data) == "null" {
		return
	}

	write(w, data)
}
