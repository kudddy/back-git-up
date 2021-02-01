package handlers

import (
	"back-git-up/MessageTypes"
	"back-git-up/models"
	"back-git-up/utils"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
)

func CheckStatusJob(res http.ResponseWriter, req *http.Request) {

	if req.Method == "GET" {
		//session, _ := models.GetCookiesStore().Get(req, utils.SessionName)
		//
		////if session.IsNew{
		////	glog.Info(fmt.Sprintf("сессия новая"))
		////	http.Redirect(res, req, utils.FrontHost, http.StatusFound)
		////	return
		////}

		token := mux.Vars(req)["token"]

		glog.Info(fmt.Sprintf("Chech status job with token: %s ", token))

		memTokenStat := token + "_status"
		ok := true
		var status MessageTypes.CheckJobStatusResp

		status.MessageName = "JOBSTATUS"

		status.CountFriendAdd = models.GetCountAddFriend(token)

		result, err := models.GetMC().Get(memTokenStat)
		if err != nil {
			status.Status = "JOB NOT START YET"
		} else {
			if !ok {
				status.Status = string(result.Value)

			} else {
				status.Status = string(result.Value)
			}
		}
		status.Token = token

		js, err := json.Marshal(status)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}

		//res.Header().Set("Content-Type", "application/json")
		//res.Header().Set("Access-Control-Allow-Origin", utils.FrontHost)
		glog.Info(fmt.Sprintf("Отправляем токен: %s ", token))
		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", utils.FrontHost)
		res.Header().Set("Access-Control-Allow-Credentials", "true")
		res.Header().Set("Access-Control-Allow-Headers", "Cache, Accept,Content-Type,Host,Accept")
		res.Header().Set("Access-Control-Request-Headers", "Cache, Accept,Content-Type,Host,Accept")
		res.Write(js)

	}

}
