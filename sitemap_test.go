package sitemap_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/snabb/sitemap"
)

// This is a web server that implements two request paths /foo and /bar
// and provides a sitemap that contains those paths at /sitemap.xml.
func Example() {
	sm := sitemap.New()

	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "foo")
	})
	sm.Add(&sitemap.URL{Loc: "http://localhost:8080/foo"})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "bar")
	})
	sm.Add(&sitemap.URL{Loc: "http://localhost:8080/bar"})

	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		sm.WriteTo(w)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Sitemap with one URL.
func ExampleSitemap() {
	sm := sitemap.New()
	t := time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC)
	sm.Add(&sitemap.URL{
		Loc:        "http://example.com/",
		LastMod:    &t,
		ChangeFreq: sitemap.Daily,
		Priority:   0.5,
	})
	sm.WriteTo(os.Stdout)
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	//   <url>
	//     <loc>http://example.com/</loc>
	//     <lastmod>1984-01-01T00:00:00Z</lastmod>
	//     <changefreq>daily</changefreq>
	//     <priority>0.5</priority>
	//   </url>
	// </urlset>
}

// Setting Minify to true omits indentation and newlines in generated sitemap.
func ExampleSitemap_minify() {
	sm := sitemap.New()
	sm.Minify = true
	t := time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC)
	sm.Add(&sitemap.URL{
		Loc:        "http://example.com/",
		LastMod:    &t,
		ChangeFreq: sitemap.Weekly,
		Priority:   0.5,
	})
	sm.WriteTo(os.Stdout)
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"><url><loc>http://example.com/</loc><lastmod>1984-01-01T00:00:00Z</lastmod><changefreq>weekly</changefreq><priority>0.5</priority></url></urlset>
}
