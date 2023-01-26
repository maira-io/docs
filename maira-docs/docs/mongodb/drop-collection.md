# mongodb drop-collection

drop mongodb collection

## Description

Delete the collection provided from the server

## Synopsis

`mongodb drop-collection [--site  <site>] <database> <collection>`

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


#### `collection` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the collection  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"collection-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! mongodb drop-collection --site "localhost" "test-database-1" "test-collection-1"
```
Output: 
```
Collection test-collection-1 is dropped
```

