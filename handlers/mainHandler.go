package handlers

import (
	"back-git-up/MessageTypes"
	"back-git-up/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Hello(in chan MessageTypes.Profile, out chan MessageTypes.Profile) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		// Пример декодирования json
		decoder := json.NewDecoder(req.Body)
		var userinfo MessageTypes.UserToken
		err := decoder.Decode(&userinfo)
		if err != nil {
			panic(err)
		}
		fmt.Println(userinfo.Token)

		profile := MessageTypes.Profile{Name: "Alex", Hobbies: []string{"snowboarding", "programming"}}

		// Пример записи в базу данных
		models.AddUser()

		// пример записи в базу
		result := models.GetUser(1)

		fmt.Println(result.Email)

		//проверка работы роута с переменой

		vars := mux.Vars(req)

		fmt.Printf(vars["token"])

		//models.GetUser(5)

		// пример обмена сообщениями между потоками
		select {
		case in <- profile:
			fmt.Println("received message from hello", profile)
		case msg1 := <-out:
			profile = msg1
		default:
			fmt.Println("no message received from hello")
		}

		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}

		res.Header().Set("Content-Type", "application/json")
		res.Write(js)
	}
}
