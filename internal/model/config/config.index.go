package config

type Config struct {
	Api   ApiConf   `json:"api"`
	App   AppConf   `json:"app"`
	Org   OrgConf   `json:"org"`
	Store StoreConf `json:"store"`
	Web   WebConf   `json:"web"`
}

type OrgConf struct {
	Abbr    string   `json:"abbr,omitempty"`
	Name    string   `json:"name,omitempty"`
	Domain  string   `json:"domain,omitempty"`
	URL     string   `json:"url,omitempty"`
	Logo    []string `json:"logo"`
	Img     []string `json:"img,omitempty"`
	Tagline string   `json:"tagline,omitempty"`
}
