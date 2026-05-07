package filter

import "net/url"

func SameDomain(seed string, target string) bool {
	s, err1 := url.Parse(seed)
	t, err2 := url.Parse(target)

	if err1 != nil || err2 != nil {
		return false
	}

	return s.Host == t.Host
}
