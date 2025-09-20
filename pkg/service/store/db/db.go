package db

import (
	"errors"
	"sync"
	"xi/pkg/lib/cfg"

	"gorm.io/gorm"
)

type DbStore struct {
	cli        *gorm.DB
	cliProfile string
	clis       map[string]*gorm.DB
	mu         sync.RWMutex
}

var (
	Db = &DbStore{clis: make(map[string]*gorm.DB)}
)

// AddCli stores a DB instance by its cliProfile
func (d *DbStore) AddCli(cliProfile string, cli *gorm.DB) error {
	if cli == nil {
		return errors.New("db cli is nil for profile '" + cliProfile + "'")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	d.clis[cliProfile] = cli

	// Set as global if this is the default profile OR if global isn't set yet
	if cfg.Store.Db.DefaultProfile == cliProfile || d.cli == nil {
		d.cli = cli
		d.cliProfile = cliProfile
	}
	return nil
}

// SetCli explicitly sets the current global DB client
func (d *DbStore) SetCli(cliProfile string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if cli, ok := d.clis[cliProfile]; ok {
		d.cli = cli
		d.cliProfile = cliProfile
	}
}

func (d *DbStore) Cli(cliProfiles ...string) *gorm.DB {
	// Fast path: no profile given, return default directly
	if len(cliProfiles) == 0 {
		return d.cli
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	// Try profiles in order
	for _, profile := range cliProfiles {
		if cli, ok := d.clis[profile]; ok && cli != nil {
			return cli
		}
	}

	// Fallback to default
	return d.cli
}