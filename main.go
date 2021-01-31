package main

import (
	"hippo_query_analysis/config"
	"hippo_query_analysis/es"
	"hippo_query_analysis/server"
)

func init() {
	config.InitConfig("")
	cf := config.GetConfig()
	es.InitEsClient(cf)
}

func main() {
	cf := config.GetConfig()
	server.ThriftInit(cf)
}
