package request

type UserUpdateRequest struct {
	Password        string `json:"password" validate:"required,min=5,max=100"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=5,max=100"`
}

type UserCreateRequest struct {
}
