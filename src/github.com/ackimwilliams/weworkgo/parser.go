// Simple string parser
package main

import "regexp"

const (
	isParserCaseInsensitive = true
)

// Parser interface
type IParser interface {
	Contains(string, string) bool
}

type Parser struct {
	searchTerm string	//search regex
}

// Create new parser
func NewParser(searchTerm string) Parser {
	p := Parser{}
	p.SetSearchTerm(searchTerm)

	return p
}

// SetSearchTerm set the search term used in this parser
func (p *Parser) SetSearchTerm(searchTerm string) {
	if isParserCaseInsensitive {
		searchTerm = "(?i)" + searchTerm
	}

	//make search term case insensitive
	p.searchTerm = searchTerm
}

// Contains determine if a search term is in a string blob
func (p *Parser) Contains(blob string) bool {
	match, _ := regexp.Match(p.searchTerm, []byte(blob))

	return match
}
