package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	q "felixwie.com/savannah/queue"
	"github.com/gorilla/mux"
)

type GithubHook struct {
	Action string `json:"action"`
	Rule   struct {
		ID                                       int       `json:"id"`
		RepositoryID                             int       `json:"repository_id"`
		Name                                     string    `json:"name"`
		CreatedAt                                time.Time `json:"created_at"`
		UpdatedAt                                time.Time `json:"updated_at"`
		PullRequestReviewsEnforcementLevel       string    `json:"pull_request_reviews_enforcement_level"`
		RequiredApprovingReviewCount             int       `json:"required_approving_review_count"`
		DismissStaleReviewsOnPush                bool      `json:"dismiss_stale_reviews_on_push"`
		RequireCodeOwnerReview                   bool      `json:"require_code_owner_review"`
		AuthorizedDismissalActorsOnly            bool      `json:"authorized_dismissal_actors_only"`
		IgnoreApprovalsFromContributors          bool      `json:"ignore_approvals_from_contributors"`
		RequiredStatusChecks                     []string  `json:"required_status_checks"`
		RequiredStatusChecksEnforcementLevel     string    `json:"required_status_checks_enforcement_level"`
		StrictRequiredStatusChecksPolicy         bool      `json:"strict_required_status_checks_policy"`
		SignatureRequirementEnforcementLevel     string    `json:"signature_requirement_enforcement_level"`
		LinearHistoryRequirementEnforcementLevel string    `json:"linear_history_requirement_enforcement_level"`
		AdminEnforced                            bool      `json:"admin_enforced"`
		AllowForcePushesEnforcementLevel         string    `json:"allow_force_pushes_enforcement_level"`
		AllowDeletionsEnforcementLevel           string    `json:"allow_deletions_enforcement_level"`
		MergeQueueEnforcementLevel               string    `json:"merge_queue_enforcement_level"`
		RequiredDeploymentsEnforcementLevel      string    `json:"required_deployments_enforcement_level"`
		RequiredConversationResolutionLevel      string    `json:"required_conversation_resolution_level"`
		AuthorizedActorsOnly                     bool      `json:"authorized_actors_only"`
		AuthorizedActorNames                     []string  `json:"authorized_actor_names"`
	} `json:"rule"`
	Changes struct {
		AuthorizedActorsOnly struct {
			From bool `json:"from"`
		} `json:"authorized_actors_only"`
		AuthorizedActorNames struct {
			From []interface{} `json:"from"`
		} `json:"authorized_actor_names"`
	} `json:"changes"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		HTMLURL          string      `json:"html_url"`
		Description      string      `json:"description"`
		Fork             bool        `json:"fork"`
		URL              string      `json:"url"`
		ForksURL         string      `json:"forks_url"`
		KeysURL          string      `json:"keys_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		TeamsURL         string      `json:"teams_url"`
		HooksURL         string      `json:"hooks_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		EventsURL        string      `json:"events_url"`
		AssigneesURL     string      `json:"assignees_url"`
		BranchesURL      string      `json:"branches_url"`
		TagsURL          string      `json:"tags_url"`
		BlobsURL         string      `json:"blobs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		TreesURL         string      `json:"trees_url"`
		StatusesURL      string      `json:"statuses_url"`
		LanguagesURL     string      `json:"languages_url"`
		StargazersURL    string      `json:"stargazers_url"`
		ContributorsURL  string      `json:"contributors_url"`
		SubscribersURL   string      `json:"subscribers_url"`
		SubscriptionURL  string      `json:"subscription_url"`
		CommitsURL       string      `json:"commits_url"`
		GitCommitsURL    string      `json:"git_commits_url"`
		CommentsURL      string      `json:"comments_url"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		ContentsURL      string      `json:"contents_url"`
		CompareURL       string      `json:"compare_url"`
		MergesURL        string      `json:"merges_url"`
		ArchiveURL       string      `json:"archive_url"`
		DownloadsURL     string      `json:"downloads_url"`
		IssuesURL        string      `json:"issues_url"`
		PullsURL         string      `json:"pulls_url"`
		MilestonesURL    string      `json:"milestones_url"`
		NotificationsURL string      `json:"notifications_url"`
		LabelsURL        string      `json:"labels_url"`
		ReleasesURL      string      `json:"releases_url"`
		DeploymentsURL   string      `json:"deployments_url"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		PushedAt         time.Time   `json:"pushed_at"`
		GitURL           string      `json:"git_url"`
		SSHURL           string      `json:"ssh_url"`
		CloneURL         string      `json:"clone_url"`
		SvnURL           string      `json:"svn_url"`
		Homepage         string      `json:"homepage"`
		Size             int         `json:"size"`
		StargazersCount  int         `json:"stargazers_count"`
		WatchersCount    int         `json:"watchers_count"`
		Language         string      `json:"language"`
		HasIssues        bool        `json:"has_issues"`
		HasProjects      bool        `json:"has_projects"`
		HasDownloads     bool        `json:"has_downloads"`
		HasWiki          bool        `json:"has_wiki"`
		HasPages         bool        `json:"has_pages"`
		ForksCount       int         `json:"forks_count"`
		MirrorURL        interface{} `json:"mirror_url"`
		Archived         bool        `json:"archived"`
		Disabled         bool        `json:"disabled"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		License          interface{} `json:"license"`
		Forks            int         `json:"forks"`
		OpenIssues       int         `json:"open_issues"`
		Watchers         int         `json:"watchers"`
		DefaultBranch    string      `json:"default_branch"`
	} `json:"repository"`
	Organization struct {
		Login            string `json:"login"`
		ID               int    `json:"id"`
		NodeID           string `json:"node_id"`
		URL              string `json:"url"`
		ReposURL         string `json:"repos_url"`
		EventsURL        string `json:"events_url"`
		HooksURL         string `json:"hooks_url"`
		IssuesURL        string `json:"issues_url"`
		MembersURL       string `json:"members_url"`
		PublicMembersURL string `json:"public_members_url"`
		AvatarURL        string `json:"avatar_url"`
		Description      string `json:"description"`
	} `json:"organization"`
	Sender struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
}

func ReceiveGithubWebhook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data GithubHook
	if err := decoder.Decode(&data); err != nil {
		log.Fatal(err)
	}

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	queue := q.GetQueue()
	queue.Submit(&q.WebhookJob{ID: data.Rule.CreatedAt.String()})

	log.Printf("ID: %#v", vars)
}
