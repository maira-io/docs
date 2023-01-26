# bitbucket create-issue

Create bitbucket issue

## Description

Create bitbucket issue

## Synopsis

`bitbucket create-issue [--site  <site>] --repo  <repo> [--title  <title>] [--content  <content>] [--milestone  <milestone>] [--type  <type>] [--priority  <priority>] [--assignee  <assignee>]`

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


#### `assignee` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; User to assign to this issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--assignee  "assignee"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! bitbucket create-issue --site "localhost:8081" --title "testTitle1" --repo "sbansal7/repo1"
```
Output: 
```
ID	TYPE	PRIORITY	STATE	ASSIGNEE	MILESTONE	TITLE
 4	bug 	major   	new  	-       	null     	testTitle1
```

