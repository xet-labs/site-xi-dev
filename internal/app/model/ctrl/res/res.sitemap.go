package model

import (
	"encoding/xml"
	model_config "xi/internal/app/model/config"
)

// Sitemap is the root <urlset> of the XML sitemap.
type Sitemap struct {
	XMLName xml.Name                   `xml:"urlset"`
	Xmlns   string                     `xml:"xmlns,attr"`
	URLs    []model_config.MetaSitemap `xml:"url"`
}
