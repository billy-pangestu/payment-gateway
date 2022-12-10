package viewmodel

// MerchantVM ...
type MerchantVM struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Amount    float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt string  `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
}
