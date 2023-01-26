# mongodb get-document

Get mongodb document

## Description

Get the document with given ID in given collection

## Synopsis

`mongodb get-document [--site  <site>] <database> <collection> <Index>`

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


#### `Index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Index of the document  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"Index-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! mongodb get-document --site "localhost" "test-database-1" "test-collection-1" "{\"sample\":\"document\"}"
```
Output: 
```
[{_id ObjectID(\"621785023620990be3cc378d\")} {sample document} {sample2 1}]
```

