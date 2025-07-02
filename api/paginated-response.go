package api

type PaginatedResponse[T interface{}] struct {
	Elements []T `json:"elements"`
	Total    int `json:"total_elements"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func CreatePaginatedResponse[T interface{}](data []T, total int, params PaginatedRequest) PaginatedResponse[T] {
	return PaginatedResponse[T]{
		Elements: data,
		Total:    total,
		Page:     params.Page,
		PageSize: params.PageSize,
	}
}
