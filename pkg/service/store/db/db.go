package db

import (
	"errors"
	"sync"
	"xi/pkg/lib/cfg"

	"gorm.io/gorm"
)

type DbStore struct {
	Cli        *gorm.DB
	CliProfile string
	clis       map[string]*gorm.DB
	mu         sync.RWMutex
}

var Db = &DbStore{
	clis: make(map[string]*gorm.DB),
}

// AddCli stores a DB instance by its cliProfile
func (d *DbStore) AddCli(cliProfile string, cli *gorm.DB) error {
	if cli == nil {
		return errors.New("db cli is nil for profile '" + cliProfile + "'")
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	d.clis[cliProfile] = cli

	// Set as global if this is the default profile OR if global isn't set yet
	if cfg.Store.Db.DefaultProfile == cliProfile || d.Cli == nil {
		d.Cli = cli
		d.CliProfile = cliProfile
	}
	return nil
}

// SetCli explicitly sets the current global DB client
func (d *DbStore) SetCli(cliProfile string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if cli, ok := d.clis[cliProfile]; ok {
		d.Cli = cli
		d.CliProfile = cliProfile
	}
}

// GetCli returns the DB instance by profile(s) or falls back to global
func (d *DbStore) GetCli(cliProfiles ...string) *gorm.DB {
	d.mu.RLock()
	defer d.mu.RUnlock()

	// Try profiles in order
	for _, profile := range cliProfiles {
		if cli, ok := d.clis[profile]; ok && cli != nil {
			return cli
		}
	}

	// Fall back to global
	if d.Cli != nil {
		return d.Cli
	}

	// Nothing found
	return nil
}
