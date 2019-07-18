package eventparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEventRequestBody(t *testing.T) {
	//given
	eventType := "star"
	eventTypeVersion := "v1"
	eventID := ""
	data := []byte(`[
		{
			"action": "deleted",
			"starred_at": null,
			"repository": {
			  "id": 197145944,
			  "node_id": "MDEwOlJlcG9zaXRvcnkxOTcxNDU5NDQ=",
			  "name": "hack-showcase",
			  "full_name": "ksputo/hack-showcase",
			  "private": false,
			  "owner": {
				"login": "ksputo",
				"id": 26189888,
				"node_id": "MDQ6VXNlcjI2MTg5ODg4",
				"avatar_url": "https://avatars2.githubusercontent.com/u/26189888?v=4",
				"gravatar_id": "",
				"url": "https://api.github.com/users/ksputo",
				"html_url": "https://github.com/ksputo",
				"followers_url": "https://api.github.com/users/ksputo/followers",
				"following_url": "https://api.github.com/users/ksputo/following{/other_user}",
				"gists_url": "https://api.github.com/users/ksputo/gists{/gist_id}",
				"starred_url": "https://api.github.com/users/ksputo/starred{/owner}{/repo}",
				"subscriptions_url": "https://api.github.com/users/ksputo/subscriptions",
				"organizations_url": "https://api.github.com/users/ksputo/orgs",
				"repos_url": "https://api.github.com/users/ksputo/repos",
				"events_url": "https://api.github.com/users/ksputo/events{/privacy}",
				"received_events_url": "https://api.github.com/users/ksputo/received_events",
				"type": "User",
				"site_admin": false
			  },
			  "html_url": "https://github.com/ksputo/hack-showcase",
			  "description": "The repository for a new Kyma showcase prepared by Hack Team",
			  "fork": true,
			  "url": "https://api.github.com/repos/ksputo/hack-showcase",
			  "forks_url": "https://api.github.com/repos/ksputo/hack-showcase/forks",
			  "keys_url": "https://api.github.com/repos/ksputo/hack-showcase/keys{/key_id}",
			  "collaborators_url": "https://api.github.com/repos/ksputo/hack-showcase/collaborators{/collaborator}",
			  "teams_url": "https://api.github.com/repos/ksputo/hack-showcase/teams",
			  "hooks_url": "https://api.github.com/repos/ksputo/hack-showcase/hooks",
			  "issue_events_url": "https://api.github.com/repos/ksputo/hack-showcase/issues/events{/number}",
			  "events_url": "https://api.github.com/repos/ksputo/hack-showcase/events",
			  "assignees_url": "https://api.github.com/repos/ksputo/hack-showcase/assignees{/user}",
			  "branches_url": "https://api.github.com/repos/ksputo/hack-showcase/branches{/branch}",
			  "tags_url": "https://api.github.com/repos/ksputo/hack-showcase/tags",
			  "blobs_url": "https://api.github.com/repos/ksputo/hack-showcase/git/blobs{/sha}",
			  "git_tags_url": "https://api.github.com/repos/ksputo/hack-showcase/git/tags{/sha}",
			  "git_refs_url": "https://api.github.com/repos/ksputo/hack-showcase/git/refs{/sha}",
			  "trees_url": "https://api.github.com/repos/ksputo/hack-showcase/git/trees{/sha}",
			  "statuses_url": "https://api.github.com/repos/ksputo/hack-showcase/statuses/{sha}",
			  "languages_url": "https://api.github.com/repos/ksputo/hack-showcase/languages",
			  "stargazers_url": "https://api.github.com/repos/ksputo/hack-showcase/stargazers",
			  "contributors_url": "https://api.github.com/repos/ksputo/hack-showcase/contributors",
			  "subscribers_url": "https://api.github.com/repos/ksputo/hack-showcase/subscribers",
			  "subscription_url": "https://api.github.com/repos/ksputo/hack-showcase/subscription",
			  "commits_url": "https://api.github.com/repos/ksputo/hack-showcase/commits{/sha}",
			  "git_commits_url": "https://api.github.com/repos/ksputo/hack-showcase/git/commits{/sha}",
			  "comments_url": "https://api.github.com/repos/ksputo/hack-showcase/comments{/number}",
			  "issue_comment_url": "https://api.github.com/repos/ksputo/hack-showcase/issues/comments{/number}",
			  "contents_url": "https://api.github.com/repos/ksputo/hack-showcase/contents/{+path}",
			  "compare_url": "https://api.github.com/repos/ksputo/hack-showcase/compare/{base}...{head}",
			  "merges_url": "https://api.github.com/repos/ksputo/hack-showcase/merges",
			  "archive_url": "https://api.github.com/repos/ksputo/hack-showcase/{archive_format}{/ref}",
			  "downloads_url": "https://api.github.com/repos/ksputo/hack-showcase/downloads",
			  "issues_url": "https://api.github.com/repos/ksputo/hack-showcase/issues{/number}",
			  "pulls_url": "https://api.github.com/repos/ksputo/hack-showcase/pulls{/number}",
			  "milestones_url": "https://api.github.com/repos/ksputo/hack-showcase/milestones{/number}",
			  "notifications_url": "https://api.github.com/repos/ksputo/hack-showcase/notifications{?since,all,participating}",
			  "labels_url": "https://api.github.com/repos/ksputo/hack-showcase/labels{/name}",
			  "releases_url": "https://api.github.com/repos/ksputo/hack-showcase/releases{/id}",
			  "deployments_url": "https://api.github.com/repos/ksputo/hack-showcase/deployments",
			  "created_at": "2019-07-16T07:44:59Z",
			  "updated_at": "2019-07-17T11:53:39Z",
			  "pushed_at": "2019-07-04T12:44:24Z",
			  "git_url": "git://github.com/ksputo/hack-showcase.git",
			  "ssh_url": "git@github.com:ksputo/hack-showcase.git",
			  "clone_url": "https://github.com/ksputo/hack-showcase.git",
			  "svn_url": "https://github.com/ksputo/hack-showcase",
			  "homepage": null,
			  "size": 8,
			  "stargazers_count": 0,
			  "watchers_count": 0,
			  "language": null,
			  "has_issues": false,
			  "has_projects": true,
			  "has_downloads": true,
			  "has_wiki": false,
			  "has_pages": false,
			  "forks_count": 0,
			  "mirror_url": null,
			  "archived": false,
			  "disabled": false,
			  "open_issues_count": 0,
			  "license": {
				"key": "apache-2.0",
				"name": "Apache License 2.0",
				"spdx_id": "Apache-2.0",
				"url": "https://api.github.com/licenses/apache-2.0",
				"node_id": "MDc6TGljZW5zZTI="
			  },
			  "forks": 0,
			  "open_issues": 0,
			  "watchers": 0,
			  "default_branch": "master"
			},
			"sender": {
			  "login": "ksputo",
			  "id": 26189888,
			  "node_id": "MDQ6VXNlcjI2MTg5ODg4",
			  "avatar_url": "https://avatars2.githubusercontent.com/u/26189888?v=4",
			  "gravatar_id": "",
			  "url": "https://api.github.com/users/ksputo",
			  "html_url": "https://github.com/ksputo",
			  "followers_url": "https://api.github.com/users/ksputo/followers",
			  "following_url": "https://api.github.com/users/ksputo/following{/other_user}",
			  "gists_url": "https://api.github.com/users/ksputo/gists{/gist_id}",
			  "starred_url": "https://api.github.com/users/ksputo/starred{/owner}{/repo}",
			  "subscriptions_url": "https://api.github.com/users/ksputo/subscriptions",
			  "organizations_url": "https://api.github.com/users/ksputo/orgs",
			  "repos_url": "https://api.github.com/users/ksputo/repos",
			  "events_url": "https://api.github.com/users/ksputo/events{/privacy}",
			  "received_events_url": "https://api.github.com/users/ksputo/received_events",
			  "type": "User",
			  "site_admin": false
			}
		  }	
	`)

	//when

	t.Run("When event request is created", func(t *testing.T) {
		actual := GetEventRequestPayload(eventType, eventTypeVersion, eventID, data)

		//then
		t.Run("Then struct should be created", func(t *testing.T) {
			assert.NotNil(t, actual)
		})
		t.Run("Then returned object should have proper type", func(t *testing.T) {
			assert.IsType(t, EventRequestPayload{}, actual)
		})

		t.Run("Then all struct fields should be filled", func(t *testing.T) {
			assert.NotNil(t, actual.EventType)
			assert.NotNil(t, actual.EventTypeVersion)
			assert.NotNil(t, actual.EventID)
			assert.NotNil(t, actual.EventTime)
			assert.NotNil(t, actual.Data)

		})
	})

}
