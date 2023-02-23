package sitemap_test

import (
	"os"
	"time"

	"github.com/snabb/sitemap"
)

// Sitemap index with one sitemap URL.
func ExampleSitemapIndex() {
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
}
