# mongodb update-document

Update mongodb Document

## Description

Update the document with given ID

## Synopsis

`mongodb update-document [--site  <site>] <database> <collection> <index> <update>`

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


#### `index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Index value to ID the target document  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"index-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `update` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Values to be updated  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"update-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! mongodb update-document --site "localhost" "test-database-1" "test-collection-1" "{\"sample\":\"document\"}" "{\"sample\":\"document\", \"sample2\":4351}"
```
Output: 
```
Matched and replaced an existing document
```

