package hookserve

import "github.com/bmatsuo/go-jsontree"

// LoadEvent takes a JSON and parses it into an Event
func LoadEvent(request *jsontree.JsonTree, eventType string) (event *Event, err error) {
	// Parse the request and build the Event
	event = &Event{Repo: Repo{}}
	event.Type = eventType
	event.Repo.FullName, err = request.Get("repository").Get("fullname").String()
	event.Repo.Url, err = request.Get("repository").Get("html_url").String()

	if err != nil {
		return event, err
	}

	switch eventType {
	case "push":
		err = parsePush(event, request)
	case "pull_request":
		err = parsePullRequst(event, request)
	}
	return event, err
}
