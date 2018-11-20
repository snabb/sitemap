// Package sitemap provides tools for creating XML sitemaps
// and sitemap indexes and writing them to io.Writer (such as
// http.ResponseWriter).
//
// Please see http://www.sitemaps.org/ for description of sitemap contents.
package sitemap

import (
	"encoding/xml"
	"github.com/snabb/diagio"
	"io"
	"time"
)

// ChangeFreq specifies change frequency of a sitemap entry. It is just a string.
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

// Publication for new
type Publication struct {
	Name     string `xml:"name,omitempty"` // Name of the news publication. It must exactly match the name as it appears on your articles in news.google.com, omitting any trailing parentheticals. For example, if the name appears in Google News as "The Example Times (subscription)", you should use "The Example Times". Required.
	Language string `xml:"language"`  // Language of the publication. It should be an ISO 639 Language Code (either 2 or 3 letters); see: http://www.loc.gov/standards/iso639-2/php/code_list.php Exception: For Chinese, please use zh-cn for Simplified Chinese or zh-tw for Traditional Chinese. Required.
}

// News entry following the format 
type News struct {
	Publication     Publication `xml:"publication"`  // The publication in which the article appears. Required.
	Title           string      `xml:"title,omitempty"`   // Title of the news article.
	PublicationDate *time.Time  `xml:"publication_date,omitempty"` // Article publication date in W3C format, specifying the complete date (YYYY-MM-DD) with optional timestamp. See: http://www.w3.org/TR/NOTE-datetime Please ensure that you give the original date and time at which the article was published on your site; do not give the time at which the article was added to your Sitemap. Required.
	Genres          string      `xml:"genres,omitempty"` // A comma-separated list of properties characterizing the content of the article, such as "PressRelease" or "UserGenerated". For a list of possible values, see: http://www.google.com/support/news_pub/bin/answer.py?answer=93992 Required if any genres apply to the article, otherwise this tag should be omitted.
	Keywords        string      `xml:"keywords,omitempty"`  // Comma-separated list of keywords describing the topic of the article. Keywords may be drawn from, but are not limited to, the list of existing Google News keywords; see: http://www.google.com/support/news_pub/bin/answer.py?answer=116037 Optional.
	StockTickers    string      `xml:"stock_tickers,omitempty"`  // Comma-separated list of up to 5 stock tickers of the companies, mutual funds, or other financial entities that are the main subject of the article. Relevant primarily for business articles. Each ticker must be prefixed by the name of its stock exchange, and must match its entry in Google Finance. For example, "NASDAQ:AMAT" (but not "NASD:AMAT"), or "BOM:500325" (but not "BOM:RIL"). Optional.
}

// URL entry in sitemap or sitemap index. LastMod is a pointer
// to time.Time because omitempty does not work otherwise. Loc is the
// only mandatory item. ChangeFreq and Priority must be left empty when
// using with a sitemap index.
type URL struct {
	Loc        string     `xml:"loc"`
	LastMod    *time.Time `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreq `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
	News       *News      `xml:"news,omitempty"`
}

// Sitemap represents a complete sitemap which can be marshaled to XML.
// New instances must be created with New() in order to set the xmlns
// attribute correctly. Minify can be set to make the output less human
// readable.
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`

	URLs []*URL `xml:"url"`

	Minify bool `xml:"-"`
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
	if !s.Minify {
		en.Indent("", "  ")
	}
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
