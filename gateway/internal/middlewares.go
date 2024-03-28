package internal

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

func method(m string, next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == m {
			next(w, r)
		} else {
			fmt.Fprintf(w, "%d: method not allowed\n", http.StatusMethodNotAllowed)
		}
	}

	return wrapper
}

func serveBody(handler func(ctx context.Context, body []byte) ([]byte, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(string(body))
		defer r.Body.Close()

		if result, err := handler(r.Context(), body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write(result)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func serveQuery(param string, handler func(ctx context.Context, q string) ([]byte, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Query().Get(param)
		if value == "" {
			http.Error(w, "empty param value", http.StatusBadRequest)
			return
		}

		if result, err := handler(r.Context(), value); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write(result)
			w.WriteHeader(http.StatusOK)
		}
	}
}
