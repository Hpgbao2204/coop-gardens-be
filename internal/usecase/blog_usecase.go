package usecase

import (
	"errors"

	"coop-gardens-be/internal/models"
	"coop-gardens-be/internal/repository"
)

type BlogUsecase struct {
	repo *repository.BlogRepository
}

func NewBlogUsecase(repo *repository.BlogRepository) *BlogUsecase {
	return &BlogUsecase{repo}
}

func (u *BlogUsecase) CreateBlog(blog *models.Blog) error {
	if blog.Title == "" || blog.Content == "" || blog.AuthorID == "" {
		return errors.New("title, content and author_id are required")
	}
	return u.repo.CreateBlog(blog)
}

func (u *BlogUsecase) GetAllBlogs() ([]models.Blog, error) {
	return u.repo.GetAllBlogs()
}

func (u *BlogUsecase) GetBlogByID(id uint) (*models.Blog, error) {
	return u.repo.GetBlogByID(id)
}

func (u *BlogUsecase) CreateComment(comment *models.Comment) error {
	if comment.Content == "" {
		return errors.New("comment content is required")
	}
	return u.repo.CreateComment(comment)
}

func (u *BlogUsecase) GetCommentsByBlogID(blogID uint) ([]models.Comment, error) {
	return u.repo.GetCommentsByBlogID(blogID)
}

func (u *BlogUsecase) CreateReview(review *models.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return u.repo.CreateReview(review)
}

func (u *BlogUsecase) GetReviewsByInventoryID(inventoryID uint) ([]models.Review, error) {
	return u.repo.GetReviewsByInventoryID(inventoryID)
}
