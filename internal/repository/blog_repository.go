package repository

import (
    "coop-gardens-be/internal/models"

    "gorm.io/gorm"
)

type BlogRepository struct {
    DB *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
    return &BlogRepository{DB: db}
}

func (r *BlogRepository) CreateBlog(blog *models.Blog) error {
    return r.DB.Create(blog).Error
}

func (r *BlogRepository) GetAllBlogs() ([]models.Blog, error) {
    var blogs []models.Blog
    err := r.DB.Preload("Author").Find(&blogs).Error
    return blogs, err
}

func (r *BlogRepository) GetBlogByID(id uint) (*models.Blog, error) {
    var blog models.Blog
    err := r.DB.Preload("Author").First(&blog, id).Error
    if err != nil {
        return nil, err
    }
    return &blog, nil
}

func (r *BlogRepository) CreateComment(comment *models.Comment) error {
    return r.DB.Create(comment).Error
}

func (r *BlogRepository) GetCommentsByBlogID(blogID uint) ([]models.Comment, error) {
    var comments []models.Comment
    err := r.DB.Preload("Author").Where("blog_id = ?", blogID).Find(&comments).Error
    return comments, err
}

func (r *BlogRepository) CreateReview(review *models.Review) error {
    return r.DB.Create(review).Error
}

func (r *BlogRepository) GetReviewsByInventoryID(inventoryID uint) ([]models.Review, error) {
    var reviews []models.Review
    err := r.DB.Preload("User").Where("inventory_id = ?", inventoryID).Find(&reviews).Error
    return reviews, err
}