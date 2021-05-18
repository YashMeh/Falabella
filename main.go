package main

import (
	"fmt"

	"github.com/yashmeh/doc-rank/elasticApi"
	"github.com/yashmeh/doc-rank/service"
	"github.com/yashmeh/doc-rank/tikaApi"
)

func main() {
	aE, err := elasticApi.NewElasticClient("http://localhost:9200")
	if err != nil {
		fmt.Println("[ERROR] Connecting to elasticsearch")
	}
	aT := tikaApi.NewTikaClient("http://localhost:9998")

	sI := service.NewIndexService(aE, aT)
	sI.IndexDoc("./assets")

}
