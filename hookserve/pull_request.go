package hookserve

import "github.com/bmatsuo/go-jsontree"

func parsePullRequst(event *Event, request *jsontree.JsonTree) (err error) {
	// Fill in values
	event.Owner, err = request.Get("pull_request").Get("head").Get("repo").Get("owner").Get("login").String()
	if err != nil {
		return err
	}
	event.Repo, err = request.Get("pull_request").Get("head").Get("repo").Get("name").String()
	if err != nil {
		return err
	}
	event.Branch, err = request.Get("pull_request").Get("head").Get("ref").String()
	if err != nil {
		return err
	}
	event.Commit, err = request.Get("pull_request").Get("head").Get("sha").String()
	if err != nil {
		return err
	}
	event.BaseOwner, err = request.Get("pull_request").Get("base").Get("repo").Get("owner").Get("login").String()
	if err != nil {
		return err
	}
	event.BaseRepo, err = request.Get("pull_request").Get("base").Get("repo").Get("name").String()
	if err != nil {
		return err
	}
	event.BaseBranch, err = request.Get("pull_request").Get("base").Get("ref").String()
	if err != nil {
		return err
	}
	return nil
}
