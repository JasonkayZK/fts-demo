package main

import (
	"log"
	"regexp"
	"testing"
	"time"
)

func TestRegexpSearch(t *testing.T) {
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
	matchedDocs := regexpSearch(docs, query)
	log.Printf("Search found %d documents in %v", len(matchedDocs), time.Since(start))
}

func regexpSearch(docs []Document, term string) []Document {
	// Don't do this in production, it's a security risk. term needs to be sanitized.
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)
	var r []Document
	for _, doc := range docs {
		if re.MatchString(doc.Text) {
			r = append(r, doc)
		}
	}
	return r
}
