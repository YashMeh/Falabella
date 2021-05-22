package tikaApi

import (
	"github.com/google/go-tika/tika"
	"github.com/yashmeh/doc-rank/config"
)

type TikaServer interface {
	Get() *tika.Client
}

type tikaServer struct {
	Client *tika.Client
}

func NewTikaClient(c *config.Config) TikaServer {
	config := c.Get()
	url := config.GetString("services.apacheTika")
	client := tika.NewClient(nil, url)
	return &tikaServer{Client: client}

}

func (c *tikaServer) Get() *tika.Client {
	return c.Client
}
