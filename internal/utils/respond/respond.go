package respond

import "net/http"

type Standard struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
	Meta    Meta        `json:"meta,omitempty"`
}

type Meta struct {
	Size  int `json:"size,omitempty"`
	Total int `json:"total,omitempty"`
}

type Writer struct {
	w http.ResponseWriter
}

func Write(w http.ResponseWriter) *Writer {
	return &Writer{w: w}
}

func (rw *Writer) Status(status int) *Writer {
	rw.w.WriteHeader(status)
	return rw
}
