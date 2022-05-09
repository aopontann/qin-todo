package main

type GetUserResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Avatar_url string `json:"avatar_url"`
}

type PutUserRequestBody struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
