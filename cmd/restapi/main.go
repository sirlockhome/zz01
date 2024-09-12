package main

import (
	"fmt"
	"foxomni/internal/common"
	"foxomni/internal/core/server"
	"foxomni/pkg/config"
	"foxomni/pkg/database"

	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("what?")
	conf, err := config.LoadConfig("local_config")
	if err != nil {
		log.Fatal().Err(err).Msg("can't not load config")
	}

	sql, err := database.NewSQL(conf.SQLServer)
	if err != nil {
		log.Fatal().Err(err).Msg("error init SQLServer")
	}

	common.Domain = conf.Server.Addr

	httpserver := server.NewHTTPServer(*conf, sql)

	if err := httpserver.RunHTTPServer(); err != nil {
		log.Fatal().Err(err).Msg("server error")
	}
}
