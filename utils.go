package main

import "net/url"

// GetHostname returns the hostname of the rawurl
func GetHostname(rawurl string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	} else {
		return u.Hostname(), nil
	}
}

// IsSameHostName checks if the hostname of the rawurl is th4 same as targetHost
func IsSameHostName(rawurl, targetHost string) bool {
	host, err := GetHostname(rawurl)
	if err != nil {
		return false
	} else {
		return host == targetHost
	}
}

// ResolveUrl resolves the rawurl with pageurl
func ResolveUrl(rawurl, pageurl string) string {
	u, err := url.Parse(rawurl)
	if err != nil {
		return rawurl
	}
	p, err := url.Parse(pageurl)
	if err != nil {
		return pageurl
	}
	return p.ResolveReference(u).String()
}
