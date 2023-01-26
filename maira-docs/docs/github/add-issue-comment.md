# github add-issue-comment

Add a comment to a github issue

## Description

Add a comment to a github issue

## Synopsis

`github add-issue-comment [--site  <site>] --repo  <repo> --issue  <issue> [--comment  <comment>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; issue number  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--issue  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `comment` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; comment to be added  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--comment  "Write a user facing document on how to create new integration and commands"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x=!github add-issue-comment --site "localhost:8080" --repo "maira-io/apiserver" --issue 108 --comment "new comment"
```
Output: 
```
{
    "url": "https://api.github.com/repos/maira-io/apiserver/issues/comments/1159210585",
    "html_url": "https://github.com/maira-io/apiserver/issues/108#issuecomment-1159210585",
    "issue_url": "https://api.github.com/repos/maira-io/apiserver/issues/108",
    "id": 1159210585,
    "node_id": "IC_kwDOE8hIb85FGCZZ",
    "user": {
        "login": "sbansal7",
        "id": 1282760,
        "node_id": "MDQ6VXNlcjEyODI3NjA=",
        "avatar_url": "https://avatars.githubusercontent.com/u/1282760?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/sbansal7",
        "html_url": "https://github.com/sbansal7",
        "followers_url": "https://api.github.com/users/sbansal7/followers",
        "following_url": "https://api.github.com/users/sbansal7/following{/other_user}",
        "gists_url": "https://api.github.com/users/sbansal7/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/sbansal7/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/sbansal7/subscriptions",
        "organizations_url": "https://api.github.com/users/sbansal7/orgs",
        "repos_url": "https://api.github.com/users/sbansal7/repos",
        "events_url": "https://api.github.com/users/sbansal7/events{/privacy}",
        "received_events_url": "https://api.github.com/users/sbansal7/received_events",
        "type": "User",
        "site_admin": false
    },
    "created_at": "2022-06-17T20:35:18Z",
    "updated_at": "2022-06-17T20:35:18Z",
    "author_association": "CONTRIBUTOR",
    "body": "new comment",
    "reactions": {
        "url": "https://api.github.com/repos/maira-io/apiserver/issues/comments/1159210585/reactions",
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
    "performed_via_github_app": null
}
```

