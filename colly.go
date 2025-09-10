package impcolly

import (
	"github.com/gocolly/colly"
	"github.com/imroc/req/v3"
)

var (
	impClient = req.C()
)

func init() {
	impClient.ImpersonateChrome()
}

func NewCollector(o ...func(*colly.Collector)) *colly.Collector {
	c := colly.NewCollector(o...)
	c.WithTransport(impClient.Transport)
	return c
}
