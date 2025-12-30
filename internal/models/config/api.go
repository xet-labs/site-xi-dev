package config

type ApiConf struct {
	CookieDomain  string `json:"cookie_domain,omitempty"`
	SecureCookies bool   `json:"secure_cookies,omitempty"`
}
