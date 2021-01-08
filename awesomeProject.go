package main

import (
	"back-git-up/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	//memcacheClient := memcache.New("127.0.0.1:11211")
	r := mux.NewRouter()

	// добавляем милдварю
	r.Use(handlers.JwtAuthentication)

	//in := make(chan MessageTypes.Profile)
	//
	//out := make(chan MessageTypes.Profile)
	//тут должен быть запущен контроллер фононов задач
	//go Job.Workers(in, out)
	// скорее всего нужен будет менеджер

	r.HandleFunc("/{token}/checktoken", handlers.CheckToken)

	r.HandleFunc("/{token}/start/addfriend", handlers.StarJobAdd)

	r.HandleFunc("/{token}/stop/addfriend", handlers.StopJobAdd)

	//r.HandleFunc("/status/{token}", handlers.Hello(in, out))

	r.HandleFunc("/{token}/jobstatusbytoken", handlers.CheckStatusJob)

	http.Handle("/", r)
	http.ListenAndServe(":9000", nil)
}

//func main() {
//	token := "982be9592eb2dd614d58e341d323d96ceee54e1c"
//	Job.AddFriendWorker(token)
//}
