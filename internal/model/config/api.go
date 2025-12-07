package config

type ApiConf struct {
	JwtSecret     string `json:"jwt_secret,omitempty"`
	CookieDomain  string `json:"cookie_domain,omitempty"`
	SecureCookies bool   `json:"secure_cookies,omitempty"`
}
