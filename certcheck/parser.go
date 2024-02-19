package certcheck

import (
	"fmt"
	"net/url"
	"strings"
)

func Parse(link string) (*url.URL, error) {
	if link == "" {
		return nil, fmt.Errorf("empty URL")
	}

	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "https://" + link
	}

	parsedURL, err := url.Parse(link)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	return parsedURL, nil
}
