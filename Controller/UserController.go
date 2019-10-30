package Controller

import (
	"net/http"
)

type UserController struct {
}

type RegiserResult struct {
	UserName string
	Password string
}

func (p *UserController) Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	registerResult := RegiserResult{UserName: username, Password: password}
	result := ApiResult{Success: true, Message: "ok", Data: registerResult}

	//jsons,_ := json.Marshal(result)
	//
	//_, _ = w.Write(jsons)
	JSON(w, result)

}

func (p *UserController) GetInfo(w http.ResponseWriter, r *http.Request) {
	result := ApiResult{Success: true, Message: "ok"}
	JSON(w, result)
}
