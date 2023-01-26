# github create-issue

Create github issue

## Description

Create github issue

## Synopsis

`github create-issue [--site  <site>] --repo  <repo> --title  <title> [--body  <body>] [--milestone  <milestone>] [--labels  <labels>] [--assignees  <assignees>]`

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


#### `title` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The title of the issue  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--title  "Write a user facing document on how to create new integration and commands"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `body` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The contents of the issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--body  "This document should be single place to get started and contribute integrations for the open source community"`

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
! github create-issue --site "localhost:8081" --title "testTitle1" --cluster "github-default" --repo "prabhsimran-beehyv/testRepo"
```
Output: 
```
Issue number 1 created
```

