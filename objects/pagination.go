package objects

type Pagination struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	TotalPage int   `json:"total_page"`
	TotalData int64 `json:"total_data"`
}
