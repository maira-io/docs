# bitbucket search-issues

Search bitbucket issues

## Description

Search bitbucket issues based on the given query

## Synopsis

`bitbucket search-issues [--site  <site>] --repo  <repo> <query> [--sort  <sort>] [--page  <page>]`

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


#### `query` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Query string in bitbucket query format  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"state="new" AND title ~ "text""`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `sort` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; What to sort results by  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--sort  "created_on"`
 ,  `--sort  "updated_on"`
 ,  `--sort  "priotity"`

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
!bitbucket search-issues --site "localhost:8080" --repo "simulage/daccosim" 'title ~ "Matryoshka.jar"'
```
Output: 
```
ID	TYPE    	STATE 	ASSIGNEES	MILESTONE	TITLE
32	proposal	closed	-        	null     	what is Matryoshka.jar
```

