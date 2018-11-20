package sitemap_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"path"
	"os"
	"bytes"
	"time"
	"runtime"
	"io/ioutil"
	"github.com/zirra-com/sitemap"
)

var sitemapNewsXML []byte
var sitemapXML   []byte

var _ = Describe("Sitemap#WriteTo", func() {
	It("The generated sitemap should be proper", func() {
		sm := sitemap.New()
		t := time.Unix(0, 0).UTC()
		sm.Add(&sitemap.URL{
			Loc:        "http://example.com/",
			LastMod:    &t,
			ChangeFreq: sitemap.Daily,
		})
		sm.WriteTo(os.Stdout)
		// Output:
		// <?xml version="1.0" encoding="UTF-8"?>
		// <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
		//   <url>
		//     <loc>http://example.com/</loc>
		//     <lastmod>1970-01-01T00:00:00Z</lastmod>
		//     <changefreq>daily</changefreq>
		//   </url>
		// </urlset>
	})

	It("The sitemap should be properly parsed", func() {
		sm := sitemap.New()

		_, err := sm.ReadFrom(bytes.NewReader(sitemapNewsXML))
		Expect(err).To(BeNil())

		Expect(len(sm.URLs)).To(Equal(34))
		// check the first item
		URL := sm.URLs[0]

		Expect(URL.Loc).To(Equal("https://techcrunch.com/2018/05/10/uber-to-pop-up-a-service-in-spains-costa-del-sol-in-time-for-summer/"))
		Expect(URL.News).To(Not(BeNil()))
		Expect(URL.News.Publication.Name).To(Equal("TechCrunch"))
		Expect(URL.News.Publication.Language).To(Equal("en"))
		Expect(URL.News.Title).To(Equal("Uber to pop up a service in Spain&#039;s Costa del Sol in time for summer"))
		Expect(URL.News.Genres).To(Equal("Blog"))

	})

	It("The sitemap should be properly parsed", func() {
		sm := sitemap.New()

		_, err := sm.ReadFrom(bytes.NewReader(sitemapXML))
		Expect(err).To(BeNil())

		Expect(len(sm.URLs)).To(Equal(100))
		// check the first item
		URL := sm.URLs[0]

		Expect(URL.Loc).To(Equal("http://www.finsmes.com/2018/05/oncosynergy-raises-series-a-funding-round.html"))
		Expect(URL.News).To(BeNil())
	})

})

func init() {
	_, filename, _, _ := runtime.Caller(1)
	sitemapXML, _ = ioutil.ReadFile(path.Join(path.Dir(filename), "sitemap.xml"))

	sitemapNewsXML, _ = ioutil.ReadFile(path.Join(path.Dir(filename), "sitemap-news.xml"))
}

