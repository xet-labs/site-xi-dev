package config

type AuthConf struct {
	JwtSecret string `json:"jwt_secret,omitempty"`
}
