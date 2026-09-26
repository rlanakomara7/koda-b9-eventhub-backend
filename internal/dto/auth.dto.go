package dto

//Register
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

//Login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

//response succes
type AuthResponse struct {
	UserID  uint   `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
