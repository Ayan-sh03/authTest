package docs

type Register struct {
	Firstname  string `json:"first_name" validate:"required"`
	Lastname   string `json:"last_name" validate:"required"`
	Middlename string `json:"middle_name"`
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

type OTP struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type OTP_successResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
