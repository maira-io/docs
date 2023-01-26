# github search-issues

Search github issues

## Description

Search github issues based on the given query

## Synopsis

`github search-issues [--site  <site>] --repo  <repo> <query> [--sort  <sort>] [--direction  <direction>] [--per_page  <per_page>] [--page  <page>]`

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


#### `query` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Query string in github query format  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"warning in:title"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


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
!github search-issues --site "localhost:8081" --repo "mozilla/protocol" "Deprecate in:title state:open"
```
Output: 
```
NUMBER	LABELS        	STATE	ASSIGNEES	MILESTONE	TITLE
  787	Bug :beetle:  	open 	-        	null     	Deprecate (or update) the Callout component
  694	Infra :wrench:	open 	-        	null     	Netlify "Trusty" build image is deprecated
```

