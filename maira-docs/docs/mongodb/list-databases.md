# mongodb list-databases

List mongodb databases

## Description

List the databases present in the server

## Synopsis

`mongodb list-databases [--site  <site>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! mongodb list-databases --site "localhost"
```
Output: 
```
["test-database-1", "test-database-2", "test-database-3"]
```

