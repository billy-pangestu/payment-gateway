package request

//MerchantRequest ...
type MerchantRequest struct {
	Name string `json:"name" validate:"required"`
}
