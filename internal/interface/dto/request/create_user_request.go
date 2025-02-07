package request

type CreateUserRequest struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}
