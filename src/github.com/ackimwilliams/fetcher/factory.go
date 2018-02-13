// This package contains utilities that assists with getting content remotely.
// This factory creates fetcher objects
package fetcher

import (
	"errors"
	"io/ioutil"
	"fmt"
)

//Fetch types supported, update as necessary
const (
	TypeFtp             = 0
	TypeHttp            = 1
	TypeLocalFileSystem = 2
	//@todo: add more here
)

//------------------------------------------- content fetchers
// regex used for path validation by fetchers
const (
	regexHttpStringValidator = `^((http[s]?):\/)?\/?([^:\/\s]+)((\/\w+)*\/)([\w\-\.]+[^#?\s]+)(.*)?(#[\w\-]+)?$`
	regexFtpStringValidator = `^((ftp):\/)?\/?([^:\/\s]+)((\/\w+)*\/)([\w\-\.]+[^#?\s]+)(.*)$`
	regexFileSystemStringValidator = `^(.+)/([^/]+)$`
)

//// -------------------------------------------- local file system behaviors
// localFileSystemFetch gets content from files on the file system
type localFileSystemFetch struct {
	*abstractFetcher
}

func (l *localFileSystemFetch) GetFileContentAsString(path string, reqComplete chan bool) (string, error) {
	if l.contentPath == "" {
		return "", errors.New("cannot fetch file content without setting a content path")
	}

	content, err := ioutil.ReadFile(l.contentPath)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

//-------------------------------------------- ftpFetch behaviors
type ftpFetch struct {
	*abstractFetcher
}

// GetFileContentAsString read entire file as string
func (f *ftpFetch) GetFileContentAsString(path string, reqComplete chan bool) (string, error) {
	//@todo: //continue implementation
	return "", errors.New("not yet implemented")
}

// -------------------------------------------- httpFetch behaviors
type httpFetch struct {
	*abstractFetcher
}

// GetFileContentAsString reads remote files returning their string content
func (h *httpFetch) GetFileContentAsString(path string, reqComplete chan bool) (string, error) {
	content, err := h.client.Get(path)

	defer func() {
		reqComplete <- true
	}()
	if err != nil {
		//ignore error
		return "", nil
	}

	return content, nil
}

// Create makes fetch objects of types specified
func Create(t int, path string, c ITransportClient) (IFetcher, error) {
	var f IFetcher

	switch t {
		case TypeFtp:
			a := &abstractFetcher{ pathRegex: regexFtpStringValidator, client: c}
			f =  &ftpFetch{a}
			a.IFetcher = f
		case TypeHttp:
			a := &abstractFetcher{ pathRegex: regexHttpStringValidator, client: c}
			f =  &httpFetch{a}
			a.IFetcher = f
		case TypeLocalFileSystem:
			a := &abstractFetcher{ pathRegex: regexFileSystemStringValidator, client: c}
			f =  &localFileSystemFetch{a}
			a.IFetcher = f
		default:
			return nil, errors.New(fmt.Sprintf("cannot create type %s", t))
	}

	//@todo: handle path errors, ignored for now
	f.SetContentPath(path)

	return f, nil
}
