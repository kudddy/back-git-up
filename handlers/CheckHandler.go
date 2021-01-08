package handlers

import (
	"back-git-up/MessageTypes"
	c "back-git-up/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func realCheck(token string) bool {

	// проверка api на доступ к необходимым методам
	//follow user Unfollow a user

	client := &http.Client{}
	var status bool
	for _, method := range c.AllowMethodForReq {
		// TODO пока используем Артема снигерева
		req, _ := http.NewRequest(method, c.UrlAddFriend+"ArtemSnegirev", nil)
		req.Header.Set("Accept", c.GitAccept)
		req.Header.Set("Authorization", "token"+" "+token)

		res, err := client.Do(req)

		if err != nil {
			status = false
		} else if res.StatusCode == 204 {
			status = true
		} else {
			status = false
		}

	}
	return status

}

func CheckToken(res http.ResponseWriter, req *http.Request) {
	// генерация данных для проверки
	//tokenstatus:= false
	//status := MessageTypes.CheckTokenResp{MessageName: "TOKENSTATUS", Status: tokenstatus, StatusCode: 200}

	// достаем токен

	token := mux.Vars(req)["token"]

	ok := realCheck(token)
	fmt.Printf("\n")
	fmt.Printf(token)
	fmt.Printf("\n")
	//ok := true
	var status MessageTypes.CheckTokenResp
	fmt.Printf("\n")
	print(ok)
	fmt.Printf("\n")

	status.MessageName = "TOKENSTATUS"

	if !ok {
		status.Status = false

	} else {
		status.Status = true
	}
	status.Token = token

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(js)

}
