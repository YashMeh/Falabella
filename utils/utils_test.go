package utils

import (
	"testing"

	"github.com/yashmeh/Falabella/indexer"
)

func TestJStruct(t *testing.T) {
	sampleDocument := &indexer.Document{
		Body:        "TestBody",
		ContentType: "TestContent",
		MetaData:    "TestMeta",
		FileName:    "TestFilename",
	}
	got, err := JsonStruct(sampleDocument)

	if err != nil {
		t.Errorf("[ERROR]:Error converting to JSON string")
	}

	want := `{"Body":"TestBody","ContentType":"TestContent","MetaData":"TestMeta","FileName":"TestFilename"}`
	if got != want {
		t.Errorf("got %q , want %q", got, want)
	}

}

func TestDir(t *testing.T) {
	got, err := IOReadDir("../testdir")
	if err != nil {
		t.Errorf("[ERROR]:Error reading files from directory")
	}
	want := []string{"test_file.go"}
	if got[0] != want[0] {
		t.Errorf("got %q , want %q", got, want)
	}
}
