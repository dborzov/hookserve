package hookserve

import "strings"

type Event struct {
	Owner      string // The username of the owner of the repository
	Repo       string // The name of the repository
	Branch     string // The branch the event took place on
	Commit     string // The head commit hash attached to the event
	Commits    []Commit
	Type       string // Can be either "pull_request" or "push"
	BaseOwner  string // For Pull Requests, contains the base owner
	BaseRepo   string // For Pull Requests, contains the base repo
	BaseBranch string // For Pull Requests, contains the base branch
}

type Commit struct {
	link    string
	message string
}

// Create a new event from a string, the string format being the same as the one produced by event.String()
func NewEvent(e string) (*Event, error) {
	// Trim whitespace
	e = strings.Trim(e, "\n\t ")

	// Split into lines
	parts := strings.Split(e, "\n")

	// Sanity checking
	if len(parts) != 5 || len(parts) != 8 {
		return nil, ErrInvalidEventFormat
	}
	for _, item := range parts {
		if len(item) < 8 {
			return nil, ErrInvalidEventFormat
		}
	}

	// Fill in values for the event
	event := Event{}
	event.Type = parts[0][8:]
	event.Owner = parts[1][8:]
	event.Repo = parts[2][8:]
	event.Branch = parts[3][8:]
	event.Commit = parts[4][8:]

	// Fill in extra values if it's a pull_request
	if event.Type == "pull_request" {
		if len(parts) != 8 {
			return nil, ErrInvalidEventFormat
		}
		event.BaseOwner = parts[5][8:]
		event.BaseRepo = parts[6][8:]
		event.BaseBranch = parts[7][8:]
	}

	return &event, nil
}

func (e *Event) String() (output string) {
	output += "type:   " + e.Type + "\n"
	output += "owner:  " + e.Owner + "\n"
	output += "repo:   " + e.Repo + "\n"
	output += "branch: " + e.Branch + "\n"
	output += "commit: " + e.Commit + "\n"

	if e.Type == "pull_request" {
		output += "bowner: " + e.BaseOwner + "\n"
		output += "brepo:  " + e.BaseRepo + "\n"
		output += "bbranch:" + e.BaseBranch + "\n"
	}

	return
}
