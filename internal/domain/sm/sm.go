package domain

type Secret struct {
	Host 			string `json:"host"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	JWTSign 	string `json:"jwt_sign"`
	Database 	string `json:"database"`
}