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

func (w chiWrapper) HandleFunc(p string, f func(http.ResponseWriter, *http.Request)) {
	w.r.HandleFunc(p, f)
}

func main() {
	var _ handler = http.NewServeMux()
	var _ handler = chiWrapper{r: chi.NewMux()}
	var _ handler = chi.NewMux()
}
