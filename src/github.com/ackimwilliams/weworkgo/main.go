// All packages, classes, code was written from scratch for this task. No libraries were used
package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"github.com/ackimwilliams/fetcher"
	"regexp"
	"time"
)

const (
	outputFile            = "result.txt"
	rudimentarySearchTerm = `(?i)timeline`
	urlListPath           = `https://s3.amazonaws.com/fieldlens-public/urls.txt`
)

// Parse urls from url file
func parseUrls(blob string) []string {
	//@todo: rewrite sanitizing url functionality
	reg, _ := regexp.Compile(`[^a-zA-Z0-9\.-_\\\/&]+`)	// simple sanitizer

	if blob == "" {
		return nil
	}

	var urls []string

	s := bufio.NewScanner(strings.NewReader(blob))
	s.Scan() //ignore header

	for s.Scan() {
		url := strings.Split(s.Text(), ",")[1] //explode by comma
		url = reg.ReplaceAllString(url, "")	//sanitize url

		if url != "" {
			urls = append(urls, url)
		}
	}

	return urls
}

// writeOutput writes urls with search in result file
func writeOutput(urls []string) {
	file, e := os.Create(outputFile)

	if e != nil {
		panic(e)
	}

	w := bufio.NewWriter(file)
	w.WriteString("Search term was found in the following urls:\n")
	w.WriteString(strings.Join(urls, "\n"))
	w.Flush()
}

func main() {
	start := time.Now()
	fmt.Println("Beginning program!")
	fmt.Print("\tFetching list of urls...")

	//create channels
	chReqUrlComplete := make(chan bool, 1)

	f, _ := fetcher.Create(fetcher.TypeHttp, urlListPath, &fetcher.HttpClient{})	//create http
	content, err := f.GetFileContentAsString(urlListPath, chReqUrlComplete)

	<-chReqUrlComplete //receive msg

	fmt.Print(" done.\n\n")

	if err != nil {
		panic(fmt.Sprintf("cannot retrieve remote file '%s'", urlListPath))
	}

	urls := parseUrls(content)

	fmt.Println("\tCrawling urls...")
	// crawl urls
	c := NewCrawler(urls, f, NewParser(rudimentarySearchTerm))
	matchedUrls := c.SearchUrls()
	fmt.Print(" done.\n\n")

	fmt.Print("\tWriting matched results to file...")
	writeOutput(matchedUrls)	//write output to file
	fmt.Print(" done.\n\n")

	fmt.Println("Program exit!")
	fmt.Println(fmt.Sprintf("\n\nTotal execution time: %s", time.Since(start)))
}
