package handlers

import (
	"back-git-up/MessageTypes"
	"back-git-up/models"
	"back-git-up/utils"
	"encoding/json"
	"net/http"
)

func CheckAuth(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", utils.FrontHost)
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
			print("проверяем что приходит из кэша\n")
			print(result.Value)

			if err == nil {
				if result != nil {

					status.StatusAuth = true
					status.IsNewSession = true
					status.Token = string(result.Value)

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
