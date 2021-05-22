package service

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/go-tika/tika"
	"github.com/yashmeh/doc-rank/elasticApi"
	"github.com/yashmeh/doc-rank/tikaApi"
	"github.com/yashmeh/doc-rank/utils"
)

var (
	wg   sync.WaitGroup
	Flag chan bool
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

func ReadData(s *indexService, fileDir string, fileName string, ch chan<- *elasticApi.Document) {
	//Open the file
	f, err := os.Open(fileDir)
	if err != nil {
		fmt.Println("[ERROR] Opening file")
	}
	defer f.Close()
	c := context.Background()
	docBody, err := s.TClient.Parse(c, f)
	if err != nil {
		fmt.Println("[ERROR] Reading body")
	}
	// docContent, err := s.TClient.Detect(c, f)
	// if err != nil {
	// 	fmt.Println("[ERROR] Reading MIMETYPE")
	// }
	// docMeta, err := s.TClient.Meta(c, f)
	// if err != nil {
	// 	fmt.Println("[ERROR] Reading Meta-data")
	// }
	tikaDocument := &elasticApi.Document{
		Body:        docBody,
		ContentType: "", //docContent,
		// MetaData:    "", //docMeta,
		FileName: fileName,
	}
	defer func() {
		fmt.Printf("Tika Processed: %s \n", fileName)
	}()
	//Send the document on the channel
	ch <- tikaDocument
}

func IndexData(s *indexService, ch <-chan *elasticApi.Document, index int) {
	defer wg.Done()
	var fileName string
	for tikaDocument := range ch {
		fileName = tikaDocument.FileName
		docString, err := utils.JsonStruct(tikaDocument)
		if err != nil {
			fmt.Println("[ERROR] Converting to JSON string")
		}
		// Instantiate a request object
		req := esapi.IndexRequest{
			Index:      "doc-search",
			DocumentID: strconv.Itoa(index + 1),
			Body:       strings.NewReader(docString),
			Refresh:    "true",
		}

		// Return an API response object from request
		res, err := req.Do(context.Background(), s.EClient)
		if err != nil {
			fmt.Println("[ERROR] Sending the request to elasticsearch")
		}
		defer res.Body.Close()
		break

	}
	defer func() {
		fmt.Printf("Elastic Indexed: %s \n", fileName)
	}()
}

//This is the method that loads all the documents to elastic search
func (s *indexService) IndexDoc(dir string) error {
	//Get all the files for the directory
	files, err := utils.IOReadDir(dir)
	if err != nil {
		panic(err)
	}
	wg.Add(len(files))
	//Create a buffered channel
	//TODO: Tune according to performence
	docChannel := make(chan *elasticApi.Document, 10)
	for i, fileName := range files {
		go ReadData(s, dir+fileName, fileName, docChannel)
		go IndexData(s, docChannel, i)
	}
	//TODO:Understand the deadlock condition better
	wg.Wait()
	close(docChannel)
	return nil
}
