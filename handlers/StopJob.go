package handlers

import (
	"back-git-up/MessageTypes"
	"back-git-up/models"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/mux"
	"net/http"
)

func StopJobAdd(res http.ResponseWriter, req *http.Request) {
	// нужно обернуть для получения данных от основной горутины
	// проверяем валидный ли токен
	vars := mux.Vars(req)

	token := vars["token"]

	memTokenStat := token + "_status"

	fmt.Println(token)

	//tokenStatus:=realCheck(token)
	tokenStatus := true

	var workerStatus MessageTypes.CheckTokenResp

	workerStatus.MessageName = "STOPJOBADD"

	if tokenStatus {
		//запуск воркера
		fmt.Println("стоп выполения воркера")
		// нужна проверка не запущен ли воркер уже/узнать статус и только потом запускать
		workerStatus.Status = true
		// запись команды на остановка работы воркера
		//go Job.StartWorker(token)
		models.GetMC().Set(&memcache.Item{Key: memTokenStat, Value: []byte("FINISH")})
		models.InsertCancelAction(token)

	} else {
		fmt.Println("Отправка сообщения о невалидности токена")
		workerStatus.Status = false
	}

	js, err := json.Marshal(workerStatus)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(js)

}
