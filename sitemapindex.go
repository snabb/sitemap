package sitemap

import (
	"encoding/xml"
	"io"

	"github.com/snabb/diagio"
)

// SitemapIndex is like Sitemap except the elements are named differently
// (and ChangeFreq and Priority may not be used).
// New instances must be created with NewSitemapIndex() in order to set the
// xmlns attribute correctly. Minify can be set to make the output less
// human readable.
type SitemapIndex struct {
	XMLName xml.Name `xml:"sitemapindex"`
	Xmlns   string   `xml:"xmlns,attr"`

	URLs []*URL `xml:"sitemap"`

	Minify bool `xml:"-"`
}

// NewSitemapIndex returns new SitemapIndex.
func NewSitemapIndex() *SitemapIndex {
	return &SitemapIndex{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  make([]*URL, 0),
	}
}

// Add adds an URL to a SitemapIndex.
func (s *SitemapIndex) Add(u *URL) {
	s.URLs = append(s.URLs, u)
}

// WriteTo writes XML encoded sitemap index to given io.Writer.
// Implements io.WriterTo.
func (s *SitemapIndex) WriteTo(w io.Writer) (n int64, err error) {
	cw := diagio.NewCounterWriter(w)

	_, err = cw.Write([]byte(xml.Header))
	if err != nil {
		return cw.Count(), err
	}
	en := xml.NewEncoder(cw)
	if !s.Minify {
		en.Indent("", "  ")
	}
	err = en.Encode(s)
	cw.Write([]byte{'\n'})
	return cw.Count(), err
}

var _ io.WriterTo = (*Sitemap)(nil)

// ReadFrom reads and parses an XML encoded sitemap index from io.Reader.
// Implements io.ReaderFrom.
func (s *SitemapIndex) ReadFrom(r io.Reader) (n int64, err error) {
	de := xml.NewDecoder(r)
	err = de.Decode(s)
	return de.InputOffset(), err
}

var _ io.ReaderFrom = (*Sitemap)(nil)
