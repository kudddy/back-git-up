package handlers

//var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
//
//	account := &models.Account{}
//	err := json.NewDecoder(r.Body).Decode(account) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
//	if err != nil {
//		u.Respond(w, u.Message(false, "Invalid request"))
//		return
//	}
//
//	resp := account.Create() //Создать аккаунт
//	u.Respond(w, resp)
//}
