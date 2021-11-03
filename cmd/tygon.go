package main

import (
  "time"

  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"

  bugLog "github.com/bugfixes/go-bugfixes/logs"
  bugfixes "github.com/bugfixes/go-bugfixes/middleware"
  "github.com/k8sdeploy/tygon/internal/config"
  "github.com/keloran/go-probe"
)

func main() {
  bugLog.Local().Info("Starting Tygon")

  cfg, err := config.BuildConfig()
  if err != nil {
    _ = bugLog.Errorf("buildConfig: %+v", err)
    return
  }

  if err := route(cfg); err != nil {
    _ = bugLog.Errorf("route failed: %+v", err)
    return
  }
}

func route(cfg config.Config) error {
  r := chi.NewRouter()
  r.Use(middleware.Timeout(60 * time.Second))
  r.Use(middleware.Heartbeat("/ping"))
  r.Use(bugfixes.BugFixes)

  r.Route("/", func(r chi.Router) {
    r.Post("/", )
  })

  r.Route("/probe", func(r chi.Router) {
    r.Get("/", probe.HTTP)
  })

  return nil
}
