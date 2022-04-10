package request

//UserRequest ...
type UserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	UniqueID  string `json:"unique_id"`
	Password  string `json:"password"`
	RoleID    string `json:"role_id"`
}

//UserLoginRequest ...
type UserLoginRequest struct {
	UniqueID string `json:"unique_id" validate:"required"`
	Password string `json:"password"`
}
