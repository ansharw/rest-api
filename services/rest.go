package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type RestHandler struct {
	srvSession *Cockatoo
}

type restResponse struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (rest *RestHandler) RegisterCockatoo(co *Cockatoo) {
	rest.srvSession = co
}

func (rest *RestHandler) RestCreateSession(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	request := &sessionReq{}
	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&restResponse{
			Code:    http.StatusInternalServerError,
			Data:    nil,
			Message: fmt.Sprintf("ther is problem when handling the request: %s", err.Error()),
		})
		return
	}

	res, err := rest.srvSession.ProcessSetSession(ctx, request)
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(&restResponse{
				Code:    http.StatusServiceUnavailable,
				Data:    nil,
				Message: fmt.Sprintf("Error when connect to redis, error: %v", err.Error()),
			})
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&restResponse{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: res.Message,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&restResponse{
		Code:    http.StatusOK,
		Data:    nil,
		Message: res.Message,
	})
}
