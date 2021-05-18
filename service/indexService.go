package service

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/google/go-tika/tika"
	"github.com/yashmeh/doc-rank/elasticApi"
	"github.com/yashmeh/doc-rank/tikaApi"
)

type IndexService interface {
	IndexDoc(dir string) error
}

type indexService struct {
	EClient *elasticsearch.Client
	TClient *tika.Client
}

func NewIndexService(elasticC elasticApi.ElasticServer, tikaC tikaApi.TikaServer) IndexService {
	return &indexService{EClient: elasticC.Get(), TClient: tikaC.Get()}
}

//This is the method that loads all the documents to elastic search
func (s *indexService) IndexDoc(dir string) error {
	return nil
}
