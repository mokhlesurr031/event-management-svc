package domain

type Pagination struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	TotalPages  int `json:"total_page"`
	CurrentPage int `json:"current_page"`
}
