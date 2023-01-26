# bitbucket add-issue-comment

Add a comment to a bitbucket issue

## Description

Add a comment to a bitbucket issue

## Synopsis

`bitbucket add-issue-comment [--site  <site>] --repo  <repo> --issue  <issue> [--comment  <comment>]`

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
x=!bitbucket add-issue-comment --site "localhost:8080" --repo "sbansal7/repo1" --issue 1 --comment "new comment"
```
Output: 
```
{
    "type": "issue_comment",
    "id": 64023557,
    "created_on": "2022-08-30T15:38:32.472009+00:00",
    "updated_on": null,
    "content": {
        "type": "rendered",
        "raw": "new comment",
        "markup": "markdown",
        "html": "<p>new comment</p>"
    },
    "user": {
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
    "issue": {
        "type": "issue",
        "id": 1,
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
                "href": "https://api.bitbucket.org/2.0/repositories/sbansal7/repo1/issues/1"
            }
        },
        "title": "test issue"
    },
    "links": {
        "self": {
            "href": "https://api.bitbucket.org/2.0/repositories/sbansal7/repo1/issues/1/comments/64023557"
        },
        "html": {
            "href": "https://bitbucket.org/sbansal7/repo1/issues/1#comment-64023557"
        }
    }
}
```

