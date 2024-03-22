package internal

import (
	"fmt"
	"net/http"
)

func StartApiServer() error {
	http.HandleFunc("/getHash/", method(http.MethodGet, echoEnpoint))
	http.HandleFunc("/checkHash/", method(http.MethodGet, echoEnpoint))
	http.HandleFunc("/createHash/", method(http.MethodPost, echoEnpoint))

	return http.ListenAndServe(":8080", nil)
}

func echoEnpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.RequestURI)
}
