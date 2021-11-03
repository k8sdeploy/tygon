package tygon

import (
  "encoding/json"
  "net/http"

  bugLog "github.com/bugfixes/go-bugfixes/logs"
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
  var decodedJSON interface{}

  if err := json.NewDecoder(r.Body).Decode(&decodedJSON); err != nil {
    jsonError(w, "Invalid JSON", err)
    return
  }

  // it's a ping event
  if err := mapstructure.Decode(decodedJSON, &t.PingEvent); err == nil {
    if err := t.PingEventTriggered(); err != nil {
      jsonError(w, "Ping event failed", err)
      return
    }
  }

  w.WriteHeader(http.StatusOK)
	bugLog.Infof("Request: %+v", decodedJSON)
  for name, values := range r.Header {
    for _, value := range values {
      bugLog.Infof("Header: %s: %s", name, value)
    }
  }
}
