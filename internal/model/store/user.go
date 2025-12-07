package store

import (
	"time"
)

type User struct {
	ID            uint64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Username      string     `gorm:"column:username;type:varchar(20);uniqueIndex;not null" json:"username"`
	Name          string     `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Email         string     `gorm:"column:email;type:varchar(50);uniqueIndex;not null" json:"email"`
	Verified      bool       `gorm:"column:verified;type:tinyint(1);not null;default:0" json:"verified"`
	PasswordHash  string     `gorm:"column:password_hash;type:varchar(255);not null" json:"-"`
	Role          string     `gorm:"column:role;type:enum('admin','dev','user');default:'user'" json:"role"`
	CreatedAt     time.Time  `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
	LastLogin     *time.Time `gorm:"column:last_login;type:datetime" json:"last_login,omitempty"`
	Status        string     `gorm:"column:status;type:enum('active','inactive','suspended','deleted');default:'active'" json:"status"`
	AvatarURL     *string    `gorm:"column:avatar_url;type:varchar(255)" json:"avatar_url,omitempty"`
	Address       *string    `gorm:"column:address;type:text" json:"address,omitempty"`
	PhoneNo       *string    `gorm:"column:phone_no;type:varchar(20)" json:"phone_no,omitempty"`
	DOB           *time.Time `gorm:"column:dob;type:date" json:"dob,omitempty"`
	Config        *string    `gorm:"column:config;type:longtext;check:json_valid(config)" json:"config,omitempty"`
	RememberToken *string    `gorm:"column:remember_token;type:varchar(100)" json:"remember_token,omitempty"`

	// Relationship: one user has many blogs
	RefreshTokens []RefreshToken `gorm:"foreignKey:UID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"auth_tokens,omitempty"`
	Blogs         []Blog         `gorm:"foreignKey:UID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"blogs,omitempty"`
}

// TableName overrides default pluralization
func (User) TableName() string {
	return "users"
}
