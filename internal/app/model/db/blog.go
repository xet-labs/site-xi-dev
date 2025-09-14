package db

import (
	"time"
	"xi/internal/app/model/util"
)

type Blog struct {
	ID           uint64           `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UID          uint64           `gorm:"column:uid;not null;index" json:"uid"`
	Status       string           `gorm:"column:status;type:enum('draft','published','published_hidden','archived');default:'draft'" json:"status"`
	Tags         util.StringArray `gorm:"column:tags;type:varchar(255)" json:"tags,omitempty"` // could also map to util.StringArray if you handle conversion
	Title        string           `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Headline     string           `gorm:"column:short_title;type:varchar(255)" json:"headline,omitempty"`
	Description  string           `gorm:"column:description;type:text" json:"description,omitempty"`
	FeaturedImg  string           `gorm:"column:featured_img;type:varchar(255);not null" json:"featured_img"`
	Slug         string           `gorm:"column:slug;type:varchar(255);not null" json:"slug"`
	Path         string           `gorm:"column:path;type:varchar(255)" json:"path,omitempty"`
	CreatedAt    *time.Time       `gorm:"column:created_at;type:timestamp(3);default:CURRENT_TIMESTAMP(3)" json:"created_at,omitempty"`
	UpdatedAt    *time.Time       `gorm:"column:updated_at;type:timestamp(3);default:CURRENT_TIMESTAMP(3);autoUpdateTime" json:"updated_at,omitempty"`
	Content      string           `gorm:"column:content;type:longtext" json:"content,omitempty"`
	Meta         string           `gorm:"column:meta;type:longtext;check:json_valid(meta)" json:"meta,omitempty"`
	MetaKeywords string           `gorm:"column:meta_keywords;type:text" json:"meta_keywords,omitempty"`

	// Relationship: blog belongs to one user
	User *User `gorm:"foreignKey:UID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}

// TableName overrides default pluralization
func (Blog) TableName() string {
	return "blogs"
}
