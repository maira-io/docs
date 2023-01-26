# elasticsearch raw

Run an elasticsearch raw command based REST API

## Description

Run an elasticsearch raw command based REST API, providing path, body and http method

## Synopsis

`elasticsearch raw [--site  <site>] [--cluster  <cluster>] --path  <path> [--method  <method>] [--body  <body>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of elastic search cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "elastic-default"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `elastic-default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `path` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; REST API URL path  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--path  "/my-index-01/_rollover"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `method` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; HTTP method - GET/POST/PUT/DELETE  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--method  "method-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `body` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Body for REST API, JSON formatted  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--body  "{"field1": "value1"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `{}`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

elasticsearch raw --path  "/my-index-01/_rollover" --method  "method-1"
