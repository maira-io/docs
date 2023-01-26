# bitbucket edit-issue

Edit bitbucket issue

## Description

Edit bitbucket issue

## Synopsis

`bitbucket edit-issue [--site  <site>] --repo  <repo> --issue  <issue> [--title  <title>] [--content  <content>] [--milestone  <milestone>] [--type  <type>] [--priority  <priority>] [--state  <state>] [--assignee  <assignee>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `repo` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Fully Qualified Repository Name for issue  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--repo  "bitbucket_user1/test_repo_1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `issue` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; issue id  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--issue  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `title` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The title of the issue  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--title  "Write a user facing document on how to create new integration and commands"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `content` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The contents of the issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--content  "This document should be single place to get started and contribute integrations for the open source community"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `milestone` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The number of the milestone to associate this issue with.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--milestone  "milestone1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Type of the issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--type  "bug"`
 ,  `--type  "enhancement"`
 ,  `--type  "proposal"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `priority` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Priority of the issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--priority  "major"`
 ,  `--priority  "minor"`
 ,  `--priority  "trivial"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `state` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; State of the issue. Either open or closed.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--state  "open"`
 ,  `--state  "new"`
 ,  `--state  "closed"`
 ,  `--state  "on hold"`
 ,  `--state  "resolved"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `assignee` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; User to assign to this issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--assignee  "assignee"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x=! bitbucket edit-issue --site "localhost:8080" --repo "sbansal7/repo1" --issue 1 --assignee "sbansal7"
```
Output: 
```
{
    "type": "issue",
    "id": 5,
    "repository": {
        "type": "repository",
        "full_name": "sbansal7/repo1",
        "links": {
            "self": {
                "href": "https://api.bitbucket.org/2.0/repositories/sbansal7/repo1"
            },
            "html": {
                "href": "https://bitbucket.org/sbansal7/repo1"
            },
            "avatar": {
                "href": "https://bytebucket.org/ravatar/%7Bc7b44412-1430-4789-a512-4c8cadf3f6f7%7D?ts=default"
            }
        },
        "name": "repo1",
        "uuid": "{c7b44412-1430-4789-a512-4c8cadf3f6f7}"
    },
    "links": {
        "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/sbansal7/repo1/issues/5"
        }
    },
    "title": "testTitle1",
    "content": {
        "type": "rendered",
        "raw": "",
        "markup": "markdown",
        "html": ""
    },
    "reporter": {
        "display_name": "Sachin Bansal",
        "links": {
            "self": {
                "href": "https://api.bitbucket.org/2.0/users/%7Bfa271433-988c-4a42-932b-380426e46f8e%7D"
            },
            "avatar": {
                "href": "https://secure.gravatar.com/avatar/bbcb732b26ae24604238af9c110f5f76?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FSB-2.png"
            },
            "html": {
                "href": "https://bitbucket.org/%7Bfa271433-988c-4a42-932b-380426e46f8e%7D/"
            }
        },
        "type": "user",
        "uuid": "{fa271433-988c-4a42-932b-380426e46f8e}",
        "account_id": "630d04a0316bbc56c42401e9",
        "nickname": "Sachin Bansal"
    },
    "assignee": null,
    "created_on": "2022-08-30T15:42:14.325947+00:00",
    "edited_on": null,
    "updated_on": "2022-08-30T15:42:14.325947+00:00",
    "state": "new",
    "kind": "bug",
    "milestone": null,
    "component": null,
    "priority": "major",
    "version": null,
    "votes": 0,
    "watches": null
}
```

