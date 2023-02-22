package sitemap_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/go-test/deep"
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

func TestSitemapIndex_WriteToError(t *testing.T) {
	smi := sitemap.NewSitemapIndex()
	smi.Add(&sitemap.URL{Loc: "http://example.com/sitemap.xml"})

	n, err := smi.WriteTo(failWriter{})
	if n != 0 {
		t.Error("WriteTo did not return zero")
	}
	if err == nil {
		t.Error("WriteTo did not propagate error")
	}
}

func TestSitemapIndex_ReadFrom(t *testing.T) {
	smi1 := sitemap.NewSitemapIndex()

	for i := 0; i < rand.Intn(100)+1; i++ {
		timeNow := time.Now()
		smi1.Add(&sitemap.URL{
			Loc:     fmt.Sprintf("http://example.com/sitemap-%03d.xml", i),
			LastMod: &timeNow,
		})
	}

	buf := new(bytes.Buffer)

	_, err := smi1.WriteTo(buf)
	if err != nil {
		t.Fatalf("Error writing sitemap: %v", err)
	}

	smi2 := new(sitemap.SitemapIndex)

	_, err = smi2.ReadFrom(buf)
	if err != nil {
		t.Fatalf("Error reading sitemap: %v", err)
	}

	if diff := deep.Equal(smi1.URLs, smi2.URLs); diff != nil {
		t.Error(diff)
	}
}
