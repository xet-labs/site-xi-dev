package blog

import "time"

type BlogSitemap struct {
	Username  string    `json:"username"`
	Slug      string    `json:"slug"`
	UpdatedAt time.Time `json:"updated_at"`
}
