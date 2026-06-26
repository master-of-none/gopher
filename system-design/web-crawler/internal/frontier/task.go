package frontier

type CrawlTask struct {
	URL   string `json:"url"`
	Depth int    `json:"depth"`
}
