package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/google/go-tika/tika"
)

func main() {
	f, err := os.Open("./pic.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	client := tika.NewClient(nil, "http://localhost:9998")
	// body, err := client.Parse(context.Background(), f)
	body2, err := client.Detect(context.Background(), f)
	body3, err := client.MetaRecursive(context.Background(), f)
	// fmt.Println(body)
	fmt.Println(body2)
	fmt.Println(body3)
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// Username: "user",
		// Password: "pass",
	}

	ElasticClient, err := elasticsearch.NewClient(cfg)

	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
	}

	// Have the client instance return a response
	res, err := ElasticClient.Info()

	// Deserialize the response into a map.
	if err != nil {
		log.Fatalf("client.Info() ERROR:", err)
	} else {
		log.Printf("client response:", res)
	}

}
