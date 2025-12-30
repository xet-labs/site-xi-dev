package config

type WebCache struct {
	Css CacheCss `json:"css,omitempty"`
}

type CacheCss struct {
	FilesList bool `json:"files_list,omitempty"`
}
