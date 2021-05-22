package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yashmeh/doc-rank/config"
	"github.com/yashmeh/doc-rank/elasticApi"
	"github.com/yashmeh/doc-rank/service"
	"github.com/yashmeh/doc-rank/tikaApi"
)

func main() {
	c := config.NewConfig()
	aE, err := elasticApi.NewElasticClient(c)
	if err != nil {
		log.Error("[ERROR] Connecting to elasticsearch")
	}
	aT := tikaApi.NewTikaClient(c)

	sI := service.NewIndexService(aE, aT)
	sI.IndexDoc(c)

	// select {
	// case j := <-service.Flag:
	// 	fmt.Printf("All Done %t ", j)

	// }

}
