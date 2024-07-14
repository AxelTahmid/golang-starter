package respond

import (
	"encoding/json"
	"log"
	"net/http"
)

func write(w http.ResponseWriter, data []byte) {
	_, err := w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

func (rw *Writer) WithErrs(errors []string) {
	rw.w.Header().Set("Content-Type", "application/problem+json")

	if errors == nil {
		write(rw.w, nil)
		return
	}

	p := map[string][]string{
		"message": errors,
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}

	if string(data) == "null" {
		return
	}

	write(rw.w, data)
}

func (rw *Writer) WithErr(message error) {
	rw.w.Header().Set("Content-Type", "application/problem+json")

	var p map[string]string
	if message == nil {
		write(rw.w, nil)
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

	write(rw.w, data)
}
