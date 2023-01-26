# bitbucket list-issues

List bitbucket issues

## Description

List bitbucket issues

## Synopsis

`bitbucket list-issues [--site  <site>] --repo  <repo> [--milestone  <milestone>] [--priority  <priority>] [--state  <state>] [--type  <type>] [--assignee  <assignee>] [--reporter  <reporter>] [--sort  <sort>] [--page  <page>]`

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


#### `milestone` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The number of the milestone to associate this issue with.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--milestone  "milestone1"`

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


#### `type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Type of the issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--type  "bug"`
 ,  `--type  "enhancement"`
 ,  `--type  "proposal"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `assignee` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; User to assign to this issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--assignee  "assignee"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `reporter` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The user that reported the issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--reporter  "bitbucket_user1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `sort` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; What to sort results by  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--sort  "created_on"`
 ,  `--sort  "updated_on"`
 ,  `--sort  "priority"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `page` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Page number of the results to fetch  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--page  2`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
!bitbucket list-issues --site "localhost:8080" --repo "simulage/daccosim" --type "proposal"
```
Output: 
```
ID	TYPE    	PRIORITY	STATE  	ASSIGNEE	MILESTONE	TITLE
32	proposal	major   	closed 	-       	null     	what is Matryoshka.jar
 8	proposal	major   	on hold	-       	null     	Participation in FMU Cross-Check
```

