# curl

CURL a given URL

## Description

Use CURL command to get a given URL

## Synopsis

`curl [--site  <site>] [--request  |-X  <request>] [--data  |-d  <data>] [--header  |-H  <header>] [--ssl ] [--insecure ] <url>`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site to run the command  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `request` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Request command to use  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--request  "POST"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `data` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; HTTP POST data  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--data  "data-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `header` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Custom headers  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--header  "Host: www.maira.io"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `ssl` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Use SSL  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--ssl  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `insecure` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Allow insecure server connections when using SSL  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--insecure  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `url` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; URL to get  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"www.maira.io"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

curl --request  "POST" --data  "data-1" --header  "Host: www.maira.io" --ssl   --insecure   "www.maira.io"
