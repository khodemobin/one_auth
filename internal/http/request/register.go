package request

type RegisterRequest struct {
	Phone string `json:"phone" validate:"required,min=11,max=11,number"`
}

type RegisterVerifyRequest struct {
	Phone string `json:"phone" validate:"required,min=11,max=11,number"`
	Code  string `json:"code" validate:"required,min=6,max=6,number"`
}
