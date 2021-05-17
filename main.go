package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-tika/tika"
)

func main() {
	f, err := os.Open("./negotiation_genius.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	client := tika.NewClient(nil, "http://localhost:9998")
	body, err := client.Parse(context.Background(), f)
	fmt.Println(body)
}
