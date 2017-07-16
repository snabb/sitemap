// Package sitemap provides tools for creating an XML sitemap and writing it
// to an io.Writer (such as http.ResponseWriter). Please see
// http://www.sitemaps.org/ for description of sitemap contents.
package sitemap

import (
	"encoding/xml"
	"github.com/snabb/diagio"
	"io"
	"time"
)

type ChangeFreq string

// Feel free to use these constants for ChangeFreq (or you can just supply
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

// Single URL entry in sitemap. LastMod is a pointer to time.Time because
// omitempty does not work otherwise. Loc is the only mandatory item.
type URL struct {
	Loc        string     `xml:"loc"`
	LastMod    *time.Time `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreq `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`

	URLs []*URL `xml:"url"`
}

// New returns a new Sitemap.
func New() *Sitemap {
	return &Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  make([]*URL, 0),
	}
}

// Add adds an URL to a Sitemap.
func (s *Sitemap) Add(u *URL) {
	s.URLs = append(s.URLs, u)
}

// WriteTo writes XML encoded sitemap to given io.Writer.
// Implements io.WriterTo.
func (s *Sitemap) WriteTo(w io.Writer) (n int64, err error) {
	cw := diagio.NewCounterWriter(w)

	_, err = cw.Write([]byte(xml.Header))
	if err != nil {
		return cw.Count(), err
	}
	en := xml.NewEncoder(cw)
	en.Indent("", "  ")
	err = en.Encode(s)
	cw.Write([]byte{'\n'})
	return cw.Count(), err
}

var _ io.WriterTo = (*Sitemap)(nil)

// ReadFrom reads and parses an XML encoded sitemap from io.Reader.
// Implements io.ReaderFrom.
func (s *Sitemap) ReadFrom(r io.Reader) (n int64, err error) {
	de := xml.NewDecoder(r)
	err = de.Decode(s)
	return de.InputOffset(), err
}

var _ io.ReaderFrom = (*Sitemap)(nil)
