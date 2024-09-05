package databases

import (
	"fmt"
	"net/url"
)

type Database interface {
	Dump(bool, string) (string, error)
	Restore(string) (bool, error)
	RunQuery(string) (string, error)
}

func ParseUrl(urlString string) (Database, error) {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	if parsedUrl.Scheme == "postgresql" {
		return parsePostgres(parsedUrl)
	}

	return nil, fmt.Errorf("unsupported database type: %s", parsedUrl.Scheme)
}
