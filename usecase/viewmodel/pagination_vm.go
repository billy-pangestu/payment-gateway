package viewmodel

// PaginationVM ...
type PaginationVM struct {
	CurrentPage   int `json:"current_page"`
	LastPage      int `json:"last_page"`
	Count         int `json:"count"`
	RecordPerPage int `json:"record_per_page"`
}

// SimplePaginationVM ...
type SimplePaginationVM struct {
	PrevPage string `json:"prev_page"`
	NextPage string `json:"next_page"`
}
