package indexer

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/yashmeh/Falabella/config"
)

type ElasticServer interface {
	Get() *elasticsearch.Client
}

type elasticServer struct {
	Client *elasticsearch.Client
}

func NewElasticClient(c *config.Config) (ElasticServer, error) {
	config := c.Get()
	url := config.GetString("services.elasticSearch")
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
	_, err2 := ElasticClient.Info()

	// Deserialize the response into a map.
	if err2 != nil {
		return nil, err2
	}
	return &elasticServer{Client: ElasticClient}, nil
}

func (c *elasticServer) Get() *elasticsearch.Client {
	return c.Client
}
