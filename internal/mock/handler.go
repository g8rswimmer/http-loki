package mock

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/g8rswimmer/http-loki/internal/model"
)

type pair struct {
	request  model.Request
	response model.Response
}

type Handler struct {
	pairs []pair
}

func (h *Handler) Add(req model.Request, resp model.Response) {
	h.pairs = append(h.pairs, pair{request: req, response: resp})
}

func (h *Handler) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	p, err := h.requestPair()
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(p.response.StatusCode)
	payload, _ := json.Marshal(p.response.Body)
	w.Write(payload)
}

func (h *Handler) requestPair() (pair, error) {
	if len(h.pairs) == 0 {
		return pair{}, fmt.Errorf("no request pair")
	}
	for _, p := range h.pairs {
		if p.request.Body == nil {
			return p, nil
		}
	}
	return pair{}, fmt.Errorf("no request pair")
}
