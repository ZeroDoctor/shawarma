package model

type GithubUser struct {
	Login                   string         `json:"login"`
	ID                      int            `json:"id" db:"id"`
	NodeID                  string         `json:"node_id"`
	AvatarURL               string         `json:"avatar_url" db:"avatar_url"`
	GravatarID              string         `json:"gravatar_id"`
	URL                     string         `json:"url" db:"url"`
	HtmlURL                 string         `json:"html_url"`
	FollowersURL            string         `json:"followers_url"`
	FollowingURL            string         `json:"following_url"`
	GistsURL                string         `json:"gists_url"`
	StarredURL              string         `json:"starred_url"`
	SubscriptionsURL        string         `json:"subscriptions_url"`
	OrganizationsURL        string         `json:"organizations_url" db:"organizations_url"`
	ReposURL                string         `json:"repos_url" db:"repos_url"`
	EventsURL               string         `json:"events_url"`
	ReceivedEventsURL       string         `json:"received_events_url"`
	Type                    string         `json:"type" db:"type"`
	SiteAdmin               bool           `json:"site_admin"`
	Name                    string         `json:"name" db:"name"`
	Company                 string         `json:"company"`
	Blog                    string         `blog:"blog"`
	Location                string         `json:"location"`
	Email                   string         `json:"email"`
	Hireable                bool           `json:"hireable"`
	Bio                     string         `json:"bio"`
	TwitterUsername         string         `json:"twitter_username"`
	PublicRepos             int            `json:"public_repos"`
	PublicGists             int            `json:"public_gists"`
	Followers               int            `json:"followers"`
	Following               int            `json:"following"`
	CreatedAt               Time           `json:"created_at" db:"created_at"`
	UpdatedAt               Time           `json:"updated_at" db:"updated_at"`
	PrivateGists            int            `json:"private_gists"`
	TotalPrivateRepos       int            `json:"total_private_repos"`
	OwnedPrivateRepos       int            `json:"owned_private_repos"`
	DiskUsage               int            `json:"disk_usage"`
	Collaborators           int            `json:"collaborators"`
	TwoFactorAuthentication bool           `json:"two_factor_authentication"`
	Plan                    GithubUserPlan `json:"plan"`
}

type GithubUserPlan struct {
	Name          string `json:"name"`
	Space         int    `json:"space"`
	PrivateRepos  int    `json:"private_repos"`
	Collaborators int    `json:"collaborators"`
}

type GithubTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}
