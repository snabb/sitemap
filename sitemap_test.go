package sitemap_test

import (
	"github.com/snabb/sitemap"
	"os"
	"time"
)

func Example() {
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
}
