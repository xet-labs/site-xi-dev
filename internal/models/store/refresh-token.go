package store

import "time"

type RefreshToken struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UID          uint      `gorm:"column:uid;index;not null" json:"uid"`
	TokenHash    string    `gorm:"column:token_hash;uniqueIndex;size:128;not null" json:"token_hash"`
	ExpiresAt    time.Time `gorm:"column:expires_at;not null" json:"expires_at"`
	Revoked      bool      `gorm:"column:revoked;default:false;not null" json:"revoked"`
	ReplacedById *uint     `gorm:"column:replaced_by_id" json:"replaced_by_id,omitempty"`
	UserAgent    *string   `gorm:"column:user_agent;size:255" json:"user_agent,omitempty"`
	IPAddress    *string   `gorm:"column:ip_address;size:45" json:"ip_address,omitempty"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	// Relation (preload when needed)
	User *User `gorm:"foreignKey:UID;references:ID" json:"-"`
}
