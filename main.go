package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yashmeh/doc-rank/config"
	"github.com/yashmeh/doc-rank/indexer"
	"github.com/yashmeh/doc-rank/parser"
	"github.com/yashmeh/doc-rank/service"
)

func main() {
	c := config.NewConfig()
	aE, err := indexer.NewElasticClient(c)
	if err != nil {
		log.Error("[ERROR] Connecting to elasticsearch")
	}
	aT := parser.NewTikaClient(c)

	sI := service.NewIndexService(aE, aT)
	sI.IndexDoc(c)

}
