package conf

import (
	"github.com/caarlos0/env"
	"github.com/rs/zerolog/log"
	"time"
)

type Config struct {
	Addr            string        `env:"ADDR" envDefault:":3000"`
	IsProduction    bool          `env:"PRODUCTION"`
	FetcherDuration time.Duration `env:"FETCHER_DURATION" envDefault:"1m"`
}

var Env Config

func init() {
	err := env.Parse(&Env)
	if err != nil {
		log.Error().Msgf("failed to parse Conf: %v", err)
	} else {
		log.Debug().Msgf("app config: %+v", Env)
	}
}
