package service

import (
	"context"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/go-tika/tika"
	log "github.com/sirupsen/logrus"
	"github.com/yashmeh/Falabella/config"
	"github.com/yashmeh/Falabella/indexer"
	"github.com/yashmeh/Falabella/parser"
	"github.com/yashmeh/Falabella/utils"
)

var (
	wg   sync.WaitGroup
	Flag chan bool
)

type IndexService interface {
	IndexDoc(c *config.Config) error
}

type indexService struct {
	EClient *elasticsearch.Client
	TClient *tika.Client
}

func NewIndexService(elasticC indexer.ElasticServer, tikaC parser.TikaServer) IndexService {
	return &indexService{EClient: elasticC.Get(), TClient: tikaC.Get()}
}

func ReadData(s *indexService, fileDir string, fileName string, ch chan<- *indexer.Document) {
	//Open the file
	f1, err := os.Open(fileDir)
	if err != nil {
		log.Error("[ERROR] Opening file")
	}
	f2, err := os.Open(fileDir)
	if err != nil {
		log.Error("[ERROR] Opening file")
	}
	f3, err := os.Open(fileDir)
	if err != nil {
		log.Error("[ERROR] Opening file")
	}
	defer f1.Close()
	defer f2.Close()
	defer f3.Close()
	c := context.Background()
	docBody, err := s.TClient.Parse(c, f1)
	if err != nil {
		log.Error("[ERROR] Reading body")
	}
	docContent, err := s.TClient.Detect(c, f2)
	if err != nil {
		log.Error("[ERROR] Reading MIMETYPE")
	}
	docMeta, err := s.TClient.Meta(c, f3)
	if err != nil {
		log.Error("[ERROR] Reading Meta-data")
	}
	tikaDocument := &indexer.Document{
		Body:        docBody,
		ContentType: docContent,
		MetaData:    docMeta,
		FileName:    fileName,
	}
	defer func() {
		log.Infof("Tika Processed: %s \n", fileName)
	}()
	//Send the document on the channel
	ch <- tikaDocument
}

func IndexData(s *indexService, ch <-chan *indexer.Document, index int) {
	defer wg.Done()
	var fileName string
	var statusCode int
	for tikaDocument := range ch {
		fileName = tikaDocument.FileName
		docString, err := utils.JsonStruct(tikaDocument)
		if err != nil {
			log.Error("[ERROR] Converting to JSON string")
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
		statusCode = res.StatusCode

		if err != nil {
			log.Error("[ERROR] Sending the request to elasticsearch")
		}
		defer res.Body.Close()
		break

	}
	defer func() {
		if statusCode == 200 {
			log.Infof("Elastic Indexed: %s \n", fileName)
		} else {
			log.Errorf("[ERROR] Elastic Indexed: %s \n", fileName)
		}
	}()
}

//This is the method that loads all the documents to elastic search
func (s *indexService) IndexDoc(c *config.Config) error {
	config := c.Get()
	dir := config.GetString("appConfig.filePath")
	//Get all the files for the directory
	files, err := utils.IOReadDir(dir)
	if err != nil {
		panic(err)
	}
	wg.Add(len(files))
	//Create a buffered channel
	//TODO: Tune according to performence
	docChannel := make(chan *indexer.Document, 10)
	for i, fileName := range files {
		go ReadData(s, dir+fileName, fileName, docChannel)
		go IndexData(s, docChannel, i)
	}
	//TODO:Understand the deadlock condition better
	wg.Wait()
	close(docChannel)
	return nil
}
