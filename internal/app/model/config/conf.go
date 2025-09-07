package config

type Config struct {
	Api   ApiConf   `json:"api"`
	App   AppConf   `json:"app"`
	Org   OrgConf   `json:"org"`
	Build BuildConf `json:"build"`
	Db    DbConf    `json:"db"`
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

type BuildConf struct {
	Date     string `json:"date,omitempty"`
	Name     string `json:"name,omitempty"`
	Revision string `json:"revision,omitempty"`
	Version  string `json:"version,omitempty"`
	Mode     string `json:"mode,omitempty"`
}
