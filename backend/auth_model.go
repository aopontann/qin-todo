package main

type UserRegisterReqb struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,alphanum,min=6"`
}

type UserLoginReqb struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,alphanum"`
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}