package elasticApi

import (
	"github.com/elastic/go-elasticsearch/v7"
)

type ElasticServer interface {
	Get() *elasticsearch.Client
}

type elasticServer struct {
	Client *elasticsearch.Client
}

func NewElasticClient(url string) (ElasticServer, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			url,
		},
		// Username: "user",
		// Password: "pass",
	}
	ElasticClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &elasticServer{Client: ElasticClient}, nil
}

func (c *elasticServer) Get() *elasticsearch.Client {
	return c.Client
}
