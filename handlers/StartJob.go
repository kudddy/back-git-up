package handlers

import (
	"back-git-up/Job"
	"back-git-up/MessageTypes"
	"back-git-up/models"
	"encoding/json"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/golang/glog"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func StarJobAdd(res http.ResponseWriter, req *http.Request) {
	// нужно обернуть для получения данных от основной горутины
	// проверяем валидный ли токен

	vars := mux.Vars(req)
	// добавил префикс для разделения ключей
	token := vars["token"]
	// добавил префикс для разделения ключей, префикс _status - для статуса задачи

	glog.Info(fmt.Sprintf("Start Job with token: %s ", token))

	memTokenStat := token + "_status"

	//TODO проверить нужно ли проверить валидность токена. Наверное нет
	//tokenStatus:=realCheck(token)
	tokenStatus := true

	var workerStatus MessageTypes.CheckTokenResp

	workerStatus.MessageName = "STARTJOBADD"

	//генерация уникального id задачи
	id := guuid.New()
	fmt.Println(id.String())
	// TODO требуется запускать воркеры с статусом RUNNING  при старте приложения

	//session, _ := models.GetCookiesStore().Get(req, sessionName)

	glog.Info(fmt.Sprintf("Start Job with token: %s ", token))

	if tokenStatus {
		//запуск воркера
		fmt.Println("запуск воркера")
		// нужна проверка не запущен ли воркер уже/узнать статус и только потом запускать
		// пишем в memcached
		models.GetMC().Set(&memcache.Item{Key: memTokenStat, Value: []byte("RUNNING")})
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
