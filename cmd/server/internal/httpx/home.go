package httpx

import (
	"encoding/json"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	payload := struct {
		Msg string `json:"msg"`
	}{
		Msg: "welcome to the loki server",
	}
	w.Header().Add("content-type", "application/json")
	enc, _ := json.Marshal(payload)
	w.Write(enc)
}
