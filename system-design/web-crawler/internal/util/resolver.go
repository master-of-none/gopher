package util

import (
	"net/url"
	"strings"
)

func Normalize(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", err
	}

	u.Fragment = ""
	u.RawQuery = ""
	u.Path = strings.TrimSuffix(u.Path, "/")

	return u.String(), nil
}

func ResolveURL(baseStr string, href string) (string, error) {
	base, err := url.Parse(baseStr)

	if err != nil {
		return "", err
	}

	ref, err := url.Parse(href)
	if err != nil {
		return "", err
	}

	return base.ResolveReference(ref).String(), nil
}
