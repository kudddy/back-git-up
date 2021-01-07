package handlers

import (
	"git-up-back/MessageTypes"
	"git-up-back/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CheckStatusJob(res http.ResponseWriter, req *http.Request) {
	// генерация данных для проверки
	//tokenstatus:= false
	//status := MessageTypes.CheckTokenResp{MessageName: "TOKENSTATUS", Status: tokenstatus, StatusCode: 200}

	// достаем токен
	// обязательно проверять авторизацию

	token := mux.Vars(req)["token"]

	//ok := realCheck(token)
	ok := true
	var status MessageTypes.CheckJobStatusResp

	status.MessageName = "JOBSTATUS"

	// достаем статус задач

	//result := models.GetJobStatusFromDb(token)


	result, err := models.GetMC().Get(token)
	if err != nil{
		status.Status = "JOB NOT START YET"
	}else{
		if !ok {
			status.Status = string(result.Value)

		} else {
			status.Status = string(result.Value)
		}
	}
	//fmt.Printf("Тут мы проверяем результат")
	//fmt.Printf("\n")
	//print(len(result.Status))
	//fmt.Printf("Закончили проверять результат")
	//if len(result.Status) == 0 {
	//	result.Status = "JOB NOT START YET"
	//}
	//if !ok {
	//	status.Status = result.Status
	//
	//} else {
	//	status.Status = result.Status
	//}
	status.Token = token

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Write(js)

}
