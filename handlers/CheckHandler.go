package handlers

import (
	"back-git-up/MessageTypes"
	"back-git-up/models"
	"back-git-up/utils"
	c "back-git-up/utils"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/glog"
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

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", c.FrontHost)
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Headers", "Cache, Accept,content-type,Host,Accept")
	res.Header().Set("Access-Control-Request-Headers", "Cache, Accept,content-type,Host,Accept")

	// достаем токен
	if req.Method == "GET" {
		sessionn, _ := models.GetCookiesStore().Get(req, utils.SessionName)

		if sessionn.IsNew {
			glog.Info(fmt.Sprintf("сессия новая"))
			http.Redirect(res, req, utils.FrontHost, http.StatusFound)
			return
		}
		token := mux.Vars(req)["token"]

		ok := realCheck(token)

		//ok := true
		var status MessageTypes.CheckTokenResp
		fmt.Printf("\n")
		print("статус авторизации:")
		print(ok)
		fmt.Printf("\n")

		status.MessageName = "TOKENSTATUS"

		session, _ := models.GetCookiesStore().Get(req, "session-name")
		print("йо\n")
		print(session.ID)
		// проверяем старая ли сессия

		// проверяем аутентифицировался ли пользователь ранее путем проверки в кэше
		//TODO добавить обработчик ошибок
		results, _ := models.GetMC().Get(session.ID)
		if results == nil {
			// в кэшэ ничего нет, пользователь заходит первый раз
			print("в кэшэ ничего нет, пользователь заходит первый раз\n")
			if ok {
				//если с авторизацией ок то пишем в кэш
				print("если с авторизацией ок то пишем в кэш:")
				print(session.ID)
				print("\n")
				status.Status = true
				models.GetMC().Set(&memcache.Item{Key: session.ID, Value: []byte(token + "_ok")})
			} else {
				print("если не ок, то отправляем отрицательный статус")
				//если не ок, то отправляем отрицательный статус
				status.Status = false
			}
		} else {
			status.Status = true
		}
		//if session.IsNew{
		//	print("сессия новая\n")
		//	//status.Status = false
		//
		//	//если сессия новая  но авторизация прошла успешно, то пишем в кэш и возвращаем тру
		//	if ok {
		//		print("токен валидный, пишем в кэш\n")
		//		print("смотрим что пишем в кэш:")
		//		print(session.ID)
		//		print("\n")
		//		status.Status = true
		//		models.GetMC().Set(&memcache.Item{Key: session.ID, Value: []byte(token + "_ok")})
		//		//если нет, то возвращаем false
		//	}else{
		//		print("токен невалидный, не пишем в кэш\n")
		//		status.Status = true
		//	}
		//
		//
		//	// если сессия старая
		//} else{
		//	print("сесия старая\n")
		//	if ok {
		//		// возвращаем тру и ничего не делаем
		//		print("c токеном все ок, пропускаем дальше\n")
		//		status.Status = true
		//		//models.GetMC().Set(&memcache.Item{Key: session.ID, Value: []byte(token + "_notok")})
		//
		//	} else {
		//		status.Status = false
		//		// тут достаем id сессии для записи в kv
		//		print("токен невалидный\n")
		//		//models.GetMC().Set(&memcache.Item{Key: session.ID, Value: []byte(token + "_ok")})
		//
		//	}
		//}

		status.Token = token

		js, err := json.Marshal(status)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		res.Write(js)
	}

}
