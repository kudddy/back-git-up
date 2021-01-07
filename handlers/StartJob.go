package handlers

import (
	"git-up-back/Job"
	"git-up-back/MessageTypes"
	"git-up-back/models"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func StarJobAdd(res http.ResponseWriter, req *http.Request) {
	// нужно обернуть для получения данных от основной горутины
	// проверяем валидный ли токен

	vars := mux.Vars(req)

	token := vars["token"]

	fmt.Println(token)

	//tokenStatus:=realCheck(token)
	tokenStatus := true

	var workerStatus MessageTypes.CheckTokenResp

	workerStatus.MessageName = "STARTJOBADD"

	//генерация уникального id задачи
	id := guuid.New()
	fmt.Println(id.String())
	// TODO требуется запускать воркеры с статусом RUNNING  при старте приложения



	if tokenStatus {
		//запуск воркера
		fmt.Println("запуск воркера")
		// нужна проверка не запущен ли воркер уже/узнать статус и только потом запускать
		// пишем в memcached
		models.GetMC().Set(&memcache.Item{Key: token, Value: []byte("RUNNING")})
		jobStatus := models.GetJobStatusFromDb(token)
		if jobStatus.Status == "RUNNING" {
			workerStatus.Status = false
			workerStatus.Desc = "the task is already running, complete the last"
			fmt.Printf("Задача запущена, для запуска требуется прервать предыдущую")
		} else if jobStatus.Status == "START" {
			workerStatus.Status = false
			workerStatus.Desc = "the task is starting, complete the last"
			fmt.Printf("Задача запускается, для запуска требуется прервать предыдущую")
		} else {
			workerStatus.Status = true
			workerStatus.Desc = "OK, start working"
			go Job.StartWorker(token)
		}

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
