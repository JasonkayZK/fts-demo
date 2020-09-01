package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var dumpPath string

	flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.Parse()

	log.Println("Full-text engine starting...")

	start := time.Now()
	docs, err := loadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(Index)
	idx.add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("input query str:")
	for input.Scan() {
		query := input.Text()
		start = time.Now()
		matchedIDs := idx.search(query)
		log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

		for _, id := range matchedIDs {
			doc := docs[id]
			log.Printf("%d\t%s\n", id, doc.Text)
		}
	}
}
