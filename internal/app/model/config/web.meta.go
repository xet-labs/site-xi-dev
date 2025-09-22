package config

import "time"

type WebMeta struct {
	// Basic SEO
	Alternate         []string       `json:"alternate,omitempty"` // hreflang alternate URLs
	AltJson           string         `json:"alt_json,omitempty"`  // Json data url for page (api or similar)
	Canonical         string         `json:"canonical,omitempty"`
	CreatedAt         *time.Time     `json:"created_at,omitempty"`
	UpdatedAt         *time.Time     `json:"updated_at,omitempty"`
	Description       string         `json:"description,omitempty"`
	HrefLang          []MetaHrefLang `json:"href_lang,omitempty"`
	Img               MetaImg        `json:"img,omitempty"`
	Locale            string         `json:"locale,omitempty"`   // en_US etc (for og:locale)
	Robots            string         `json:"robots,omitempty"`   // e.g., "index, follow"
	Referrer          string         `json:"referrer,omitempty"` // default "no-referrer-when-downgrade"
	ShortLink         string         `json:"short_link,omitempty"`
	Tagline           string         `json:"tagline,omitempty"`
	Tags              []string       `json:"tags,omitempty"`
	Title             string         `json:"title,omitempty"`
	TitleOrgSuffixInc *bool          `json:"title_org_suffix_inc,omitempty"`
	TitleOrgSuffixSep string         `json:"title_org_suffix_sep,omitempty"`
	Type              string         `json:"type,omitempty"` // WebSite, WebPage, Article, BlogPosting, NewsArticle, Product, Offer, Person, Organization, FAQPage
	URL               string         `json:"url,omitempty"`

	// Sitemap
	Sitemap MetaSitemap `json:"sitemap,omitempty"`

	// Social/owners
	Author  MetaAuthor  `json:"author,omitempty"`
	OG      MetaOG      `json:"og,omitempty"`
	Twitter MetaTwitter `json:"twitter,omitempty"`

	// Article/Product specifics
	Category    string `json:"category,omitempty"`
	IsFree      *bool  `json:"isAccessibleForFree,omitempty"`
	ReadingTime string `json:"readingTime,omitempty"` // e.g., "5 min"

	// Extra JSON-LD (raw block to merge)
	LD     map[string]any `json:"ld,omitempty"`
	LDPre  map[string]any `json:"ld_pre,omitempty"`
	LDPost map[string]any `json:"ld_post,omitempty"`
}

type MetaAuthor struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	Img         *string `json:"img,omitempty"`
	JobTitle    string `json:"jobTitle,omitempty"`
	Description string `json:"description,omitempty"`
	SameAs      string `json:"sameAs,omitempty"` // single URL or CSV
}
type MetaHrefLang struct {
	Lang string `json:"lang,omitempty"` // e.g. en, en-IN, fr
	URL  string `json:"url,omitempty"`
}
type MetaImg struct {
	URL string `json:"url,omitempty"`
	Alt string `json:"alt,omitempty"`
}
type MetaOG struct {
	Type        string            `json:"type,omitempty"`
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Img         string            `json:"img,omitempty"`
	URL         string            `json:"url,omitempty"`
	Extra       map[string]string `json:"extra,omitempty"`
}
type MetaSitemap struct {
	Loc        string `json:"loc,omitempty" xml:"loc"`
	LastMod    string `json:"lastmod,omitempty" xml:"lastmod,omitempty"`
	ChangeFreq string `json:"change_freq,omitempty" xml:"changefreq,omitempty"`
	Priority   string `json:"priority,omitempty" xml:"priority,omitempty"`
}
type MetaTwitter struct {
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Site        string            `json:"site,omitempty"`    // @handle
	Creator     string            `json:"creator,omitempty"` // @author
	Card        string            `json:"card,omitempty"`    // summary, summary_large_image
	Img         string            `json:"img,omitempty"`
	Extra       map[string]string `json:"extra,omitempty"` // label1/data1... or any kv
}
