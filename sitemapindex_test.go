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

var sitemapIndex []byte

var _ = Describe("SitemapIndex#WriteTo", func() {
	It("The generated sitemap index should be proper", func() {
		smi := sitemap.NewSitemapIndex()
		t := time.Unix(0, 0).UTC()
		smi.Add(&sitemap.URL{
			Loc:     "http://example.com/sitemap-1.xml",
			LastMod: &t,
		})
		smi.WriteTo(os.Stdout)
		// Output:
		// <?xml version="1.0" encoding="UTF-8"?>
		// <sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
		//   <sitemap>
		//     <loc>http://example.com/sitemap-1.xml</loc>
		//     <lastmod>1970-01-01T00:00:00Z</lastmod>
		//   </sitemap>
		// </sitemapindex>
	})

	It("The sitemap index should be properly parsed", func() {
		smi := sitemap.NewSitemapIndex()

		_, err := smi.ReadFrom(bytes.NewReader(sitemapIndex))
		Expect(err).To(BeNil())

		Expect(len(smi.URLs)).To(Equal(1))
		// check the first item
		URL := smi.URLs[0]

		Expect(URL.Loc).To(Equal("http://www.finsmes.com/sitemap-pt-post-2018-05.xml"))
		Expect(URL.News).To(BeNil())
	})

})


func init() {
	_, filename, _, _ := runtime.Caller(1)
	sitemapIndex, _ = ioutil.ReadFile(path.Join(path.Dir(filename), "sitemapindex.xml"))
}


