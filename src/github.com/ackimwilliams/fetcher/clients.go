// List of transport packages used by fetcher objects
package fetcher

import (
	"net/http"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// Client interface for ftpFetch, web, etc
type ITransportClient interface {
	Get(url string) (string, error)

	//@todo: more methods here
}

//custom default protocol
const (
	defaultProtocol = "http"
	timeout = 60	//secs
)

//HttpClient http transport client
type HttpClient struct {}

// Get method makes httpFetch GET calls
func (h *HttpClient) Get(url string) (string, error) {
	client := http.Client{
		Timeout: timeout * time.Second,
	}

	//check if protocol is included
	//@todo: remove this functionality here and move it into the parser
	if !strings.Contains(url, defaultProtocol) {
		url = defaultProtocol + "://" + url
	}

	r, err := client.Get(url)

	if err != nil {
		return "", errors.New(fmt.Sprintf("err fetching webpage %s, %s", url, err))
	}

	content, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return "", errors.New(fmt.Sprintf("err reading r from page %s, %s", url, err))
	}

	return string(content), nil
}

//@todo: add other methods such as ftp or filesystem