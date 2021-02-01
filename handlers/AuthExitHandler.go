package handlers

import (
	"back-git-up/models"
	"back-git-up/utils"
	"fmt"
	"github.com/golang/glog"
	"net/http"
)

func ExitAuth(res http.ResponseWriter, req *http.Request) {

	print("что за метод:")

	print(req.Method)

	print("\n")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", utils.FrontHost)
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Headers", "Cache, Accept,Content-Type,Host,Accept")
	res.Header().Set("Access-Control-Request-Headers", "Cache, Accept,Content-Type,Host,Accept")

	if req.Method == "GET" {
		store := models.GetCookiesStore()

		session, _ := store.Get(req, utils.SessionName)

		models.GetMC().Delete(session.ID)

		glog.Info(fmt.Sprintf("client exit with sesion id: %s ", session.ID))

	}
}
