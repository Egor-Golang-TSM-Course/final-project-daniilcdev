package internal

import (
	"fmt"
	"net/http"
)

func method(m string, source func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == m {
			source(w, r)
		} else {
			fmt.Fprintf(w, "%d: method not allowed\n", http.StatusMethodNotAllowed)
		}
	}

	return wrapper
}
