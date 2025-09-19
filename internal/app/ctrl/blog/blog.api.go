// cntr/blog.api.go
package blog

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	model_store "xi/internal/app/model/store"
	"xi/pkg/lib"
	"xi/pkg/service/store"
)

type BlogApiCtrl struct {
	// dbCli  *gorm.DB
	// rdbCli *redis.Client
	mu   sync.RWMutex
	once sync.Once
}

// Singleton controller
var (
	BlogApi = &BlogApiCtrl{}

	ErrInvalidUserName = errors.New("invalid username")
	ErrInvalidUID      = errors.New("invalid UID")
	ErrInvalidSlug     = errors.New("invalid slug")
	ErrBlogNotFound    = errors.New("blog not found")
)

// GET /blog or /blog?Page=2&Limit=6
func (b *BlogApiCtrl) Index(c *gin.Context) {
	page := c.DefaultQuery("Page", "1")
	limit := c.DefaultQuery("Limit", "6")
	pageNum, err1 := strconv.Atoi(page)
	limitNum, err2 := strconv.Atoi(limit)

	if err1 != nil || err2 != nil || pageNum <= 0 || limitNum <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Page or Limit params"})
		return
	}

	// Try cache
	rdbKey := c.Request.URL.String()
	blogs := []model_store.Blog{}
	if err := store.Rdb.GetJson(rdbKey, &blogs); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"blogsExhausted": len(blogs) == 0,
			"blogs":          blogs,
		})
		return
	}

	// Try DB
	offset := (pageNum - 1) * limitNum
	if err := b.IndexCore(&blogs, offset, limitNum); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch blogs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"blogsExhausted": len(blogs) == 0,
		"blogs":          blogs,
	})

	// Async cache
	go func(data any) { store.Rdb.SetJson(rdbKey, data, 10*time.Minute) }(blogs)
}

func (b *BlogApiCtrl) IndexCore(blogs *[]model_store.Blog, offset, limit int) error {
	return store.Db.Cli.
		Preload("User").
		Where("status IN ?", []string{"published", "published_hidden"}).
		Order("updated_at DESC").
		Offset(offset).
		Limit(limit).
		Find(blogs).
		Error
}

// GET api/blog/uid/id
func (b *BlogApiCtrl) Show(c *gin.Context) {
	rawUID := c.Param("uid") // @username or UID
	rawID := c.Param("id")   // blog ID or slug

	// Try cache
	blog := model_store.Blog{}
	rdbKey := "/api/blog/" + rawUID + "/" + rawID
	if err := store.Rdb.GetJson(rdbKey, &blog); err == nil {
		c.JSON(http.StatusOK, blog)
		return
	}

	// Validate parameters
	if err := b.Validate(rawUID, rawID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fallback to DB
	if err := b.ShowCore(&blog, rawUID, rawID); err != nil {
		status := http.StatusNotFound
		if err == ErrInvalidUID {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)

	// Cache asynchronously
	go func(data model_store.Blog) { store.Rdb.SetJson(rdbKey, data, 10*time.Minute) }(blog)
}

// FetchBlog fetches a blog and stores it in the given pointer.
// It returns the Redis key and any error.
func (b *BlogApiCtrl) ShowCore(dest *model_store.Blog, rawUID, rawID string) error {
	var err error

	// Case 1: @username format
	if username, ok := strings.CutPrefix(rawUID, "@"); ok {
		// DB fallback
		err = store.Db.Cli.Preload("User").
			Joins("JOIN users ON users.id = blogs.uid").
			Where("users.username = ? AND (blogs.slug = ? OR blogs.id = ?)", username, rawID, rawID).
			First(dest).Error

		// Case 2: UID (numeric)
	} else if isNumeric(rawUID) {

		err = store.Db.Cli.Preload("User").
			Where("uid = ? AND (slug = ? OR id = ?)", rawUID, rawID, rawID).
			First(dest).Error

		// Invalid UID format
	} else {
		return ErrInvalidUID
	}

	if err != nil {
		return ErrBlogNotFound
	}

	return nil
}

// POST api/blog/uid/id
func (b *BlogApiCtrl) Post(c *gin.Context) {
	var blog model_store.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	blog.CreatedAt = ptrTime(time.Now())

	if err := store.Db.Cli.Create(&blog).Error; err != nil {
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

	if err := store.Db.Cli.First(&blog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	blog.UpdatedAt = ptrTime(time.Now())

	if err := store.Db.Cli.Save(&blog).Error; err != nil {
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

	if err := store.Db.Cli.Delete(&model_store.Blog{}, id).Error; err != nil {
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
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func (b *BlogApiCtrl) Validate(rawUID, rawID string) error {
	if strings.HasPrefix(rawUID, "@") {
		if !lib.Validate.Uname(rawUID) {
			return ErrInvalidUserName
		}
	} else if !lib.Validate.UID(rawUID) {
		return ErrInvalidUID
	}
	if !lib.Validate.Slug(rawID) {
		return ErrInvalidSlug
	}
	return nil
}
