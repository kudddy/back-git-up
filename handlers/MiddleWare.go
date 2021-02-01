package handlers

import (
	"back-git-up/models"
	"back-git-up/utils"
	"fmt"
	"github.com/golang/glog"
	"net/http"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		glog.Info(fmt.Sprintf("смотрим какой запрос %s", r.Method))
		if r.Method == "GET" || r.Method == "POST" {
			notAuth := []string{"/github/login", "/github/callback"}

			requestPath := r.URL.Path

			//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
			for _, value := range notAuth {

				if value == requestPath {
					next.ServeHTTP(w, r)
					return
				}
			}

			cookie, err := r.Cookie(utils.SessionName)
			print(fmt.Sprintf("cookie  %s ", r.Header.Get("Cookie")))
			//cookie, err := req.Cookie("example-github-app")

			//addCookie(w, sessionName, cookie.Value, 30*time.Minute)

			if err != nil {
				fmt.Printf("Cant find cookie :/\r\n")
			} else {
				print(fmt.Sprintf("we find cookie  %s ", cookie.Value))
			}

			glog.Info(fmt.Sprintf("We inter in middleware"))
			session, _ := models.GetCookiesStore().Get(r, utils.SessionName)
			if session.IsNew {
				glog.Info(fmt.Sprintf("its new session"))
				http.Redirect(w, r, utils.FrontHost, http.StatusFound)
				return

			} else {
				glog.Info(fmt.Sprintf("its old session"))
			}

			next.ServeHTTP(w, r) //передать управление следующему обработчику!
		}

	})
}
