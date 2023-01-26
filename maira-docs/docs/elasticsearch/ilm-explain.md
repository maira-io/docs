# elasticsearch ilm-explain

Explain index lifecycle management

## Description

Retrieve the current lifecycle status for one or more indices

## Synopsis

`elasticsearch ilm-explain [--site  <site>] [--cluster  <cluster>] [--target  <target>] [--only_errors ]`

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


#### `only_errors` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, Filters the returned indices to only indices that are managed by ILM and are in an error state  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--only_errors  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch ilm-explain --site "localhost" --cluster "elastic-1" --only_errors
```
Output: 
```
{
  "indices": {}
}
```

