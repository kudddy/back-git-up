package handlers

import (
	"back-git-up/MessageTypes"
	"back-git-up/models"
	"encoding/json"
	"net/http"
	"strings"
)

func CheckAuth(res http.ResponseWriter, req *http.Request) {

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

		var status MessageTypes.CheckCookie
		// новая сессия, просим клиента авторизоваться
		if session.IsNew {
			print("\nсессия новая")
			// запоминаем id сессии и прочую приблуду
			session.Values["foo"] = "bar"
			session.Values[42] = 43
			//session.Options.Domain ="127.0.0.1"
			//session.Options.Path = "/checkout"

			session.Save(req, res)
			status.StatusAuth = false
			status.IsNewSession = true
			print("пизда")
		} else {
			print("а теперь точно пизда")
			sessionId := session.ID
			print("узнаем сессию в момент когда сессия старая:")
			print(sessionId)
			print("\n")
			result, err := models.GetMC().Get(sessionId)

			if err == nil {
				if result != nil {

					tokenStatus := string(result.Value)
					tokenStat := strings.Split(tokenStatus, "_")
					checkTokenStatus := tokenStat[1]
					print("\n")
					print("статус токена:")
					print(checkTokenStatus)

					if checkTokenStatus == "ok" {
						status.StatusAuth = true
						status.IsNewSession = false
						status.Token = tokenStat[0]

					} else {

						status.StatusAuth = false
						status.IsNewSession = false
					}
				} else {
					print("пустое хранилище\n")

					status.StatusAuth = false
					status.IsNewSession = false
				}

			} else {
				print("обрабатываем ошибку")
				status.StatusAuth = false
				status.IsNewSession = false
			}

		}

		status.MessageName = "COOCKIESTATUS"

		js, err := json.Marshal(status)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)

		}
		res.Write(js)
	}
}
