# github list-issue

List github issues

## Description

List github issues

## Synopsis

`github list-issue [--site  <site>] --repo  <repo> [--milestone  <milestone>] [--state  <state>] [--assignee  <assignee>] [--creator  <creator>] [--mentioned  <mentioned>] [--labels  <labels>] [--sort  <sort>] [--direction  <direction>] [--since  <since>] [--per_page  <per_page>] [--page  <page>]`

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


#### `milestone` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The number of the milestone to associate this issue with.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--milestone  "milestone1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `state` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; State of the issue. Either open or closed.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--state  "open"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `assignee` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Can be the name of a user. Pass in none for issues with no assigned user, and * for issues assigned to any user.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--assignee  "assignee"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `creator` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The user that created the issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--creator  "github_user1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `mentioned` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A user that is mentioned in the issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--mentioned  "github_user1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `labels` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A list of comma separated label names.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--labels  "bug,ui,@high"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `sort` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; What to sort results by. Default- created  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--sort  "created/updated/comments"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `direction` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; One of asc (ascending) or desc (descending). Default- desc  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--direction  "asc/desc"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `since` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Only show notifications updated after the given time. This is a timestamp in ISO 8601 format- YYYY-MM-DDTHH:MM:SSZ
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--since  "2022-07-01T12:00:00Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `per_page` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Results per page (max 100). Default- 30  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--per_page  100`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `page` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Page number of the results to fetch. Default- 1  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--page  2`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! github list-issue --site "localhost:8081" --cluster "github-default" --repo "prabhsimran-beehyv/testRepo"
```
Output: 
```
NUMBER  LABELS  STATE   ASSIGNEES               MILESTONE       TITLE      
     4  bug     open    prabhsimran-beehyv      null            testTitle5
     3  -       open    -                       null            testTitle3
     2  -       open    -                       null            testTitle1
     1  -       open    -                       null            testTitle1
```

