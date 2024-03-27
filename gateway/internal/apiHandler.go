package internal

import (
	"context"
	"encoding/json"
	"errors"
)

type apiHandler struct {
	cl shared.HashServiceClient
}

func New() (*apiHandler, error) {
	cl, err := shared.New(":9000")

	if err != nil {
		return nil, err
	}

	return &apiHandler{
		cl: cl,
	}, nil
}

func (api *apiHandler) getHash(ctx context.Context, payload string) ([]byte, error) {
	if payload == "" {
		return nil, errors.New("payload cannot be empty")
	}
	type response struct {
		Hash string `json:"hash"`
	}

	hash, err := api.cl.GetHash(ctx, payload)
	if err != nil {
		return nil, err
	}

	return json.Marshal(response{Hash: hash})
}

func (api *apiHandler) checkHash(ctx context.Context, hash string) ([]byte, error) {
	if hash == "" {
		return nil, errors.New("cannot check empty hash")
	}

	exists, err := api.cl.CheckHash(ctx, hash)

	if err != nil {
		return nil, err
	}

	type response struct {
		Exists bool `json:"exists"`
	}

	return json.Marshal(response{Exists: exists})
}

func (api *apiHandler) createHash(ctx context.Context, body []byte) ([]byte, error) {
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

	hash, err := api.cl.CreateHash(ctx, p.Src)

	if err != nil {
		return nil, err
	}

	type response struct {
		Hash string `json:"hash"`
	}

	return json.Marshal(response{Hash: hash})
}
