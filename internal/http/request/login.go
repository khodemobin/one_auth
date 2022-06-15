package request

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required,min=11,max=11,number"`
	Password string `json:"password" validate:"required,min=5"`
}
