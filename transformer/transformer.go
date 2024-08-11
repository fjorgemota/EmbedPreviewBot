package transformer

import (
	"errors"
	"net/url"
)

// TransformURL processes the given URL to remove the query string and change the domain based on the input domain.
func TransformURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil || !parsedURL.IsAbs() {
		return "", errors.New("invalid or malformed URL")
	}

	var zero struct{}
	maintainQueryString := make(map[string]struct{})

	// Change the domain based on the input URL
	switch parsedURL.Host {
	case "twitter.com", "www.twitter.com", "x.com", "www.x.com":
		parsedURL.Host = "vxtwitter.com"
	case "instagram.com", "www.instagram.com":
		parsedURL.Host = "ddinstagram.com"
	case "reddit.com", "www.reddit.com":
		parsedURL.Host = "rxddit.com"
	case "old.reddit.com":
		parsedURL.Host = "old.rxddit.com"
	case "youtube.com", "www.youtube.com":
		parsedURL.Host = "yfxtube.com"
		maintainQueryString["v"] = zero
		maintainQueryString["t"] = zero
	case "youtu.be":
		parsedURL.Host = "fxyoutu.be"
	default:
		// Return the original URL if it's not from Twitter/X or Instagram
		return inputURL, nil
	}
	// exclude every query string that's not on maintainQueryString
	query := parsedURL.Query()
	var modified bool
	for key := range query {
		if _, ok := maintainQueryString[key]; !ok {
			modified = true
			query.Del(key)
		}
	}
	if modified {
		parsedURL.RawQuery = query.Encode()
	}

	return parsedURL.String(), nil
}
