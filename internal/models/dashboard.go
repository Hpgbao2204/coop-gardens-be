package models

// DashboardSummary chứa các thông tin tổng hợp cho dashboard.
type DashboardSummary struct {
	TotalCrops        int     `json:"total_crops"`
	TotalSeasons      int     `json:"total_seasons"`
	TotalHarvestYield float64 `json:"total_harvest_yield"`
	TotalUsersAdmin   int     `json:"total_users_admin"`
	TotalUsersFarmer  int     `json:"total_users_farmer"`
	TotalUsers        int     `json:"total_users"`
	AvgProductRating  float64 `json:"avg_product_rating"`
	TotalOrders       int     `json:"total_orders"`
}
