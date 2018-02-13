// This crawler takes a collection of urls and determines if a particular string
// is in the response body of the given collection
package main

import (
	"github.com/ackimwilliams/fetcher"
	"fmt"
	"time"
)

// Crawler configuration
const (
	defaultWorkerPoolSize = 20
)

type Crawler struct {
	urls []string

	//thread management
	currentUrlIndex int	//also tells no of urls used
	noConcurrentThreads int

	// crawler dependencies
	client fetcher.IFetcher
	parser Parser
}

// NewCrawler uses composite literal to create a new crawler
func NewCrawler(urls []string, client fetcher.IFetcher, parser Parser) Crawler {
	return Crawler{urls, 0, defaultWorkerPoolSize, client, parser}
}

//determine if url contains search term
//@todo: some requests take too extremely to complete, future iterations should timeout the request
func (c *Crawler) worker(jobId int, jobs <-chan string, results chan<- string) {
	for url := range jobs {
		start := time.Now()

		//prepare http request
		reqHttpComplete := make(chan bool, 1)
		content, err := c.client.GetFileContentAsString(url, reqHttpComplete)
		<-reqHttpComplete

		if err != nil {
			//@todo: better error handling, for now ignore error handling and ignore url
			results <- ""
			continue
		}

		if !c.parser.Contains(content) {
			results <- ""
		} else {
			results <- url	//add url to list of results found
		}

		fmt.Println(fmt.Sprintf("Worker ID = %d; site = %s; time to complete = %s", jobId, url,time.Since(start)))
	}
}

//search urls for search rules in parser
func (c *Crawler) SearchUrls() []string {
	noUrls := len(c.urls)
	foundUrls := make([]string, 0)

	jobs := make(chan string, noUrls)	//jobs queue
	results := make(chan string, noUrls)

	for w := 0; w < defaultWorkerPoolSize; w++ {
		go c.worker(w, jobs, results)
	}

	for j := 0; j < noUrls; j++ {
		jobs <- c.urls[j]	//add url to jobs queue
	}

	close(jobs)	// close jobs queue

	matched:= ""

	for r := 0; r < noUrls; r++ {
		matched = <-results

		if matched != "" {
			foundUrls = append(foundUrls, matched)
		}
	}

	return foundUrls
}
