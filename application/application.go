package application

import (
	"bitbucket.org/bridce/go_exercise2/config"
	"bitbucket.org/bridce/go_exercise2/router"
	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"
)

func StartApp() {
	addr := config.Config.ServiceHost + ":" + config.Config.ServicePort
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.NewRouter(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}
