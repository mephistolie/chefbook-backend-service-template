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

	ctx := context.Background()

	wait := shutdown.Graceful(ctx, 5*time.Second, map[string]shutdown.Operation{})
	<-wait
}
