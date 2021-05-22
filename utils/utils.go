package utils

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/yashmeh/Falabella/indexer"
)

// A function for marshaling structs to JSON string
func JsonStruct(doc *indexer.Document) (string, error) {

	// Create struct instance of the Elasticsearch fields struct object
	docStruct := &indexer.Document{
		Body:        doc.Body,
		ContentType: doc.ContentType,
		MetaData:    doc.MetaData,
		FileName:    doc.FileName,
	}

	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(docStruct)
	if err != nil {
		log.Error("json.Marshal ERROR:", err)
		return "", err
	}
	return string(b), nil
}

//A function to return the list of files in a directory
func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
