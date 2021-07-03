package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/yashmeh/Falabella/utils"
)

func main() {
	// c := config.NewConfig()
	// aE, err := indexer.NewElasticClient(c)
	// if err != nil {
	// 	log.Error("[ERROR] Connecting to elasticsearch")
	// }
	// aT := parser.NewTikaClient(c)

	// sI := service.NewIndexService(aE, aT)
	// sI.IndexDoc(c)

	got, err := utils.IOReadDir("./testdir")
	if err == nil {
		log.Info(got)
	}

}
