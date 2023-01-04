package form

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
} //@name UserDTO

type UserLogin struct {
	Email string `json:"email" binding:"required,email"`
} //@name UserLoginDTO

type DeleteResponse struct {
	Message string `json:"message"`
} //@name DeleteResponse

type LoginResponse struct {
	Id    string `json:"id,omitempty"`
	Email string `json:"email" binding:"required,email"`
	Token string `json:"token,omitempty"`
} //@name LoginResponse
