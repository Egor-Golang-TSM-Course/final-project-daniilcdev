package internal

import (
	"context"
	"encoding/json"
	"errors"
	client "hash-system/hashing/pkg"
	"net/http"
)

var cl client.HashServiceClient

func StartApiServer() error {
	var err error
	cl, err = client.New()

	if err != nil {
		return err
	}

	http.HandleFunc("/getHash", method(http.MethodGet, serveQuery("payload", getHash)))
	http.HandleFunc("/checkHash", method(http.MethodGet, serveQuery("hash", checkHash)))
	http.HandleFunc("/createHash", method(http.MethodPost, serveBody(createHash)))

	return http.ListenAndServe(":8080", nil)
}

func getHash(ctx context.Context, payload string) ([]byte, error) {
	if payload == "" {
		return nil, errors.New("payload cannot be empty")
	}
	type response struct {
		Hash string `json:"hash"`
	}

	hash, err := cl.GetHash(ctx, payload)
	if err != nil {
		return nil, err
	}

	return json.Marshal(response{Hash: hash})
}

func checkHash(ctx context.Context, hash string) ([]byte, error) {
	if hash == "" {
		return nil, errors.New("cannot check empty hash")
	}

	exists, err := cl.CheckHash(ctx, hash)

	if err != nil {
		return nil, err
	}

	type response struct {
		Exists bool `json:"exists"`
	}

	return json.Marshal(response{Exists: exists})
}

func createHash(ctx context.Context, body []byte) ([]byte, error) {
	if len(body) == 0 {
		return nil, errors.New("cannot create hash for empty payload")
	}

	type payload struct {
		Src string `json:"src"`
	}

	var p payload
	if err := json.Unmarshal(body, &p); err != nil {
		return nil, err
	}

	hash, err := cl.CreateHash(ctx, p.Src)

	if err != nil {
		return nil, err
	}

	type response struct {
		Hash string `json:"hash"`
	}

	return json.Marshal(response{Hash: hash})
}
