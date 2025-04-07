package handlers

import (
    "net/http"
    "strconv"

    "coop-gardens-be/internal/models"
    "coop-gardens-be/internal/usecase"

    "github.com/labstack/echo/v4"
)

type BlogHandler struct {
    Usecase *usecase.BlogUsecase
}

func NewBlogHandler(usecase *usecase.BlogUsecase) *BlogHandler {
    return &BlogHandler{Usecase: usecase}
}

// CreateBlog tạo một blog mới
func (h *BlogHandler) CreateBlog(c echo.Context) error {
    var blog models.Blog
    if err := c.Bind(&blog); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
    }
    if err := h.Usecase.CreateBlog(&blog); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, blog)
}

// GetAllBlogs trả về danh sách blog
func (h *BlogHandler) GetAllBlogs(c echo.Context) error {
    blogs, err := h.Usecase.GetAllBlogs()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, blogs)
}

// GetBlogByID trả về blog theo id
func (h *BlogHandler) GetBlogByID(c echo.Context) error {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid blog ID"})
    }
    blog, err := h.Usecase.GetBlogByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Blog not found"})
    }
    return c.JSON(http.StatusOK, blog)
}

// CreateComment tạo comment cho blog
func (h *BlogHandler) CreateComment(c echo.Context) error {
    idParam := c.Param("blog_id")
    blogID, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid blog ID"})
    }
    var comment models.Comment
    if err := c.Bind(&comment); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
    }
    comment.BlogID = uint(blogID)
    if err := h.Usecase.CreateComment(&comment); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, comment)
}

// GetComments trả về danh sách comment theo blog
func (h *BlogHandler) GetComments(c echo.Context) error {
    idParam := c.Param("blog_id")
    blogID, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid blog ID"})
    }
    comments, err := h.Usecase.GetCommentsByBlogID(uint(blogID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, comments)
}

// CreateReview tạo review cho inventory
func (h *BlogHandler) CreateReview(c echo.Context) error {
    var review models.Review
    if err := c.Bind(&review); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid payload"})
    }
    if err := h.Usecase.CreateReview(&review); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusCreated, review)
}

// GetReviewsByInventory trả về danh sách review theo inventory
func (h *BlogHandler) GetReviewsByInventory(c echo.Context) error {
    idParam := c.Param("inventory_id")
    inventoryID, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid inventory ID"})
    }
    reviews, err := h.Usecase.GetReviewsByInventoryID(uint(inventoryID))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, reviews)
}