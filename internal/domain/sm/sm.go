package domain

type Secret struct {
	Host 			string `json:"host"`
	Username 	string `json:"username"`
	Password 	string `json:"password"`
	JWTSign 	string `json:"jwtSign"`
	IsServ		bool	 `json:"isServ"`
	Database 	string `json:"database"`
}