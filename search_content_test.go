package main

import (
	"log"
	"strings"
	"testing"
	"time"
)

func TestSearchContent(t *testing.T) {
	dumpPath := "enwiki-latest-abstract1.xml.gz"
	query := "cat"

	log.Println("Starting")

	start := time.Now()
	docs, err := loadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchedDocs := searchContent(docs, query)
	log.Printf("Search found %d documents in %v", len(matchedDocs), time.Since(start))
}

func searchContent(docs []Document, term string) []Document {
	var r []Document
	for _, doc := range docs {
		if strings.Contains(doc.Text, term) {
			r = append(r, doc)
		}
	}
	return r
}
