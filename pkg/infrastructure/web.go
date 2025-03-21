package infrastructure

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request) error

func Handle(handler Handler) http.HandlerFunc {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	return h
}
