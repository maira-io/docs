# elasticsearch clear-cache

Clear the caches of one or more indices

## Description

Clear the caches of one or more indices

## Synopsis

`elasticsearch clear-cache [--site  <site>] [--cluster  <cluster>] [--target  <target>] [--fielddata ] [--fields  <fields>] [--query ] [--request ]`

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


#### `target` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of data streams, indices, and index aliases used to limit the request  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--target  "students or _all"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `fielddata` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, clears the cache. If fields argument is specified, cache is cleared for specified fields only  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--fielddata  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `fields` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of field names used to limit the fielddata parameter. Defaults to all fields.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--fields  "foo"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `query` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, clears the query cache  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--query  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `request` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, clears the request cache  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--request  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch clear-cache --site "localhost" --cluster "elastic-1" --fielddata --target "class"
```
Output: 
```
{
  "_shards": {
    "total": 4,
    "successful": 2,
    "failed": 0
  }
}
```

