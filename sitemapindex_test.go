package sitemap_test

import (
	"github.com/snabb/sitemap"
	"os"
	"time"
)

func ExampleSitemapIndex() {
	smi := sitemap.NewSitemapIndex()
	t := time.Unix(0, 0).UTC()
	smi.Add(&sitemap.URL{
		Loc:     "http://example.com/",
		LastMod: &t,
	})
	smi.WriteTo(os.Stdout)
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	//   <sitemap>
	//     <loc>http://example.com/</loc>
	//     <lastmod>1970-01-01T00:00:00Z</lastmod>
	//   </sitemap>
	// </sitemapindex>
}
