package model

type GithubUser struct {
	Login                   string         `json:"login,omitempty"`
	ID                      int            `json:"id,omitempty" db:"id"`
	NodeID                  string         `json:"node_id,omitempty"`
	AvatarURL               string         `json:"avatar_url,omitempty" db:"avatar_url"`
	GravatarID              string         `json:"gravatar_id,omitempty" db:"gravatar_id"`
	URL                     string         `json:"url,omitempty" db:"url"`
	HtmlURL                 string         `json:"html_url,omitempty"`
	FollowersURL            string         `json:"followers_url,omitempty"`
	FollowingURL            string         `json:"following_url,omitempty"`
	GistsURL                string         `json:"gists_url,omitempty"`
	StarredURL              string         `json:"starred_url,omitempty"`
	SubscriptionsURL        string         `json:"subscriptions_url,omitempty"`
	OrganizationsURL        string         `json:"organizations_url,omitempty" db:"organizations_url"`
	ReposURL                string         `json:"repos_url,omitempty" db:"repos_url"`
	EventsURL               string         `json:"events_url,omitempty"`
	ReceivedEventsURL       string         `json:"received_events_url,omitempty"`
	Type                    string         `json:"type,omitempty" db:"type"`
	SiteAdmin               bool           `json:"site_admin,omitempty"`
	Name                    string         `json:"name,omitempty" db:"name"`
	Company                 string         `json:"company,omitempty" db:"company"`
	Blog                    string         `blog:"blog,omitempty"`
	Location                string         `json:"location,omitempty"`
	Email                   string         `json:"email,omitempty"`
	Hireable                bool           `json:"hireable,omitempty"`
	Bio                     string         `json:"bio,omitempty"`
	TwitterUsername         string         `json:"twitter_username,omitempty"`
	PublicRepos             int            `json:"public_repos,omitempty"`
	PublicGists             int            `json:"public_gists,omitempty"`
	Followers               int            `json:"followers,omitempty"`
	Following               int            `json:"following,omitempty"`
	CreatedAt               string         `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt               string         `json:"updated_at,omitempty" db:"updated_at"`
	PrivateGists            int            `json:"private_gists,omitempty"`
	TotalPrivateRepos       int            `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos       int            `json:"owned_private_repos,omitempty"`
	DiskUsage               int            `json:"disk_usage,omitempty"`
	Collaborators           int            `json:"collaborators,omitempty"`
	TwoFactorAuthentication bool           `json:"two_factor_authentication,omitempty"`
	Plan                    GithubUserPlan `json:"plan,omitempty"`
}

type GithubUserPlan struct {
	Name          string `json:"name,omitempty"`
	Space         int    `json:"space,omitempty"`
	PrivateRepos  int    `json:"private_repos,omitempty"`
	Collaborators int    `json:"collaborators,omitempty"`
}

type GithubTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	Scope       string `json:"scope,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}

type GithubUserOrg struct {
	Login            string `json:"login,omitempty"`
	ID               int    `json:"id,omitempty"`
	NodeID           string `json:"node_id,omitempty"`
	URL              string `json:"url,omitempty"`
	ReposURL         string `json:"repos_url,omitempty"`
	EventsURL        string `json:"events_url,omitempty"`
	HooksURL         string `json:"hooks_url,omitempty"`
	IssuesURL        string `json:"issues_url,omitempty"`
	MembersURL       string `json:"members_url,omitempty"`
	PublicMembersURL string `json:"public_members_url,omitempty"`
	AvatarURL        string `json:"avatar_url,omitempty"`
	Description      string `json:"description,omitempty"`
}

type GithubOrg struct {
	Login                   string `json:"login,omitempty"`
	ID                      int    `json:"id,omitempty" db:"id"`
	NodeID                  string `json:"node_id,omitempty"`
	URL                     string `json:"url,omitempty" db:"url"`
	ReposURL                string `json:"repos_url,omitempty" db:"repos_url"`
	EventsURL               string `json:"events_url,omitempty"`
	HooksURL                string `json:"hooks_url,omitempty" db:"hooks_url"`
	IssuesURL               string `json:"issues_url,omitempty" db:"issues_url"`
	MembersURL              string `json:"members_url,omitempty" db:"members_url"`
	PublicMembersURL        string `json:"public_members_url,omitempty" db:"public_members_url"`
	AvatarURL               string `json:"avatar_url,omitempty" db:"avatar_url"`
	Description             string `json:"description,omitempty" db:"description"`
	Name                    string `json:"name,omitempty" db:"name"`
	Company                 string `json:"company,omitempty" db:"company"`
	Blog                    string `json:"blog,omitempty"`
	Location                string `json:"location,omitempty"`
	Email                   string `json:"email,omitempty"`
	TwitterUsername         string `json:"twitter_username,omitempty"`
	IsVerified              bool   `json:"is_verified,omitempty"`
	HasOrganizationProjects bool   `json:"has_organization_projects,omitempty"`
	HasRepositoryProjects   bool   `json:"has_repository_projects,omitempty"`
	PublicRepos             int    `json:"public_repos,omitempty"`
	PublicGists             int    `json:"public_gists,omitempty"`
	Followers               int    `json:"followers,omitempty"`
	Following               int    `json:"followings,omitempty"`
	HtmlURL                 string `json:"html_url,omitempty"`
	CreatedAt               string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt               string `json:"updated_at,omitempty" db:"updated_at"`
	ArchivedAt              string `json:"archived_at,omitempty" db:"archived_at"`
	Type                    string `json:"type,omitempty" db:"type"`
}

type GithubRepos struct {
	ID                       int        `json:"id,omitempty" db:"id"`
	NodeID                   string     `json:"node_id,omitempty"`
	Name                     string     `json:"name,omitempty" db:"name"`
	FullName                 string     `json:"full_name,omitempty" db:"full_name"`
	Private                  bool       `json:"private,omitempty"`
	Owner                    GithubUser `json:"owner,omitempty"`
	HtmlURL                  string     `json:"html_url,omitempty"`
	Description              string     `json:"description,omitempty" db:"description"`
	Fork                     bool       `json:"fork,omitempty"`
	URL                      string     `json:"url,omitempty" db:"url"`
	ForksURL                 string     `json:"forks_url,omitempty"`
	KeysURL                  string     `json:"keys_url,omitempty"`
	CollaboratorsURL         string     `json:"collaborators_url,omitempty" db:"collaborators_url"`
	TeamsURL                 string     `json:"teams_url,omitempty"`
	HooksURL                 string     `json:"hooks_url,omitempty" db:"hooks_url"`
	IssueEventsURL           string     `json:"issue_events_url,omitempty" db:"issue_events_url"`
	EventsURL                string     `json:"events_url,omitempty"`
	AssigneesURL             string     `json:"assignees_url,omitempty"`
	BranchesURL              string     `json:"branches_url,omitempty" db:"branches_url"`
	TagsURL                  string     `json:"tags_url,omitempty" db:"tags_url"`
	BlobsURL                 string     `json:"blobs_url,omitempty"`
	GitTagsURL               string     `json:"git_tags_url,omitempty"`
	GitRefsURL               string     `json:"git_refs_url,omitempty"`
	TreesURL                 string     `json:"trees_url,omitempty"`
	StatusesURL              string     `json:"statuses_url,omitempty" db:"statuses_url"`
	LanguagesURL             string     `json:"languages_url,omitempty"`
	StargazersURL            string     `json:"stargazers_url,omitempty"`
	ContributorsURL          string     `json:"contributors_url,omitempty"`
	SubscribersURL           string     `json:"subscribers_url,omitempty"`
	SubscriptionURL          string     `json:"subscription_url,omitempty"`
	CommitsURL               string     `json:"commits_url,omitempty" db:"commits_url"`
	GitCommitsURL            string     `json:"git_commits_url,omitempty"`
	CommentsURL              string     `json:"comments_url,omitempty"`
	IssueCommentURL          string     `json:"issue_comment_url,omitempty"`
	ContentsURL              string     `json:"contents_url,omitempty"`
	CompareURL               string     `json:"compare_url,omitempty"`
	MergesURL                string     `json:"merges_url,omitempty" db:"merges_url"`
	ArchiveURL               string     `json:"archive_url,omitempty"`
	DownloadsURL             string     `json:"downloads_url,omitempty"`
	IssuesURL                string     `json:"issues_url,omitempty" db:"issues_url"`
	PullsURL                 string     `json:"pulls_url,omitempty" db:"pulls_url"`
	MilestonesURL            string     `json:"milestones_url,omitempty"`
	NotificationsURL         string     `json:"notifications_url,omitempty"`
	LabelsURL                string     `json:"labels_url,omitempty"`
	ReleasesURL              string     `json:"releases_url,omitempty"`
	DeploymentsURL           string     `json:"deployments_url,omitempty"`
	CreatedAt                string     `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt                string     `json:"updated_at,omitempty" db:"updated_at"`
	PushedAt                 string     `json:"pushed_at,omitempty" db:"pushed_at"`
	GitURL                   string     `json:"git_url,omitempty"`
	SshURL                   string     `json:"ssh_url,omitempty"`
	CloneURL                 string     `json:"clone_url,omitempty"`
	SvnURL                   string     `json:"svn_url,omitempty"`
	HomePage                 string     `json:"homepage,omitempty"`
	Size                     int        `json:"size,omitempty"`
	StargazersCount          int        `json:"stargazers_count,omitempty"`
	WatchersCount            int        `json:"watchers_count,omitempty"`
	Language                 string     `json:"language,omitempty"`
	HasIssues                bool       `json:"has_issues,omitempty"`
	HasProjects              bool       `json:"has_projects,omitempty"`
	HasDownloads             bool       `json:"has_downloads,omitempty"`
	HasWiki                  bool       `json:"has_wiki,omitempty"`
	HasPages                 bool       `json:"has_pages,omitempty"`
	HasDiscussions           bool       `json:"has_discussions,omitempty"`
	ForksCount               int        `json:"forks_counts,omitempty"`
	MirrorURL                string     `json:"mirror_url,omitempty"`
	Archived                 bool       `json:"archived,omitempty" db:"archived"`
	Disable                  bool       `json:"disable,omitempty"`
	OpenIssuesCount          int        `json:"open_issues_count,omitempty"`
	License                  string     `json:"license,omitempty"`
	AllowForking             bool       `json:"allow_forking,omitempty"`
	IsTemplate               bool       `json:"is_template,omitempty"`
	WebCommitSignOffRequired bool       `json:"web_commit_signoff_required,omitempty"`
	Topics                   []string   `json:"topics,omitempty"`
	Visibility               string     `json:"visibility,omitempty" db:"visibility"`
	Forks                    int        `json:"forks,omitempty"`
	OpenIssues               int        `json:"open_issues,omitempty"`
	Watchers                 int        `json:"watchers,omitempty"`
	DefaultBranch            string     `json:"default_branch,omitempty"`
}
