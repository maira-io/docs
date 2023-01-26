# github edit-issue

Edit github issue

## Description

Edit github issue

## Synopsis

`github edit-issue [--site  <site>] --repo  <repo> --issue  <issue> [--title  <title>] [--body  <body>] [--state  <state>] [--milestone  <milestone>] [--labels  <labels>] [--assignees  <assignees>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `repo` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Fully Qualified Repository Name for issue  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--repo  "github_user1/test_repo_1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `issue` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; issue_number parameter  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--issue  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `title` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The title of the issue  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--title  "Write a user facing document on how to create new integration and commands"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `body` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The contents of the issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--body  "This document should be single place to get started and contribute integrations for the open source community"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `state` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; State of the issue. Either open or closed.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--state  "open"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `milestone` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The number of the milestone to associate this issue with.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--milestone  "milestone1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `labels` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Labels to associate with this issue.  Pass one or more Labels to replace the set of Labels on this Issue.  Send an empty array ([]) to clear all Labels from the Issue.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--labels  "bug/ui/@high"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `assignees` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Logins for Users to assign to this issue.  Pass one or more user logins to replace the set of assignees on this Issue.  Send an empty array ([]) to clear all assignees from the Issue.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--assignees  "assignee"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
x=! github edit-issue --site "localhost:8081" --cluster "github-default" --repo "prabhsimran-beehyv/testRepo" --issue 1
```
Output: 
```
{
  "url": "https://api.github.com/repos/prabhsimran-beehyv/testRepo/issues/1",
  "repository_url": "https://api.github.com/repos/prabhsimran-beehyv/testRepo",
  "labels_url": "https://api.github.com/repos/prabhsimran-beehyv/testRepo/issues/1/labels{/name}",
  "comments_url": "https://api.github.com/repos/prabhsimran-beehyv/testRepo/issues/1/comments",
  "events_url": "https://api.github.com/repos/prabhsimran-beehyv/testRepo/issues/1/events",
  "html_url": "https://github.com/prabhsimran-beehyv/testRepo/issues/1",
  "id": 1170705310,
  "node_id": "I_kwDOHAspuM5Fx4ue",
  "number": 1,
  "title": "testTitle1",
  "user": {
      "login": "prabhsimran-beehyv",
      "id": 100373722,
      "node_id": "U_kgDOBfuU2g",
      "avatar_url": "https://avatars.githubusercontent.com/u/100373722?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/prabhsimran-beehyv",
      "html_url": "https://github.com/prabhsimran-beehyv",
      "followers_url": "https://api.github.com/users/prabhsimran-beehyv/followers",
      "following_url": "https://api.github.com/users/prabhsimran-beehyv/following{/other_user}",
      "gists_url": "https://api.github.com/users/prabhsimran-beehyv/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/prabhsimran-beehyv/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/prabhsimran-beehyv/subscriptions",
      "organizations_url": "https://api.github.com/users/prabhsimran-beehyv/orgs",
      "repos_url": "https://api.github.com/users/prabhsimran-beehyv/repos",
      "events_url": "https://api.github.com/users/prabhsimran-beehyv/events{/privacy}",
      "received_events_url": "https://api.github.com/users/prabhsimran-beehyv/received_events",
      "type": "User",
      "site_admin": false
  },
  "labels": [],
  "state": "open",
  "locked": false,
  "assignee": null,
  "assignees": [],
  "milestone": null,
  "comments": 0,
  "created_at": "2022-03-16T08:39:31Z",
  "updated_at": "2022-03-16T08:39:31Z",
  "closed_at": null,
  "author_association": "OWNER",
  "active_lock_reason": null,
  "body": null,
  "closed_by": null,
  "reactions": {
      "url": "https://api.github.com/repos/prabhsimran-beehyv/testRepo/issues/1/reactions",
      "total_count": 0,
      "+1": 0,
      "-1": 0,
      "laugh": 0,
      "hooray": 0,
      "confused": 0,
      "heart": 0,
      "rocket": 0,
      "eyes": 0
  },
  "timeline_url": "https://api.github.com/repos/prabhsimran-beehyv/testRepo/issues/1/timeline",
  "performed_via_github_app": null
}
```

