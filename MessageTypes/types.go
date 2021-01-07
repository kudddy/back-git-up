package MessageTypes

type Profile struct {
	Name    string
	Hobbies []string
}

type UserToken struct {
	MessageName string
	Token       string
}

type CheckTokenResp struct {
	MessageName string
	Status      bool
	Desc        string
	Token string
}

type StarGazers struct {
	Login             string `json:"login"`
	Id                int32  `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

//type StarGazersStruct struct {
//	data []StarGazers
//}

type StartJobResp struct {
	MessageName string
	Status      bool
}

type CheckJobStatusResp struct {
	MessageName string
	Status      string
	Token string
}
