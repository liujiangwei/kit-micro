package main

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func InfoEndpoint(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(InfoRequest)

	return InfoResponse{1, "", req.Str}, nil
}

type InfoRequest struct {
	Str string
}

type InfoResponse struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    string `json:"data"`
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request InfoRequest
	r.ParseForm()

	if len(r.Form["string"]) > 0 {
		request.Str = r.Form["string"][0]
		return request, nil
	} else {
		return request, errors.New("string not exists")
	}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	http.Handle("/info", httptransport.NewServer(InfoEndpoint, decodeRequest, encodeResponse))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
