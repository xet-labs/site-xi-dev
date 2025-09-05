package blog

import "time"

type SitemapBlogs struct {
    Username  string    `json:"username"`
    Slug      string    `json:"slug"`
    UpdatedAt time.Time `json:"updated_at"`
}
