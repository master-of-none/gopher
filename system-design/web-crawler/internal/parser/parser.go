package parser

import (
	"strings"
	"web-crawler/internal/util"

	"github.com/PuerkitoBio/goquery"
)

func ExtractLinks(baseURL string, html string) []string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		return nil
	}

	var links []string

	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")

		if !ok {
			return
		}

		link, err := util.ResolveURL(baseURL, href)

		if err == nil {
			links = append(links, link)
		}
	})

	return links
}
