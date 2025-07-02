package api

type PaginatedRequest struct {
	Page     int `form:"page" json:"page" binding:"gte=1"`
	PageSize int `form:"page_size" json:"page_size" binding:"gte=1,lte=1000"`
}
