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
  var unknownPayload interface{}
	if err := json.NewDecoder(r.Body).Decode(&unknownPayload); err != nil {
		jsonError(w, "Invalid payload", err)
		return
	}

  // if ok, err := t.validateSecret(r.Header.Get("X-Hub-Signature-256"), body); !ok {
  //   jsonError(w, "Invalid secret", err)
  //   return
  // }

	// test different types
	if ok, parsedPayload := isPingEvent(unknownPayload); ok {
		if err := t.handlePingEvent(parsedPayload); err != nil {
			jsonError(w, "handlePingEvent failed", err)
			return
		}
	}

  if ok, parsedPayload := isPackageEvent(unknownPayload); ok {
    if err := t.handlePackageEvent(parsedPayload); err != nil {
      jsonError(w, "handlePackageEvent failed", err)
      return
    }
  }

  for name, values := range r.Header {
    for _, value := range values {
      bugLog.Infof("Header: %s: %s", name, value)
    }
  }

	w.WriteHeader(http.StatusOK)
	bugLog.Infof("Beep")
}
