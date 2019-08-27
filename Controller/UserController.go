package Controller

import (
	YoyoGo "github.com/maxzhang1985/yoyogo/framework"
	"net/http"
)

type UserController struct {

}


type RegiserResult struct {
	UserName string
	Password string
}

func (p *UserController) Register(w http.ResponseWriter,r *http.Request){
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	registerResult := RegiserResult{ UserName: username, Password:password }
	result := YoyoGo.ApiResult{ Success: true, Message: "ok", Data: registerResult }

	//jsons,_ := json.Marshal(result)
	//
	//_, _ = w.Write(jsons)
	YoyoGo.JSON(w,result)

}

func (p *UserController) GetInfo(w http.ResponseWriter,r *http.Request) {
	result := YoyoGo.ApiResult{ Success: true, Message: "ok"}
	YoyoGo.JSON(w,result)
}

