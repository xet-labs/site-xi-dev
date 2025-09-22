package db

import (
	"errors"
	"fmt"
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

func (d *DbStore) RawCli() *gorm.DB { return d.cli }

func (d *DbStore) Cli(cliProfiles ...string) *gorm.DB {
	// Fast path: return default if no profiles provided
	if len(cliProfiles) == 0 {
		if d.cli != nil {
			return d.cli
		}
		return &gorm.DB{Error: fmt.Errorf("DbStore: no database connection available")}
	}

	// Check profiles under a single read lock
	d.mu.RLock()
	defer d.mu.RUnlock()

	for _, profile := range cliProfiles {
		if cli, ok := d.clis[profile]; ok && cli != nil {
			return cli
		}
	}

	// Fallback to default if profile not found
	if d.cli != nil {
		return d.cli
	}

	// No DB found — return dummy
	return &gorm.DB{Error: fmt.Errorf("DbStore: no database connection available")}
}
