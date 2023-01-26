# bitbucket get-issue

Get bitbucket issue

## Description

Get bitbucket issue

## Synopsis

`bitbucket get-issue [--site  <site>] --repo  <repo> --issue  <issue>`

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



## Examples

Input: 
```
!bitbucket get-issue --site "localhost:8080" --repo "sbansal7/repo1" --issue 1
```
Output: 
```
ID	TYPE	PRIORITY	STATE	ASSIGNEE     	MILESTONE	TITLE
 1	bug 	major   	new  	Sachin Bansal	null     	test issue
```

