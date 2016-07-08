package sitemap

import (
	"os"
	"time"
)

func ExampleWriteTo() {
	sm := New()
	t := time.Unix(0, 0).UTC()
	sm.Add(&URL{
		Loc:        "http://example.com/",
		LastMod:    &t,
		ChangeFreq: Daily,
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
