package parser

import (
	"context"
	"os"
	"testing"

	"github.com/yashmeh/Falabella/config"
)

func TestTika(t *testing.T) {
	config := config.NewConfig("../")
	tikaServer := NewTikaClient(config)
	tikaClient := tikaServer.Get()
	c := context.Background()
	f, err := os.Open("../testdir/test.txt")
	if err != nil {
		t.Errorf("[ERROR]:Cannot open file")
	}
	got, err := tikaClient.Detect(c, f)
	want := "text/plain"
	if err != nil {
		t.Errorf("[ERROR]:Detecting using Tika")
	}
	if want != got {
		t.Errorf("%q want, %q got", want, got)
	}
}
