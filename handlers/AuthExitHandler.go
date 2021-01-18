package handlers

import (
	"back-git-up/models"
	"net/http"
)

func ExitAuth(res http.ResponseWriter, req *http.Request) {

	print("что за метод:")

	print(req.Method)

	print("\n")

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Headers", "Cache, Accept,Content-Type,Host,Accept")
	res.Header().Set("Access-Control-Request-Headers", "Cache, Accept,Content-Type,Host,Accept")

	if req.Method == "GET" {
		store := models.GetCookiesStore()

		session, _ := store.Get(req, "session-name")

		//sessionId := session.ID
		//print(sessionId+"\n")

		//var status MessageTypes.CheckCookie
		// новая сессия, просим клиента авторизоваться
		models.GetMC().Delete(session.ID)

		//status.MessageName = "COOCKIESTATUS"

		//js, err := json.Marshal(status)
		//if err != nil {
		//	http.Error(res, err.Error(), http.StatusInternalServerError)
		//
		//}
		//res.Write(js)
	}
}
