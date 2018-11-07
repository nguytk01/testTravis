package businesslogic

type userLoginRequest struct {
	UserName string
  Email string
  SessionKey string
}

type LoginResult struct{
  LoginResult string
}