package internal

import (
	"context"
	"encoding/json"
	"errors"
	"gateway/transport"
	"log"
	"net/http"
	"shared/cfg"
	grpcProxy "shared/grpc/client"
)

type apiHandler struct {
	cl grpcProxy.HashServiceClient
}

func StartApiServer(vars *cfg.EnvConfig) error {
	apiHandler, err := New(vars)
	if err != nil {
		return err
	}

	http.HandleFunc("/getHash", method(http.MethodGet, serveQuery("payload", apiHandler.getHash)))
	http.HandleFunc("/checkHash", method(http.MethodGet, serveQuery("hash", apiHandler.checkHash)))
	http.HandleFunc("/createHash", method(http.MethodPost, serveBody(apiHandler.createHash)))

	log.Printf("Starting gateway at %s", vars.GatewayPort)
	return http.ListenAndServe(vars.GatewayPort, nil)
}

func New(vars *cfg.EnvConfig) (*apiHandler, error) {
	cl, err := grpcProxy.New(vars.HashingServiceAddr)

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

	hash, err := api.cl.GetHash(ctx, payload)
	if err != nil {
		return nil, err
	}

	return json.Marshal(transport.ResponseWithHash{Hash: hash})
}

func (api *apiHandler) checkHash(ctx context.Context, hash string) ([]byte, error) {
	if hash == "" {
		return nil, errors.New("cannot check empty hash")
	}

	exists, err := api.cl.CheckHash(ctx, hash)

	if err != nil {
		return nil, err
	}

	return json.Marshal(transport.ResponseExists{Exists: exists})
}

func (api *apiHandler) createHash(ctx context.Context, body []byte) ([]byte, error) {
	if len(body) == 0 {
		return nil, errors.New("cannot create hash for empty payload")
	}

	var p transport.RequestWithPayload
	if err := json.Unmarshal(body, &p); err != nil {
		return nil, err
	}

	hash, err := api.cl.CreateHash(ctx, p.Payload)

	if err != nil {
		return nil, err
	}

	return json.Marshal(transport.ResponseWithHash{Hash: hash})
}
