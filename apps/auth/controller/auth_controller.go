package controller

import (
	"net/http"

	"github.com/asishshaji/startup/apps/auth/usecase"
	"github.com/gorilla/mux"
)

// Authcontroller struct
type Authcontroller struct {
	usecase usecase.UseCase
}

// NewAuthController created authcontroller
func NewAuthController(usecase usecase.UseCase) *Authcontroller {
	return &Authcontroller{
		usecase: usecase,
	}
}

// Signin methods login the user
func (controller *Authcontroller) Signin(response http.ResponseWriter, request *http.Request) {

	token, err := controller.usecase.SignIn(request.Context(), mux.Vars(request)["username"], mux.Vars(request)["password"])

	if err != nil {
		http.Error(response, "Wrong password", http.StatusBadRequest)
		return
	}

	http.SetCookie(response, &http.Cookie{
		Name:  "token",
		Value: token,
	})
}

// Signup creates a new user
func (controller *Authcontroller) Signup(response http.ResponseWriter, request *http.Request) {

	// _ := mux.Vars(request)

	data := request.URL.Query()
	username, _ := data["username"]
	password, _ := data["password"]
	err := controller.usecase.SignUp(request.Context(), username[0], password[0])

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	response.Write([]byte("Created user"))
}
