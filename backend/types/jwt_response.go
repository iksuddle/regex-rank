package types

type JWTResponse struct {
	Name  string
	Value string
}

func NewJWTResponse(jwt string) JWTResponse {
	return JWTResponse{
		Name:  "jwt",
		Value: jwt,
	}
}
