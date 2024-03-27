package internal

import (
	"net/http"
)

func StartApiServer() error {
	apiHandler, err := New()
	if err != nil {
		return err
	}

	http.HandleFunc("/getHash", method(http.MethodGet, serveQuery("payload", apiHandler.getHash)))
	http.HandleFunc("/checkHash", method(http.MethodGet, serveQuery("hash", apiHandler.checkHash)))
	http.HandleFunc("/createHash", method(http.MethodPost, serveBody(apiHandler.createHash)))

	return http.ListenAndServe(":8080", nil)
}
