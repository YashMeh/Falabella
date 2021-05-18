package tikaApi

import "github.com/google/go-tika/tika"

type TikaServer interface {
	Get() *tika.Client
}

type tikaServer struct {
	Client *tika.Client
}

func NewTikaClient(url string) TikaServer {
	client := tika.NewClient(nil, url)
	return &tikaServer{Client: client}

}

func (c *tikaServer) Get() *tika.Client {
	return c.Client
}
