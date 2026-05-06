package app

import (
	"context"
	"github.com/mephistolie/chefbook-backend-common/log"
	"github.com/mephistolie/chefbook-backend-common/shutdown"
	"github.com/mephistolie/chefbook-backend-service-template/internal/config"
	"time"
)

func Run(cfg *config.Config) {
	log.InitWithService("template", *cfg.LogsPath, *cfg.Environment == config.EnvDev)
	cfg.Print()

	wait := shutdown.Graceful(context.Background(), 5*time.Second, map[string]shutdown.Operation{})
	<-wait
}
