package models

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authorization struct {
	Token string `json:"token"`
}
