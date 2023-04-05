package sitemap_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-test/deep"
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
	sm := sitemap.New().WithXHTML()
	t := time.Date(1984, 1, 1, 0, 0, 0, 0, time.UTC)
	sm.Add(&sitemap.URL{
		Loc:        "http://example.com/",
		LastMod:    &t,
		ChangeFreq: sitemap.Daily,
		Priority:   0.5,
		XHTMLLinks: []sitemap.XHTMLLink{
			{Rel: "alternate", HrefLang: "en", Href: "http://example.com"},
			{Rel: "alternate", HrefLang: "fr", Href: "http://example.com?lang=fr"},
		},
	})
	sm.WriteTo(os.Stdout)
	// Output:
	// <?xml version="1.0" encoding="UTF-8"?>
	// <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">
	//   <url>
	//     <loc>http://example.com/</loc>
	//     <lastmod>1984-01-01T00:00:00Z</lastmod>
	//     <changefreq>daily</changefreq>
	//     <priority>0.5</priority>
	//     <xhtml:link rel="alternate" hreflang="en" href="http://example.com"></xhtml:link>
	//     <xhtml:link rel="alternate" hreflang="fr" href="http://example.com?lang=fr"></xhtml:link>
	//   </url>
	// </urlset>
}

// WithMinify omits indentation and newlines in generated sitemap.
func ExampleSitemap_minify() {
	sm := sitemap.New().WithMinify()
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

// failWriter is a Writer that always fails.
type failWriter struct{}

func (failWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("write failure")
}

var _ io.Writer = (*failWriter)(nil)

func TestSitemap_WriteToError(t *testing.T) {
	sm := sitemap.New()
	sm.Add(&sitemap.URL{Loc: "http://example.com/"})

	n, err := sm.WriteTo(failWriter{})
	if n != 0 {
		t.Error("WriteTo did not return zero")
	}
	if err == nil {
		t.Error("WriteTo did not propagate error")
	}
}

func TestSitemap_ReadFrom(t *testing.T) {
	sm1 := sitemap.New()

	for i := 0; i < rand.Intn(100)+1; i++ {
		timeNow := time.Now()
		sm1.Add(&sitemap.URL{
			Loc:        fmt.Sprintf("http://example.com/%03d.html", i),
			LastMod:    &timeNow,
			ChangeFreq: sitemap.Always,
			Priority:   rand.Float32(),
		})
	}

	buf := new(bytes.Buffer)

	_, err := sm1.WriteTo(buf)
	if err != nil {
		t.Fatalf("Error writing sitemap: %v", err)
	}

	sm2 := new(sitemap.Sitemap)

	_, err = sm2.ReadFrom(buf)
	if err != nil {
		t.Fatalf("Error reading sitemap: %v", err)
	}

	if diff := deep.Equal(sm1.URLs, sm2.URLs); diff != nil {
		t.Error(diff)
	}
}
