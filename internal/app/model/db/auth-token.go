package db

import "time"

type RefreshToken struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UID          uint      `gorm:"column:uid;index" json:"uid"`
	Revoked      bool      `gorm:"column:revoked" json:"revoked"`
	TokenHash    string    `gorm:"column:refresh_token;uniqueIndex;size:128" json:"token_hash"`
	ExpiresAt    time.Time `gorm:"column:expires_at" json:"expires_at"`
	ReplacedById uint      `gorm:"column:replaced_by_id" json:"replaced_by_id"`
	UserAgent    string    `gorm:"column:user_agent;size:255" json:"user_agent,omitempty"`
	IPAddress    string    `gorm:"column:ip_address;size:45" json:"ip_address,omitempty"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relation
	User *User `gorm:"foreignKey:UID;references:ID" json:"user,omitempty"`
}
