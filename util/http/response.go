package http

import h "net/http"

func OK(w h.ResponseWriter, body []byte) {
	w.WriteHeader(h.StatusOK)
	w.Write(body)
}

func BadRequest(w h.ResponseWriter, message string) {
	w.WriteHeader(h.StatusBadRequest)
	w.Write([]byte(message))
}

func InternalServerError(w h.ResponseWriter) {
	w.WriteHeader(h.StatusInternalServerError)
	w.Write([]byte("something went wrong"))
}
