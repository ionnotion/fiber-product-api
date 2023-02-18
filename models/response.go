package models

type Response struct {
	Message string `json:"message"`
	LoggedUser string `json:"loggedUser,omitempty"`
	Access_token string `json:"access_token,omitempty"`
}