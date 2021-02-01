package utils

import "os"

// константы для Github
const UrlGetStargazers = "https://api.github.com/repos/tiangolo/fastapi/stargazers"
const UrlAddFriend = "https://api.github.com/user/following/"
const GitAccept = "application/vnd.github.v3+json"

const (
	SessionName       = "example-github-app"
	SessionSecret     = "example cookie signing secret"
	SessionUserKey    = "githubID"
	SessionUsername   = "githubUsername"
	SessionGitHubUrls = "githubUrl"
)

var FrontHost = os.Getenv("frontend_adress")

var AllowMethodForReq = [2]string{"PUT", "DELETE"}

var MapMainLangToRepo = map[string]string{
	"python": "https://api.github.com/repos/tiangolo/fastapi/stargazers",
	"r":      "2138",
	"gri":    "1908",
	"adg":    "8484",
}
