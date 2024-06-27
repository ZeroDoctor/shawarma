package model

type GithubUser struct {
	Login                   string     `json:"login,omitempty" db:"name"`
	ID                      int        `json:"id,omitempty" db:"id"`
	NodeID                  string     `json:"node_id,omitempty"`
	AvatarURL               string     `json:"avatar_url,omitempty" db:"avatar_url"`
	GravatarID              string     `json:"gravatar_id,omitempty" db:"gravatar_id"`
	URL                     string     `json:"url,omitempty" db:"url"`
	HtmlURL                 string     `json:"html_url,omitempty"`
	FollowersURL            string     `json:"followers_url,omitempty"`
	FollowingURL            string     `json:"following_url,omitempty"`
	GistsURL                string     `json:"gists_url,omitempty"`
	StarredURL              string     `json:"starred_url,omitempty"`
	SubscriptionsURL        string     `json:"subscriptions_url,omitempty"`
	OrganizationsURL        string     `json:"organizations_url,omitempty" db:"organizations_url"`
	ReposURL                string     `json:"repos_url,omitempty" db:"repos_url"`
	EventsURL               string     `json:"events_url,omitempty"`
	ReceivedEventsURL       string     `json:"received_events_url,omitempty"`
	Type                    string     `json:"type,omitempty" db:"type"`
	SiteAdmin               bool       `json:"site_admin,omitempty"`
	Name                    string     `json:"name,omitempty"`
	Company                 string     `json:"company,omitempty" db:"company"`
	Blog                    string     `blog:"blog,omitempty"`
	Location                string     `json:"location,omitempty"`
	Email                   string     `json:"email,omitempty"`
	Hireable                bool       `json:"hireable,omitempty"`
	Bio                     string     `json:"bio,omitempty"`
	TwitterUsername         string     `json:"twitter_username,omitempty"`
	PublicRepos             int        `json:"public_repos,omitempty"`
	PublicGists             int        `json:"public_gists,omitempty"`
	Followers               int        `json:"followers,omitempty"`
	Following               int        `json:"following,omitempty"`
	CreatedAt               string     `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt               string     `json:"updated_at,omitempty" db:"updated_at"`
	PrivateGists            int        `json:"private_gists,omitempty"`
	TotalPrivateRepos       int        `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos       int        `json:"owned_private_repos,omitempty"`
	DiskUsage               int        `json:"disk_usage,omitempty"`
	Collaborators           int        `json:"collaborators,omitempty"`
	TwoFactorAuthentication bool       `json:"two_factor_authentication,omitempty"`
	Plan                    GithubPlan `json:"plan,omitempty"`
	Token                   string     `db:"token"`

	Orgs  []GithubOrg
	Repos []GithubRepo
}

type GithubOwner struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty" db:"id"`
	NodeID            string `json:"node_id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty" db:"avatar_url"`
	GravatarID        string `json:"gravatar_id,omitempty" db:"gravatar_id"`
	URL               string `json:"url,omitempty" db:"url"`
	HtmlURL           string `json:"html_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty" db:"organizations_url"`
	ReposURL          string `json:"repos_url,omitempty" db:"repos_url"`
	EventsURL         string `json:"events_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	Type              string `json:"type,omitempty" db:"type"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
}

type GithubPlan struct {
	Name          string `json:"name,omitempty"`
	Space         int    `json:"space,omitempty"`
	PrivateRepos  int    `json:"private_repos,omitempty"`
	Collaborators int    `json:"collaborators,omitempty"`
	FilledSeats   int    `json:"filled_seats,omitempty"`
	Seats         int    `json:"seats,omitempty"`
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

type GithubUsersOrgs struct {
	UserID    string `json:"github_user_id,omitempty" db:"github_user_id"`
	OrgID     string `json:"github_org_id,omitempty" db:"github_org_id"`
	CreatedAt Time   `json:"created_at,omitempty" db:"created_at"`
}

type GithubOrg struct {
	Login                                string     `json:"login,omitempty" db:"name"`
	ID                                   int        `json:"id,omitempty" db:"id"`
	NodeID                               string     `json:"node_id,omitempty"`
	URL                                  string     `json:"url,omitempty" db:"url"`
	ReposURL                             string     `json:"repos_url,omitempty" db:"repos_url"`
	EventsURL                            string     `json:"events_url,omitempty"`
	HooksURL                             string     `json:"hooks_url,omitempty" db:"hooks_url"`
	IssuesURL                            string     `json:"issues_url,omitempty" db:"issues_url"`
	MembersURL                           string     `json:"members_url,omitempty" db:"members_url"`
	PublicMembersURL                     string     `json:"public_members_url,omitempty" db:"public_members_url"`
	AvatarURL                            string     `json:"avatar_url,omitempty" db:"avatar_url"`
	Description                          string     `json:"description,omitempty" db:"description"`
	Name                                 string     `json:"name,omitempty"`
	Company                              string     `json:"company,omitempty" db:"company"`
	Blog                                 string     `json:"blog,omitempty"`
	Location                             string     `json:"location,omitempty"`
	Email                                string     `json:"email,omitempty"`
	TwitterUsername                      string     `json:"twitter_username,omitempty"`
	IsVerified                           bool       `json:"is_verified,omitempty"`
	HasOrganizationProjects              bool       `json:"has_organization_projects,omitempty"`
	HasRepositoryProjects                bool       `json:"has_repository_projects,omitempty"`
	PublicRepos                          int        `json:"public_repos,omitempty"`
	PublicGists                          int        `json:"public_gists,omitempty"`
	Followers                            int        `json:"followers,omitempty"`
	Following                            int        `json:"followings,omitempty"`
	HtmlURL                              string     `json:"html_url,omitempty"`
	CreatedAt                            string     `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt                            string     `json:"updated_at,omitempty" db:"updated_at"`
	ArchivedAt                           string     `json:"archived_at,omitempty" db:"archived_at"`
	Type                                 string     `json:"type,omitempty" db:"type"`
	TotalPrivateRepos                    int        `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos                    int        `json:"owned_private_repos,omitempty"`
	PrivateGists                         int        `json:"private_gists,omitempty"`
	DiskUsage                            int        `json:"disk_usage,omitempty"`
	Collaborators                        int        `json:"collaborators,omitempty"`
	BillingEmail                         string     `json:"billing_email,omitempty"`
	Plan                                 GithubPlan `json:"plan,omitempty"`
	DefaultRepositoryPermission          string     `json:"default_repository_permission,omitempty"`
	MembersCanCreateRepositories         bool       `json:"members_can_create_repositories,omitempty"`
	MembersAllowedRepositoryCreationType string     `json:"members_allowed_repository_creation_type,omitempty"`
	MembersCanCreatePublicRepositories   bool       `json:"members_can_create_public_repositories,omitempty"`
	MembersCanCreatePrivateRepositories  bool       `json:"members_can_create_private_repositories,omitempty"`
	MembersCanCreateInternalRepositories bool       `json:"members_can_create_internal_repositories,omitempty"`
}

type GithubRepo struct {
	ID                       int           `json:"id,omitempty" db:"id"`
	NodeID                   string        `json:"node_id,omitempty"`
	Name                     string        `json:"name,omitempty" db:"name"`
	FullName                 string        `json:"full_name,omitempty" db:"full_name"`
	Private                  bool          `json:"private,omitempty"`
	Owner                    GithubOwner   `json:"owner,omitempty"`
	OwnerID                  int           `db:"owner_id"`
	HtmlURL                  string        `json:"html_url,omitempty"`
	Description              string        `json:"description,omitempty" db:"description"`
	Fork                     bool          `json:"fork,omitempty"`
	URL                      string        `json:"url,omitempty" db:"url"`
	ForksURL                 string        `json:"forks_url,omitempty"`
	KeysURL                  string        `json:"keys_url,omitempty"`
	CollaboratorsURL         string        `json:"collaborators_url,omitempty" db:"collaborators_url"`
	TeamsURL                 string        `json:"teams_url,omitempty"`
	HooksURL                 string        `json:"hooks_url,omitempty" db:"hooks_url"`
	IssueEventsURL           string        `json:"issue_events_url,omitempty" db:"issue_events_url"`
	EventsURL                string        `json:"events_url,omitempty"`
	AssigneesURL             string        `json:"assignees_url,omitempty"`
	BranchesURL              string        `json:"branches_url,omitempty" db:"branches_url"`
	TagsURL                  string        `json:"tags_url,omitempty" db:"tags_url"`
	BlobsURL                 string        `json:"blobs_url,omitempty"`
	GitTagsURL               string        `json:"git_tags_url,omitempty"`
	GitRefsURL               string        `json:"git_refs_url,omitempty"`
	TreesURL                 string        `json:"trees_url,omitempty"`
	StatusesURL              string        `json:"statuses_url,omitempty" db:"statuses_url"`
	LanguagesURL             string        `json:"languages_url,omitempty"`
	StargazersURL            string        `json:"stargazers_url,omitempty"`
	ContributorsURL          string        `json:"contributors_url,omitempty"`
	SubscribersURL           string        `json:"subscribers_url,omitempty"`
	SubscriptionURL          string        `json:"subscription_url,omitempty"`
	CommitsURL               string        `json:"commits_url,omitempty" db:"commits_url"`
	GitCommitsURL            string        `json:"git_commits_url,omitempty"`
	CommentsURL              string        `json:"comments_url,omitempty"`
	IssueCommentURL          string        `json:"issue_comment_url,omitempty"`
	ContentsURL              string        `json:"contents_url,omitempty"`
	CompareURL               string        `json:"compare_url,omitempty"`
	MergesURL                string        `json:"merges_url,omitempty" db:"merges_url"`
	ArchiveURL               string        `json:"archive_url,omitempty"`
	DownloadsURL             string        `json:"downloads_url,omitempty"`
	IssuesURL                string        `json:"issues_url,omitempty" db:"issues_url"`
	PullsURL                 string        `json:"pulls_url,omitempty" db:"pulls_url"`
	MilestonesURL            string        `json:"milestones_url,omitempty"`
	NotificationsURL         string        `json:"notifications_url,omitempty"`
	LabelsURL                string        `json:"labels_url,omitempty"`
	ReleasesURL              string        `json:"releases_url,omitempty"`
	DeploymentsURL           string        `json:"deployments_url,omitempty"`
	CreatedAt                string        `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt                string        `json:"updated_at,omitempty" db:"updated_at"`
	PushedAt                 string        `json:"pushed_at,omitempty" db:"pushed_at"`
	GitURL                   string        `json:"git_url,omitempty"`
	SshURL                   string        `json:"ssh_url,omitempty"`
	CloneURL                 string        `json:"clone_url,omitempty"`
	SvnURL                   string        `json:"svn_url,omitempty"`
	HomePage                 string        `json:"homepage,omitempty"`
	Size                     int           `json:"size,omitempty"`
	StargazersCount          int           `json:"stargazers_count,omitempty"`
	WatchersCount            int           `json:"watchers_count,omitempty"`
	Language                 string        `json:"language,omitempty"`
	HasIssues                bool          `json:"has_issues,omitempty" db:"has_issues"`
	HasProjects              bool          `json:"has_projects,omitempty"`
	HasDownloads             bool          `json:"has_downloads,omitempty"`
	HasWiki                  bool          `json:"has_wiki,omitempty"`
	HasPages                 bool          `json:"has_pages,omitempty"`
	HasDiscussions           bool          `json:"has_discussions,omitempty"`
	ForksCount               int           `json:"forks_counts,omitempty"`
	MirrorURL                string        `json:"mirror_url,omitempty"`
	Archived                 bool          `json:"archived,omitempty" db:"archived"`
	Disable                  bool          `json:"disable,omitempty"`
	OpenIssuesCount          int           `json:"open_issues_count,omitempty" db:"open_issues_count"`
	License                  GithubLicense `json:"license,omitempty"`
	AllowForking             bool          `json:"allow_forking,omitempty"`
	IsTemplate               bool          `json:"is_template,omitempty"`
	WebCommitSignOffRequired bool          `json:"web_commit_signoff_required,omitempty"`
	Topics                   []string      `json:"topics,omitempty"`
	Visibility               string        `json:"visibility,omitempty" db:"visibility"`
	Forks                    int           `json:"forks,omitempty"`
	OpenIssues               int           `json:"open_issues,omitempty"`
	Watchers                 int           `json:"watchers,omitempty"`
	DefaultBranch            string        `json:"default_branch,omitempty"`
	Organization             GithubOrg     `json:"organization,omitempty"`
}

type GithubLicense struct {
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	SPDXID string `json:"spdx_id,omitempty"`
	URL    string `json:"url,omitempty"`
	NodeID string `json:"node_id,omitempty"`
}

type GithubBranch struct {
	Name          string         `json:"name,omitempty" db:"name"`
	Commit        GithubCommit   `json:"commit,omitempty"`
	URL           string         `json:"url,omitempty" db:"url"`
	HtmlURL       string         `json:"html_url,omitempty"`
	CommentsURL   string         `json:"comments_url,omitempty"`
	Author        GithubOwner    `json:"author,omitempty"`
	AuthorID      int            `db:"author_id"`
	Committer     GithubOwner    `json:"committer,omitempty"`
	CommitterID   int            `db:"committer_id"`
	Parents       []GithubCommit `json:"parents,omitempty"`
	Protected     bool           `json:"protected,omitempty"`
	ProtectionURL string         `json:"protection_url,omitempty"`
}

type GithubCommit struct {
	SHA       string             `json:"sha,omitempty" db:"sha"`
	NodeID    string             `json:"node_id,omitempty"`
	Author    GithubCommitAuthor `json:"author,omitempty"`
	Committer GithubCommitAuthor `json:"committer,omitempty"`
	Message   string             `json:"message,omitempty" db:"message"`
	Tree      GithubTree         `json:"tree,omitempty"`
	URL       string             `json:"url,omitempty" db:"url"`
	Parents   []GithubCommit     `json:"parents,omitempty"`
}

type GithubCommitAuthor struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Date  string `json:"date,omitempty"`
}

type GithubTree struct {
	SHA string `json:"sha,omitempty"`
	URL string `json:"url,omitempty"`
}
