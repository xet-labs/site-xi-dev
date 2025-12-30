# SEO & Sitemaps

Built-in SEO tools and a **Hook-based Sitemap Generator**.

## Hook-Based Sitemap
The sitemap is built dynamically by querying registered modules via the **Hooks System**. It eliminates the need for manual sitemap updates.

### Usage
To add URLs to the sitemap (`/sitemap.xml`), implement the `CoreSitemap` interface in your controller:

```go
// Logic: Return a slice of sitemap entries
func (c *MyCtrl) SitemapCore(ctx *gin.Context) (any, error) {
    return []model_config.MetaSitemap{
        {
            Loc:        "/my-page",
            ChangeFreq: "weekly",
            Priority:   "0.8",
        },
    }, nil
}
```

The generic `SitemapLib` uses generic hooks to aggregate these results automatically.

### SEO Metadata
Page titles and descriptions are managed via the `Meta` model and can be injected into templates dynamically.
