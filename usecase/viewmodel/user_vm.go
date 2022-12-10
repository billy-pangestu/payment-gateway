package viewmodel

// UserVM ...
type UserVM struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	UniqueID  string  `json:"unique_id"` // A.K.A Nomor Rekening
	Amount    float64 `json:"amount"`
	Password  string  `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt string  `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`

	RoleID   string `json:"role_id,omitempty" bson:"role_id,omitempty"`
	RoleName string `json:"role_name,omitempty" bson:"role_name,omitempty"`
}
