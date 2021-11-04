package tygon

import (
	"encoding/json"
	"net/http"

	bugLog "github.com/bugfixes/go-bugfixes/logs"
  "github.com/google/go-github/v39/github"
  "github.com/mitchellh/mapstructure"
)

func jsonError(w http.ResponseWriter, msg string, errs error) {
	bugLog.Local().Infof("jsonError: %+v", errs)

	w.Header().Set("Content-Type", "text/json")
	if err := json.NewEncoder(w).Encode(struct {
		Error string
	}{
		Error: msg,
	}); err != nil {
		bugLog.Debugf("send %s failed: %+v", msg, err)
	}
}

func (t Tygon) ParsePayload(w http.ResponseWriter, r *http.Request) {
	var decodedPing github.Hook

	// it's a ping event
	if err := json.NewDecoder(r.Body).Decode(&decodedPing); err == nil {
		if err := t.PingEventTriggered(decodedPing); err != nil {
			jsonError(w, "Ping event failed", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
  var unknownPayload interface{}
  if err := json.NewDecoder(r.Body).Decode(&unknownPayload); err != nil {
    jsonError(w, "unknown payload", err)
    return
  }
	bugLog.Infof("Request: %+v", unknownPayload)
	for name, values := range r.Header {
		for _, value := range values {
			bugLog.Infof("Header: %s: %s", name, value)
		}
	}
}
