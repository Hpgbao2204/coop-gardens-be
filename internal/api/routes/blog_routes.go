package routes

import (
	"coop-gardens-be/internal/api/handlers"

	"github.com/labstack/echo/v4"
)

func BlogRoutes(g *echo.Group, blogHandler *handlers.BlogHandler) {
	// Blog endpoints
	g.POST("", blogHandler.CreateBlog)
	g.GET("", blogHandler.GetAllBlogs)
	g.GET("/:id", blogHandler.GetBlogByID)

	// Comment endpoints under a blog
	g.POST("/:blog_id/comments", blogHandler.CreateComment)
	g.GET("/:blog_id/comments", blogHandler.GetComments)

	// Review endpoints (for reviews on inventory)
	g.POST("/reviews", blogHandler.CreateReview)
	g.GET("/reviews/:inventory_id", blogHandler.GetReviewsByInventory)
}
