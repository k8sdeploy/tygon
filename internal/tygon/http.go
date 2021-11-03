package tygon

import (
  "encoding/json"
  "net/http"

	bugLog "github.com/bugfixes/go-bugfixes/logs"
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
  var a interface{}

  if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
    jsonError(w, "Invalid JSON", err)
    return
  }

  w.WriteHeader(http.StatusOK)
	bugLog.Infof("Request: %+v", a)

  for name, values := range r.Header {
    for _, value := range values {
      bugLog.Infof("Header: %s: %s", name, value)
    }
  }
}
