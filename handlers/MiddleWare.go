package handlers

import (
	u "git-up-back/utils"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/user/new", "/api/user/login"} //Список эндпоинтов, для которых не требуется авторизация
		requestPath := r.URL.Path                               //текущий путь запроса
		fmt.Printf("мы тут")
		fmt.Printf(requestPath)

		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		//tokenHeader := r.Header.Get("Authorization") //Получение токена

		vars := mux.Vars(r)

		token := vars["token"]

		fmt.Printf(token)

		if token == "" { //Токен отсутствует, возвращаем  403 http-код Unauthorized
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//tokenPart := splitted[1] //Получаем вторую часть токена
		//tk := &models.Token{}

		//token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		//	return []byte(os.Getenv("token_password")), nil
		//})

		//if err != nil { //Неправильный токен, как правило, возвращает 403 http-код
		//	response = u.Message(false, "Malformed authentication token")
		//	w.WriteHeader(http.StatusForbidden)
		//	w.Header().Add("Content-Type", "application/json")
		//	u.Respond(w, response)
		//	return
		//}

		if token == "123" { //токен недействителен, возможно, не подписан на этом сервере
			response = u.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		//Всё прошло хорошо, продолжаем выполнение запроса
		//fmt.Sprintf("User %", tk.Username) //Полезно для мониторинга
		ctx := context.WithValue(r.Context(), "user", 123)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //передать управление следующему обработчику!
	})
}
