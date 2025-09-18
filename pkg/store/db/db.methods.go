package db

// import (
// 	"github.com/rs/zerolog/log"
// 	"gorm.io/gorm"
// )

// // Get returns the DB instance by name or default
// func (d *DbStore) GetCli(name ...string) *gorm.DB {
// 	d.mu.RLock()

// 	dbName := d.Cli
// 	if len(name) > 0 && name[0] != "" {
// 		dbName = name[0]
// 	}

// 	if db, ok := d.clis[dbName]; ok {
// 		d.mu.RUnlock()
// 		return db
// 	}

// 	log.Warn().Msgf("requested db '%s' not found", dbName)
// 	d.mu.RUnlock()
// 	return nil
// }

// // Set sets a DB by name
// func (d *DbStore) AddCli(name string, db *gorm.DB) {
// 	d.mu.Lock()
// 	defer d.mu.Unlock()
// 	d.clients[name] = db
// }

// SetDefault sets the default DB name
func (d *DbStore) SetDefault(name string) {
	d.CliProfile = name
}

// You can similarly add Redis setters/getters if needed
