package config

type DbConf struct {
	DbDefault  string             `json:"db_default,omitempty"`
	RdbDefault string             `json:"rdb_default,omitempty"`
	RdbPrefix  string             `json:"rdb_prefix,omitempty"`
	Conn       map[string]DbProfile `json:"conn,omitempty"`
}
type DbProfile struct {
	Enable        bool   `json:"enable,omitempty"`
	Db            string `json:"db,omitempty"`
	Rdb           int    `json:"rdb,omitempty"`
	User          string `json:"user,omitempty"`
	Pass          string `json:"pass"`
	Driver        string `json:"driver,omitempty"`
	Host          string `json:"host,omitempty"`
	Port          string `json:"port,omitempty"`
	Engine        string `json:"engine,omitempty"`
	Socket        string `json:"socket,omitempty"`
	Charset       string `json:"charset,omitempty"`
	Collation     string `json:"collation,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
	PrefixIndexes bool   `json:"prefixindexes,omitempty"`
	Strict        bool   `json:"strict,omitempty"`
}
