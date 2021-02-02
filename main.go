package main

import (
	"back-git-up/handlers"
	"back-git-up/models"
	"back-git-up/utils"
	"flag"
	"fmt"
	"github.com/dghubble/gologin"
	"github.com/dghubble/gologin/github"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"strconv"
	"time"

	//"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	githubOAuth2 "golang.org/x/oauth2/github"
	"net/http"
)

type Config struct {
	GithubClientID     string
	GithubClientSecret string
}

const (
	sessionName       = "example-github-app"
	sessionSecret     = "example cookie signing secret"
	sessionUserKey    = "githubID"
	sessionUsername   = "githubUsername"
	sessionGitHubUrls = "githubUrl"
)

// addCookie will apply a new cookie to the response of a http request
// with the key/value specified.
func addCookie(w http.ResponseWriter, name, value string, ttl time.Duration) {
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

// issueSession issues a cookie session after successful Github login
func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		githubUser, err := github.UserFromContext(ctx)
		//token, _ := oauth2Login.TokenFromContext(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 2. Implement a success handler to issue some form of session
		session, _ := models.GetCookiesStore().New(req, sessionName)
		userId := strconv.Itoa(int(*githubUser.ID))
		session.Values[sessionUserKey] = *githubUser.ID
		session.Values[sessionUsername] = *githubUser.Login
		session.Values[sessionGitHubUrls] = *githubUser.AvatarURL

		session.Save(req, w)

		//cookie, err := req.Cookie(sessionName)
		//print(fmt.Sprintf("cookie  %s ", w.Header().Get("Set-Cookie")))
		//
		//
		//print(fmt.Sprintf("cookie new  %s ", w.Header().Values("Set-Cookie")))

		header := http.Header{}
		header.Add("Cookie", w.Header().Get("Set-Cookie"))
		request := http.Request{Header: header}
		cookie, err := request.Cookie("example-github-app")

		if err != nil {
			print(fmt.Sprintf("Cant find cookie :/\r\n"))
		} else {
			print(fmt.Sprintf("we find cookie  %s ", cookie.Value))
		}

		http.Redirect(w, req, utils.FrontHost+"/jobinfo/"+userId+"&"+cookie.Value, http.StatusFound)
		//http.Redirect(w, req, "/profile", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

// profileHandler shows a personal profile or a login button (unauthenticated).
func profileHandler(w http.ResponseWriter, req *http.Request) {
	//session, err := sessionStore.Get(req, sessionName)
	session, _ := models.GetCookiesStore().Get(req, sessionName)
	//if err != nil {
	//
	//}
	//print("мы тут\n")
	//print("лалаленд")
	//print(fmt.Sprintf("cookie  %s ", req.Header.Get("Cookie")))

	// Read cookie TODO убрать
	//cookie, err := req.Cookie("example-github-app")
	//print(fmt.Sprintf("cookie  %s ", req.Header.Get("Cookie")))
	//
	//if err != nil {
	//	fmt.Printf("Cant find cookie :/\r\n")
	//} else{
	//	print(fmt.Sprintf("we find cookie  %s ", cookie.Value))
	//}

	//addCookie(w, "TestCookieName", "TestValue", 30*time.Minute)

	if session.IsNew {
		http.Redirect(w, req, utils.FrontHost, http.StatusFound)
		return
	}

	//
	//githubUser, err := github.UserFromContext(ctx)
	//print("смотрим инфу по юзеру")
	//print(githubUser.Login)
	// authenticated profile
	fmt.Fprintf(w, `<p>You are loggedfff in %s! %s</p><form action="/logout" method="post"><input type="submit" value="Logout"></form>`, session.Values[sessionUsername], session.Values[sessionGitHubUrls])
}

// logoutHandler destroys the session on POSTs and redirects to home.
//func logoutHandler(w http.ResponseWriter, req *http.Request) {
//	print("смотрим какой метод:")
//	print(req.Method)
//	if req.Method == "POST" {
//		store := models.GetCookiesStore()
//
//
//
//		session, _ := store.Get(req, "example-github-app")
//
//		print(session.ID)
//		models.GetMC().Delete(session.ID)
//		//models.GetCookiesStore().Delete(w, sessionName)
//
//	}
//	http.Redirect(w, req, utils.FrontHost, http.StatusFound)
//
//}
// New returns a new ServeMux with app routes.
func New(config *Config) *mux.Router {

	mux := mux.NewRouter()

	mux.Use(handlers.JwtAuthentication)

	mux.HandleFunc("/", profileHandler)

	mux.HandleFunc("/{token}/checktoken", handlers.CheckToken)

	mux.HandleFunc("/{token}/start/addfriend", handlers.StarJobAdd)

	mux.HandleFunc("/{token}/stop/addfriend", handlers.StopJobAdd)

	mux.HandleFunc("/{token}/jobstatusbytoken", handlers.CheckStatusJob)

	//тест
	mux.HandleFunc("/checkout", handlers.CheckAuth)
	mux.HandleFunc("/exit", handlers.ExitAuth)

	// 1. Register LoginHandler and CallbackHandler
	oauth2Config := &oauth2.Config{
		ClientID:     config.GithubClientID,
		ClientSecret: config.GithubClientSecret,
		RedirectURL:  "http://localhost:8080/github/callback",
		Endpoint:     githubOAuth2.Endpoint,
		Scopes:       []string{"user:follow"},
	}
	stateConfig := gologin.DebugOnlyCookieConfig
	mux.Handle("/github/login", github.StateHandler(stateConfig, github.LoginHandler(oauth2Config, nil)))
	mux.Handle("/github/callback", github.StateHandler(stateConfig, github.CallbackHandler(oauth2Config, issueSession(), nil)))
	return mux
}

func main() {
	const address = "0.0.0.0:8080"

	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")

	glog.Info("----INIT APP----")

	config := &Config{
		GithubClientID:     "f97d3137e37cf2bc803f",
		GithubClientSecret: "5ce8ed966ffaccf65e9ba069f5f09b4063613cb7",
	}

	glog.Info("----START APP----")

	//http.Handle("/", r)
	http.ListenAndServe(address, New(config))

}

//func main() {
//	token := "982be9592eb2dd614d58e341d323d96ceee54e1c"
//	Job.AddFriendWorker(token)
//}
