package tygon

import (
  "net/http"

  bugLog "github.com/bugfixes/go-bugfixes/logs"
)

func ParsePayload(w http.ResponseWriter, r *http.Request) {
  bugLog.Infof("Request: %+v", r.Body)
}
