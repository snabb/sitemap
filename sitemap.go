// Package sitemap provides tools for creating XML sitemaps
// and sitemap indexes and writing them to [io.Writer] (such as
// [net/http.ResponseWriter]).
//
// Please see https://www.sitemaps.org/ for description of sitemap contents.
package sitemap

import (
	"encoding/xml"
	"io"
	"time"

	"github.com/snabb/diagio"
)

// ChangeFreq specifies change frequency of a [Sitemap] or [SitemapIndex]
// [URL] entry. It is just a string.
type ChangeFreq string

// Feel free to use these constants for [ChangeFreq] (or you can just supply
// a string directly).
const (
	Always  ChangeFreq = "always"
	Hourly  ChangeFreq = "hourly"
	Daily   ChangeFreq = "daily"
	Weekly  ChangeFreq = "weekly"
	Monthly ChangeFreq = "monthly"
	Yearly  ChangeFreq = "yearly"
	Never   ChangeFreq = "never"
)

// XHTMLLink entry in [URL].
type XHTMLLink struct {
	Rel      string `xml:"rel,attr"`
	HrefLang string `xml:"hreflang,attr"`
	Href     string `xml:"href,attr"`
}

// URL entry in [Sitemap] or [SitemapIndex]. LastMod is a pointer
// to [time.Time] because omitempty does not work otherwise. Loc is the
// only mandatory item. ChangeFreq and Priority must be left empty when
// using with a sitemap index.
type URL struct {
	Loc        string      `xml:"loc"`
	LastMod    *time.Time  `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreq  `xml:"changefreq,omitempty"`
	Priority   float32     `xml:"priority,omitempty"`
	XHTMLLinks []XHTMLLink `xml:"xhtml:link"`
}

// Sitemap represents a complete sitemap which can be marshaled to XML.
// New instances must be created with [New] in order to set the xmlns
// attribute correctly. Minify can be set to make the output less human
// readable.
type Sitemap struct {
	XMLName    xml.Name `xml:"urlset"`
	Xmlns      string   `xml:"xmlns,attr"`
	XmlnsXHTML *string  `xml:"xmlns:xhtml,attr"`

	URLs []*URL `xml:"url"`

	Minify bool `xml:"-"`
}

// New returns a new [Sitemap].
func New() *Sitemap {
	return &Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  make([]*URL, 0),
	}
}

// WithXHTML adds the xmlns:xhtml to a [Sitemap].
func (s *Sitemap) WithXHTML() *Sitemap {
	xhtml := "http://www.w3.org/1999/xhtml"
	s.XmlnsXHTML = &xhtml
	return s
}

// WithMinify enables minification on a [Sitemap].
func (s *Sitemap) WithMinify() *Sitemap {
	s.Minify = true
	return s
}

// Add adds an [URL] to a [Sitemap].
func (s *Sitemap) Add(u *URL) {
	s.URLs = append(s.URLs, u)
}

// WriteTo writes XML encoded sitemap to given [io.Writer].
// Implements [io.WriterTo].
func (s *Sitemap) WriteTo(w io.Writer) (n int64, err error) {
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
	if err != nil {
		return cw.Count(), err
	}
	_, err = cw.Write([]byte{'\n'})
	return cw.Count(), err
}

var _ io.WriterTo = (*Sitemap)(nil)

// ReadFrom reads and parses an XML encoded sitemap from [io.Reader].
// Implements [io.ReaderFrom]. Due to https://github.com/golang/go/issues/9519,
// unmarshaling xhtml links doesn't work.
func (s *Sitemap) ReadFrom(r io.Reader) (n int64, err error) {
	de := xml.NewDecoder(r)
	err = de.Decode(s)
	return de.InputOffset(), err
}

var _ io.ReaderFrom = (*Sitemap)(nil)
