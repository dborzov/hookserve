package hookserve

type Event struct {
	Owner      string // The username of the owner of the repository
	Repo       Repo   // The name of the repository
	Branch     string // The branch the event took place on
	Commit     string // The head commit hash attached to the event
	Commits    []Commit
	Type       string // Can be either "pull_request" or "push"
	BaseOwner  string // For Pull Requests, contains the base owner
	BaseRepo   string // For Pull Requests, contains the base repo
	BaseBranch string // For Pull Requests, contains the base branch
}

type Repo struct {
	FullName string
	Url      string
}

type Commit struct {
	link    string
	message string
}
