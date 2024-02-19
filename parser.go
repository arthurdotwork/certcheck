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

	return url.Parse(link)
}
