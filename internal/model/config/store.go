package config

type StoreConf struct {
	Db  DbStore  `json:"db,omitempty"`
	Rdb RdbStore `json:"rdb,omitempty"`
}

type DbStore struct {
	DefaultProfile string                 `json:"default_profile,omitempty"`
	Conn           map[string]ConnProfile `json:"conn,omitempty"`
}
type RdbStore struct {
	DefaultProfile string                 `json:"default_profile,omitempty"`
	Prefix         string                 `json:"prefix,omitempty"`
	Conn           map[string]ConnProfile `json:"conn,omitempty"`
}

type ConnProfile struct {
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
