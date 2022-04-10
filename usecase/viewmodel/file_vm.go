package viewmodel

// FileVM ....
type FileVM struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	URL        string `json:"url"`
	TempURL    string `json:"temp_url"`
	UserUpload string `json:"user_upload"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}
