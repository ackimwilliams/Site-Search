// Fetcher package used to grab content from remote stores
package fetcher

import (
	"errors"
	"regexp"
	"fmt"
)

// IFetcher defines categories of objects used for fetching content
type IFetcher interface {
	// set path to content
	SetContentPath(contentPath string) error

	// get content of file as string
	GetFileContentAsString(path string, reqComplete chan bool) (string, error)
}

// abstractFetcher fetcher used to create new fetchers
type abstractFetcher struct {
	IFetcher
	contentPath string
	pathRegex string	//@todo: use for validation

	//inject dependencies
	client ITransportClient
}

// SetContentPath validates and sets content path
func (a *abstractFetcher) SetContentPath(contentPath string) error {
	match, _ := regexp.Match(a.pathRegex, []byte(contentPath))

	if !match {
		return errors.New(fmt.Sprintf("cannot use path '%s' with this fetcher", contentPath))
	}

	a.contentPath = contentPath
	return nil
}

// GetFileContentAsString abstract method to get content of remote file as string
func (a *abstractFetcher) GetFileContentAsString(path string, reqComplete chan bool) (string, error) {
	return "",  errors.New("method not implemented")
}
