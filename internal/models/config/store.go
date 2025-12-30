package config

type StoreConf struct {
	Db  DbStore  `json:"db"`
	Rdb RdbStore `json:"rdb"`
}

type DbStore struct {
	Enable         bool                   `json:"enable"`
	DefaultProfile string                 `json:"default_profile,omitempty"`
	Conn           map[string]ConnProfile `json:"conn,omitempty"`
}
type RdbStore struct {
	Enable         bool                   `json:"enable"`
	DefaultProfile string                 `json:"default_profile,omitempty"`
	Conn           map[string]ConnProfile `json:"conn,omitempty"`
	Prefix         string                 `json:"prefix,omitempty"`
}

type ConnProfile struct {
	Enable        bool   `json:"enable"`
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
