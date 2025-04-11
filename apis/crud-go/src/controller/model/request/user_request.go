package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:""` //containsany=!@#$...
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Age      int8   `json:"age" binding:"required,min=0,max=130"`
}
