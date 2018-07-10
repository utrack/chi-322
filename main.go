package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

type handler interface {
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}

type chiWrapper struct {
	r chi.Router
}

// HandleFunc is practically the same, but instead of alias http.HandlerFunc it accepts
// its expanded version:
// type HandlerFunc func(ResponseWriter, *Request)
func (w chiWrapper) HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) {
	// func(ResponseWriter,*Request) == http.HandlerFunc, so no conversion needed
	w.r.HandleFunc(p, f)
}

func main() {
	// std mux implements the interface
	var _ handler = http.NewServeMux()
	// wrapper implements the interface
	var _ handler = chiWrapper{r: chi.NewMux()}
	// chi does not
	var _ handler = chi.NewMux()
}
