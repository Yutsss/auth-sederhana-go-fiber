package dtos

type (
	UserRegisterRequest struct {
		Username string `json:"username" form:"username" validate:"required,min=1,max=255"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8,max=255"`
	}

	UserRegisterResponse struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8,max=255"`
	}

	UserLoginResponse struct {
		AccessToken string `json:"access_token"`
	}

	UserGetByIdRequest struct {
		Id int64 `json:"id" form:"id" validate:"required"`
	}

	UserGetByIdResponse struct {
		Id       int64  `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)
