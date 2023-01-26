# mongodb list-collections

list mongodb collections

## Description

list the collections in the provided database

## Synopsis

`mongodb list-collections [--site  <site>] <database>`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `database` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the database  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"database-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! mongodb list-collections --site "localhost" "test-database-1"
```
Output: 
```
["test-collection-1", "test-collection-2", "test-collection-3"]
```

