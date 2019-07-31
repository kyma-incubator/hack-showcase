package registerservice

import (
	"time"

	"github.com/kyma-incubator/hack-showcase/github-connector/internal/apperrors"
	"github.com/kyma-incubator/hack-showcase/github-connector/internal/model"
	log "github.com/sirupsen/logrus"
)

var jsonBody = model.ServiceDetails{
	Provider:    "kyma",
	Name:        "github-connector",
	Description: "Boilerplate for GitHub connector",
	API: &model.API{
		TargetURL:        "https://api.github.com",
		SpecificationURL: "https://raw.githubusercontent.com/colunira/github-openapi/master/githubopenAPI.json",
	},
	Events: &model.Events{
		Spec: []byte(`{
			"asyncapi": "1.0.0",
			"info": {
				"title": "github-events",
				"version": "v1",
				"description": "Github Events v1"
			},
			"topics": {
				"issuesevent.opened.v1": {
					"subscribe": {
						"summary": "New issue opened on github repository",
						"payload": {
							"type": "object",
							"required": [
								"action"
							],
							"properties": {
								"action": {
									"type": "string",
									"example": "edited",
									"description": "The action that was performed.",
									"title": "Action"
								},
								"issue": {
									"$ref": "#/components/schemas/issue"
								},
								"changes": {
									"type": "object"
								},
								"repository": {
									"$ref": "#/components/schemas/repository"
								},
								"sender": {
									"$ref": "#/components/schemas/user"
								}
							}
						}
					}
				},
				"pullrequestreviewevent.submitted.v1": {
					"subscribe": {
						"summary": "Triggered when a pull request review is submitted into a non-pending state.",
						"payload": {
							"type": "object",
							"required": [
								"action"
							],
							"properties": {
								"action": {
									"type": "string",
									"example": "submitted",
									"description": "The action that was performed.",
									"title": "Action"
								},
								"review": {
									"$ref": "#/components/schemas/review"
								},
								"pull_request": {
									"$ref": "#/components/schemas/pull_request"
								},
								"repository": {
									"$ref": "#/components/schemas/repository"
								},
								"sender": {
									"$ref": "#/components/schemas/user"
								}
							}
						}
					}
				}
			},
			"components": {
				"schemas": {
					"issue": {
						"type": "object",
						"properties": {
							"url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/1"
							},
							"repository_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World"
							},
							"labels_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/1/labels{/name}"
							},
							"comments_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/1/comments"
							},
							"events_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/1/events"
							},
							"html_url": {
								"type": "string",
								"example": "https://github.com/Codertocat/Hello-World/issues/1"
							},
							"id": {
								"type": "integer",
								"example": 444500041
							},
							"node_id": {
								"type": "string",
								"example": "MDU6SXNzdWU0NDQ1MDAwNDE="
							},
							"number": {
								"type": "integer",
								"example": 1
							},
							"title": {
								"type": "string",
								"example": "Spelling error in the README file"
							},
							"user": {
								"$ref": "#/components/schemas/user"
							},
							"labels": {
								"type": "array",
								"items": {
									"$ref": "#/components/schemas/label"
								}
							},
							"state": {
								"type": "string",
								"example": "open"
							},
							"locked": {
								"type": "boolean",
								"example": false
							},
							"assignee": {
								"$ref": "#/components/schemas/user"
							},
							"assignees": {
								"type": "array",
								"items": {
									"$ref": "#/components/schemas/user"
								}
							},
							"milestone": {
								"$ref": "#/components/schemas/milestone"
							},
							"comments": {
								"type": "integer",
								"example": 0
							},
							"created_at": {
								"type": "string",
								"example": "2019-05-15T15:20:18Z"
							},
							"updated_at": {
								"type": "string",
								"example": "2019-05-15T15:20:18Z"
							},
							"closed_at": {
								"type": "string",
								"example": null
							},
							"author_association": {
								"type": "string",
								"example": "OWNER"
							},
							"body": {
								"type": "string",
								"example": "It looks like you accidently spelled 'commit' with two 't's."
							}
						}
					},
					"user": {
						"type": "object",
						"properties": {
							"login": {
								"type": "string",
								"example": "Codertocat"
							},
							"id": {
								"type": "integer",
								"example": 21031067
							},
							"node_id": {
								"type": "string",
								"example": "MDQ6VXNlcjIxMDMxMDY3"
							},
							"avatar_url": {
								"type": "string",
								"example": "https://avatars1.githubusercontent.com/u/21031067?v=4"
							},
							"gravatar_id": {
								"type": "string",
								"example": ""
							},
							"url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat"
							},
							"html_url": {
								"type": "string",
								"example": "https://github.com/Codertocat"
							},
							"followers_url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat/followers"
							},
							"following_url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat/following{/other_user}"
							},
							"gists_url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat/gists{/gist_id}"
							},
							"starred_url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat/starred{/owner}{/repo}"
							},
							"subscriptions_url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat/subscriptions"
							},
							"organizations_url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat/orgs"
							},
							"repos_url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat/repos"
							},
							"events_url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat/events{/privacy}"
							},
							"received_events_url": {
								"type": "string",
								"example": "https://api.github.com/users/Codertocat/received_events"
							},
							"type": {
								"type": "string",
								"example": "User"
							},
							"site_admin": {
								"type": "boolean",
								"example": false
							}
						}
					},
					"label": {
						"type": "object",
						"properties": {
							"id": {
								"type": "integer",
								"example": 1362934389
							},
							"node_id": {
								"type": "string",
								"example": "MDU6TGFiZWwxMzYyOTM0Mzg5"
							},
							"url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/labels/bug"
							},
							"name": {
								"type": "string",
								"example": "bug"
							},
							"color": {
								"type": "string",
								"example": "d73a4a"
							},
							"default": {
								"type": "boolean",
								"example": true
							}
						}
					},
					"milestone": {
						"type": "object",
						"properties": {
							"url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/milestones/1"
							},
							"html_url": {
								"type": "string",
								"example": "https://github.com/Codertocat/Hello-World/milestone/1"
							},
							"labels_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/milestones/1/labels"
							},
							"id": {
								"type": "integer",
								"example": 4317517
							},
							"node_id": {
								"type": "string",
								"example": "MDk6TWlsZXN0b25lNDMxNzUxNw=="
							},
							"number": {
								"type": "integer",
								"example": 1
							},
							"title": {
								"type": "string",
								"example": "v1.0"
							},
							"description": {
								"type": "string",
								"example": "Add new space flight simulator"
							},
							"creator": {
								"$ref": "#/components/schemas/user"
							},
							"open_issues": {
								"type": "integer",
								"example": 1
							},
							"closed_issues": {
								"type": "integer",
								"example": 0
							},
							"state": {
								"type": "string",
								"example": "closed"
							},
							"created_at": {
								"type": "string",
								"example": "2019-05-15T15:20:17Z"
							},
							"updated_at": {
								"type": "string",
								"example": "2019-05-15T15:20:18Z"
							},
							"due_on": {
								"type": "string",
								"example": "2019-05-23T07:00:00Z"
							},
							"closed_at": {
								"type": "string",
								"example": "2019-05-15T15:20:18Z"
							}
						}
					},
					"repository": {
						"type": "object",
						"properties": {
							"id": {
								"type": "integer",
								"example": 186853002
							},
							"node_id": {
								"type": "string",
								"example": "MDEwOlJlcG9zaXRvcnkxODY4NTMwMDI="
							},
							"name": {
								"type": "string",
								"example": "Hello-World"
							},
							"full_name": {
								"type": "string",
								"example": "Codertocat/Hello-World"
							},
							"private": {
								"type": "boolean",
								"example": false
							},
							"owner": {
								"$ref": "#/components/schemas/user"
							},
							"html_url": {
								"type": "string",
								"example": "https://github.com/Codertocat/Hello-World"
							},
							"description": {
								"type": "string",
								"example": null
							},
							"fork": {
								"type": "boolean",
								"example": false
							},
							"url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World"
							},
							"forks_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/forks"
							},
							"keys_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/keys{/key_id}"
							},
							"collaborators_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/collaborators{/collaborator}"
							},
							"teams_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/teams"
							},
							"hooks_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/hooks"
							},
							"issue_events_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/events{/number}"
							},
							"events_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/events"
							},
							"assignees_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/assignees{/user}"
							},
							"branches_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/branches{/branch}"
							},
							"tags_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/tags"
							},
							"blobs_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/git/blobs{/sha}"
							},
							"git_tags_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/git/tags{/sha}"
							},
							"git_refs_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/git/refs{/sha}"
							},
							"trees_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/git/trees{/sha}"
							},
							"statuses_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/statuses/{sha}"
							},
							"languages_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/languages"
							},
							"stargazers_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/stargazers"
							},
							"contributors_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/contributors"
							},
							"subscribers_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/subscribers"
							},
							"subscription_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/subscription"
							},
							"commits_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/commits{/sha}"
							},
							"git_commits_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/git/commits{/sha}"
							},
							"comments_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/comments{/number}"
							},
							"issue_comment_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/comments{/number}"
							},
							"contents_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/contents/{+path}"
							},
							"compare_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/compare/{base}...{head}"
							},
							"merges_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/merges"
							},
							"archive_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/{archive_format}{/ref}"
							},
							"downloads_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/downloads"
							},
							"issues_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/issues{/number}"
							},
							"pulls_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls{/number}"
							},
							"milestones_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/milestones{/number}"
							},
							"notifications_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/notifications{?since,all,participating}"
							},
							"labels_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/labels{/name}"
							},
							"releases_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/releases{/id}"
							},
							"deployments_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/deployments"
							},
							"created_at": {
								"type": "string",
								"example": "2019-05-15T15:19:25Z"
							},
							"updated_at": {
								"type": "string",
								"example": "2019-05-15T15:19:27Z"
							},
							"pushed_at": {
								"type": "string",
								"example": "2019-05-15T15:20:13Z"
							},
							"git_url": {
								"type": "string",
								"example": "git://github.com/Codertocat/Hello-World.git"
							},
							"ssh_url": {
								"type": "string",
								"example": "git@github.com:Codertocat/Hello-World.git"
							},
							"clone_url": {
								"type": "string",
								"example": "https://github.com/Codertocat/Hello-World.git"
							},
							"svn_url": {
								"type": "string",
								"example": "https://github.com/Codertocat/Hello-World"
							},
							"homepage": {
								"type": "string",
								"example": null
							},
							"size": {
								"type": "integer",
								"example": 0
							},
							"stargazers_count": {
								"type": "integer",
								"example": 0
							},
							"watchers_count": {
								"type": "integer",
								"example": 0
							},
							"language": {
								"type": "string",
								"example": "Ruby"
							},
							"has_issues": {
								"type": "boolean",
								"example": true
							},
							"has_projects": {
								"type": "boolean",
								"example": true
							},
							"has_downloads": {
								"type": "boolean",
								"example": true
							},
							"has_wiki": {
								"type": "boolean",
								"example": true
							},
							"has_pages": {
								"type": "boolean",
								"example": true
							},
							"forks_count": {
								"type": "integer",
								"example": 0
							},
							"mirror_url": {
								"type": "string",
								"example": null
							},
							"archived": {
								"type": "boolean",
								"example": false
							},
							"disabled": {
								"type": "boolean",
								"example": false
							},
							"open_issues_count": {
								"type": "integer",
								"example": 1
							},
							"license": {
								"type": "string",
								"example": null
							},
							"forks": {
								"type": "integer",
								"example": 0
							},
							"open_issues": {
								"type": "integer",
								"example": 1
							},
							"watchers": {
								"type": "integer",
								"example": 0
							},
							"default_branch": {
								"type": "string",
								"example": "master"
							}
						}
					},
					"review": {
						"type": "object",
						"properties": {
							"id": {
								"type": "integer",
								"example": 237895671
							},
							"node_id": {
								"type": "string",
								"example": "MDE3OlB1bGxSZXF1ZXN0UmV2aWV3MjM3ODk1Njcx"
							},
							"user": {
								"$ref": "#/components/schemas/user"
							},
							"body": {
								"type": "string",
								"example": null
							},
							"commit_id": {
								"type": "string",
								"example": "ec26c3e57ca3a959ca5aad62de7213c562f8c821"
							},
							"submitted_at": {
								"type": "string",
								"example": "2019-05-15T15:20:38Z"
							},
							"state": {
								"type": "string",
								"example": "commented"
							},
							"html_url": {
								"type": "string",
								"example": "https://github.com/Codertocat/Hello-World/pull/2#pullrequestreview-237895671"
							},
							"pull_request_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2"
							},
							"author_association": {
								"type": "string",
								"example": "OWNER"
							},
							"_links": {
								"type": "object",
								"properties": {
									"html": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://github.com/Codertocat/Hello-World/pull/2#pullrequestreview-237895671"
											}
										}
									},
									"pull_request": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2"
											}
										}
									}
								}
							}
						}
					},
					"pull_request": {
						"type": "object",
						"properties": {
							"url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2"
							},
							"id": {
								"type": "integer",
								"example": 279147437
							},
							"node_id": {
								"type": "string",
								"example": "MDExOlB1bGxSZXF1ZXN0Mjc5MTQ3NDM3"
							},
							"html_url": {
								"type": "string",
								"example": "https://github.com/Codertocat/Hello-World/pull/2"
							},
							"diff_url": {
								"type": "string",
								"example": "https://github.com/Codertocat/Hello-World/pull/2.diff"
							},
							"patch_url": {
								"type": "string",
								"example": "https://github.com/Codertocat/Hello-World/pull/2.patch"
							},
							"issue_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/2"
							},
							"number": {
								"type": "integer",
								"example": 2
							},
							"state": {
								"type": "string",
								"example": "open"
							},
							"locked": {
								"type": "boolean",
								"example": false
							},
							"title": {
								"type": "string",
								"example": "Update the README with new information."
							},
							"user": {
								"$ref": "#/components/schemas/user"
							},
							"body": {
								"type": "string",
								"example": "This is a pretty simple change that we need to pull into master."
							},
							"created_at": {
								"type": "string",
								"example": "2019-05-15T15:20:33Z"
							},
							"updated_at": {
								"type": "string",
								"example": "2019-05-15T15:20:38Z"
							},
							"closed_at": {
								"type": "string",
								"example": null
							},
							"merged_at": {
								"type": "string",
								"example": null
							},
							"merge_commit_sha": {
								"type": "string",
								"example": "c4295bd74fb0f4fda03689c3df3f2803b658fd85"
							},
							"assignee": {
								"$ref": "#/components/schemas/user",
								"example": null
							},
							"assignees": {
								"type": "array",
								"items": {
									"$ref": "#/components/schemas/user"
								},
								"example": []
							},
							"requested_reviewers": {
								"type": "array",
								"items": {
									"$ref": "#/components/schemas/user"
								},
								"example": []
							},
							"requested_teams": {
								"type": "array",
								"items": {
									"type": "object"
								},
								"example": []
							},
							"labels": {
								"type": "array",
								"items": {
									"$ref": "#/components/schemas/label"
								},
								"example": []
							},
							"milestone": {
								"$ref": "#/components/schemas/milestone",
								"example": null
							},
							"commits_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2/commits"
							},
							"review_comments_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2/comments"
							},
							"review_comment_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/comments{/number}"
							},
							"comments_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/2/comments"
							},
							"statuses_url": {
								"type": "string",
								"example": "https://api.github.com/repos/Codertocat/Hello-World/statuses/ec26c3e57ca3a959ca5aad62de7213c562f8c821"
							},
							"head": {
								"$ref": "#/components/schemas/head"
							},
							"base": {
								"$ref": "#/components/schemas/base"
							},
							"_links": {
								"type": "object",
								"properties": {
									"self": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2"
											}
										}
									},
									"html": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://github.com/Codertocat/Hello-World/pull/2"
											}
										}
									},
									"issue": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/2"
											}
										}
									},
									"comments": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://api.github.com/repos/Codertocat/Hello-World/issues/2/comments"
											}
										}
									},
									"review_comments": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2/comments"
											}
										}
									},
									"review_comment": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/comments{/number}"
											}
										}
									},
									"commits": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2/commits"
											}
										}
									},
									"statuses": {
										"type": "object",
										"properties": {
											"href": {
												"type": "string",
												"example": "https://api.github.com/repos/Codertocat/Hello-World/statuses/ec26c3e57ca3a959ca5aad62de7213c562f8c821"
											}
										}
									}
								}
							},
							"author_association": {
								"type": "string",
								"example": "OWNER"
							}
						}
					},
					"head": {
						"type": "object",
						"properties": {
							"label": {
								"type": "string",
								"example": "Codertocat:changes"
							},
							"ref": {
								"type": "string",
								"example": "changes"
							},
							"sha": {
								"type": "string",
								"example": "ec26c3e57ca3a959ca5aad62de7213c562f8c821"
							},
							"user": {
								"$ref": "#/components/schemas/user"
							},
							"repo": {
								"$ref": "#/components/schemas/repository"
							}
						}
					},
					"base": {
						"type": "object",
						"properties": {
							"label": {
								"type": "string",
								"example": "Codertocat:master"
							},
							"ref": {
								"type": "string",
								"example": "master"
							},
							"sha": {
								"type": "string",
								"example": "f95f852bd8fca8fcc58a9a2d6c842781e32a215e"
							},
							"user": {
								"$ref": "#/components/schemas/user"
							},
							"repo": {
								"$ref": "#/components/schemas/repository"
							}
						}
					}
				}
			}
		}`),
	},
}

var url = "http://application-registry-external-api.kyma-integration.svc.cluster.local:8081/github-connector-app/v1/metadata/services"

//RegisterService - register service in Kyma and get a response
func RegisterService() (string, apperrors.AppError) {

	var id string
	var err error
	for i := 0; i < 10; i++ {
		id, err = SendRegisterRequest(jsonBody, url)
		if err == nil {
			break
		}
		log.Warn(err.Error())

		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return "", apperrors.Internal("While trying to register service: %s", err.Error())
	}
	return id, nil
}
