package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yashmeh/Falabella/config"
	"github.com/yashmeh/Falabella/indexer"
	"github.com/yashmeh/Falabella/parser"
	"github.com/yashmeh/Falabella/service"
)

func main() {
	c := config.NewConfig(".")
	aE, err := indexer.NewElasticClient(c)
	if err != nil {
		log.Error("[ERROR] Connecting to elasticsearch")
	}
	aT := parser.NewTikaClient(c)

	sI := service.NewIndexService(aE, aT)
	sI.IndexDoc(c)

}
