package servicediscovery

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type (
	UriEntry struct {
		Protocol string
		Host     string
		Port     string
		Endpoint string
	}

	UriParser struct {
		urlEntry *UriEntry
	}
)

func NewUriParser(uri string) *UriParser {
	parser := &UriParser{}
	parser.parse(uri)
	return parser
}

func (parser *UriParser) parse(uri string) *UriEntry {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	parser.urlEntry = &UriEntry{}
	parser.urlEntry.Protocol = u.Scheme
	u.Host = strings.Replace(u.Host, "[", "", -1)
	u.Host = strings.Replace(u.Host, "]", "", -1)
	parser.urlEntry.Host = u.Host
	parser.urlEntry.Endpoint = u.Path
	if u.RawQuery != "" {
		parser.urlEntry.Endpoint = parser.urlEntry.Endpoint + "?" + u.RawQuery
	}
	return parser.urlEntry
}

func (parser *UriParser) GetUriEntry() *UriEntry {
	return parser.urlEntry
}

func (parser *UriParser) Generate(instanceHosting string) string {
	if parser.urlEntry != nil {
		uri := fmt.Sprintf("%s://%s%s", parser.urlEntry.Protocol, instanceHosting, parser.urlEntry.Endpoint)
		return uri
	} else {
		panic(errors.New("not parse"))
	}
}
