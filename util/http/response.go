package http

import (
	"fmt"
	h "net/http"
)

func OK(w h.ResponseWriter, body []byte) {
	w.WriteHeader(h.StatusOK)
	w.Write(body)
}

func BadRequest(w h.ResponseWriter, message string) {
	w.WriteHeader(h.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, message)))
}

func InternalServerError(w h.ResponseWriter) {
	w.WriteHeader(h.StatusInternalServerError)
	w.Write([]byte("something went wrong"))
	w.Write([]byte(`{"error": "something went wrong"}`))
}
