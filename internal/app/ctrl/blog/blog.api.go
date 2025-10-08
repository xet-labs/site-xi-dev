// cntr/blog.api.go
package blog

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	model_store "xi/internal/app/model/store"
	"xi/pkg/app"
	appErr "xi/pkg/app/err"
	"xi/pkg/lib"
	"xi/pkg/service/store"
)

type BlogApiCtrl struct{}

// Singleton controller
var (
	BlogApi = &BlogApiCtrl{}
)

// GET /blog or /blog?Page=2&Limit=6
func (b *BlogApiCtrl) Index(c *gin.Context) {
	// Parse and validate query params
	pageNum, limitNum, ok := parsePageLimit(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Page or Limit parameters"})
		return
	}

	rdbKey := c.Request.URL.String()
	var blogs []model_store.Blog

	// Try cache
	if err := store.Rdb.GetJson(rdbKey, &blogs); err == nil {
		c.JSON(http.StatusOK, map[string]any{
			"blogsExhausted": len(blogs) == 0,
			"blogs":          blogs,
		})
		return
	}

	// Fetch from DB
	offset := (pageNum - 1) * limitNum
	if err := b.IndexCore(&blogs, offset, limitNum); err != nil {
		app.Err.Handle(c, err, true)
		return
	}

	// Respond
	c.JSON(http.StatusOK, map[string]any{
		"blogsExhausted": len(blogs) == 0,
		"blogs":          blogs,
	})

	// Cache asynchronously
	go store.Rdb.SetJson(rdbKey, blogs, 10*time.Minute)
}

func parsePageLimit(c *gin.Context) (page, limit int, ok bool) {
	pageStr := c.DefaultQuery("Page", "1")
	limitStr := c.DefaultQuery("Limit", "6")

	p, err1 := strconv.Atoi(pageStr)
	l, err2 := strconv.Atoi(limitStr)
	if err1 != nil || err2 != nil || p <= 0 || l <= 0 {
		return 0, 0, false
	}
	return p, l, true
}

func (b *BlogApiCtrl) IndexCore(blogs *[]model_store.Blog, offset, limit int) error {
	db := store.Db.Cli()
	if db.Error != nil {
		return db.Error
	}

	return db.Preload("User").
		Where("status IN ?", []string{"published", "published_hidden"}).
		Order("updated_at DESC").
		Offset(offset).
		Limit(limit).
		Find(blogs).Error
}

// GET api/blog/uid/id
func (b *BlogApiCtrl) Show(c *gin.Context) {
	rawUID := c.Param("uid")
	rawID := c.Param("id")

	// Validate first
	if err := b.Validate(rawUID, rawID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}

	rdbKey := "/api/blog/" + rawUID + "/" + rawID
	var blog model_store.Blog

	// Try cache
	if err := store.Rdb.GetJson(rdbKey, &blog); err == nil {
		c.JSON(http.StatusOK, blog)
		return
	}

	// Fallback to DB
	if err := b.ShowCore(&blog, rawUID, rawID); err != nil {
		// All invalid/missing resources return 404 for pages
		c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}

	c.JSON(http.StatusOK, blog)

	// Cache asynchronously
	go store.Rdb.SetJson(rdbKey, blog, 10*time.Minute)
}

// FetchBlog fetches a blog and stores it in the given pointer.
// It returns the Redis key and any error.
func (b *BlogApiCtrl) ShowCore(dest *model_store.Blog, rawUID, rawID string) error {
	db := store.Db.Cli()
	if db.Error != nil {
		return db.Error
	}

	if username, ok := strings.CutPrefix(rawUID, "@"); ok {
		return db.Preload("User").Joins("JOIN users ON users.id = blogs.uid").Where("users.username = ? AND (blogs.slug = ? OR blogs.id = ?)", username, rawID, rawID).First(dest).Error
	}

	return db.Preload("User").Where("uid = ? AND (slug = ? OR id = ?)", rawUID, rawID, rawID).First(dest).Error
}

// POST api/blog/uid/id
func (b *BlogApiCtrl) Post(c *gin.Context) {
	var blog model_store.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	blog.CreatedAt = ptrTime(time.Now())

	if err := store.Db.Cli().Create(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create blog"})
		return
	}

	c.JSON(http.StatusCreated, blog)

	// Invalidate blog list cache
	store.Rdb.Del("blogs:all")
}

// PUT api/blog/uid/id
func (b *BlogApiCtrl) Put(c *gin.Context) {
	id := c.Param("id")
	var blog model_store.Blog

	if err := store.Db.Cli().First(&blog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	blog.UpdatedAt = ptrTime(time.Now())

	if err := store.Db.Cli().Save(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update blog"})
		return
	}

	c.JSON(http.StatusOK, blog)

	// Invalidate caches
	store.Rdb.Del("blogs:all", "blogs:id:"+id)
}

// DELETE api/blog/uid/id
func (b *BlogApiCtrl) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := store.Db.Cli().Delete(&model_store.Blog{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete blog"})
		return
	}

	// Invalidate caches
	store.Rdb.Del("blogs:all", "blogs:id:"+id)

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})
}

// --HELPERS--
// Utility: return pointer to time
func ptrTime(t time.Time) *time.Time {
	return &t
}

func (b *BlogApiCtrl) Validate(rawUID, rawID string) error {
	if strings.HasPrefix(rawUID, "@") {
		if !lib.Validate.Uname(rawUID) {
			return appErr.InvalidUserName.Err
		}
	} else if !lib.Validate.UID(rawUID) {
		return appErr.InvalidUID.Err
	}
	if !lib.Validate.Slug(rawID) {
		return appErr.InvalidSlug.Err
	}
	return nil
}
