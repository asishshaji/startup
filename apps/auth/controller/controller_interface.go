package controller

import "net/http"

type AuthController interface {
	Signin(response http.ResponseWriter, request *http.Request)
	Signup(response http.ResponseWriter, request *http.Request)
}
