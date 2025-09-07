package config

type WebConf struct {
	CssDir       string           `json:"css_dir,omitempty"`
	CssDirs      []string         `json:"css_dirs,omitempty"`
	CssBaseDir   string           `json:"css_base_dir,omitempty"`
	TemplateDir  string           `json:"template_dir,omitempty"`
	TemplateDirs []string         `json:"template_dirs,omitempty"`
	Default      *Page            `json:"default,omitempty"`
	Pages        map[string]*Page `json:"pages,omitempty"`
}

type Page struct {
	Route   string  `json:"route,omitempty"`
	Ctrl    Ctrl    `json:"ctrl,omitempty"`
	Content Content `json:"content,omitempty"`
	Org     Org     `json:"org,omitempty"`
	Meta    Meta    `json:"meta,omitempty"`
	Web     Web     `json:"web,omitempty"`

	NavMenu []NavMenu      `json:"nav_menu,omitempty"`
	Css     []string       `json:"css,omitempty"`
	Js      []string       `json:"js,omitempty"`
	Js99    []string       `json:"js99,omitempty"`
	LibHead []string       `json:"lib_head,omitempty"`
	Lib     []string       `json:"lib,omitempty"`
	Lib99   []string       `json:"lib99,omitempty"`
	Extra   map[string]any `json:"extra,omitempty"`
	Rt      map[string]any `json:"_runtime,omitempty"` // Runtime data
}
type Ctrl struct {
	Cache  *bool  `json:"cache,omitempty"`
	Layout string `json:"layout,omitempty"`
	Mode   string `json:"mode,omitempty"`   // if {true || null} route will be handled by routes.managed
	Method string `json:"method,omitempty"` // GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS || fallback to GET
	Render string `json:"render,omitempty"`
}
type Content struct {
	Raw  string `json:"raw,omitempty"`
	File string `json:"file,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Web struct {
	Menu              string `json:"menu,omitempty"`
	SubBrand          string `json:"sub_brand,omitempty"`
	SubBrandSuffixInc *bool  `json:"sub_brand_suffix_inc,omitempty"`
	SubBrandSuffixSep string `json:"sub_brand_suffix_sep,omitempty"`
}
type NavMenu struct {
	Type  string `json:"type,omitempty"` // Button, Link,
	Label string `json:"label,omitempty"`
	Href  string `json:"href,omitempty"`
	URL   string `json:"url,omitempty"`
	Data  string `json:"data,omitempty"`
}
type Org struct {
	AltName string `json:"alt_name,omitempty"`
	Name    string `json:"name,omitempty"`
	Domain  string `json:"domain,omitempty"`
	URL     string `json:"url,omitempty"`
	Logo    string `json:"logo,omitempty"`
	Img     Img    `json:"img,omitempty"`
	Tagline string `json:"tagline,omitempty"`
}
