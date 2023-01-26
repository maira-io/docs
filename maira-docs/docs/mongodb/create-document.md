# mongodb create-document

Create mongodb document

## Description

Create the document in the provided collection

## Synopsis

`mongodb create-document [--site  <site>] <database> <collection> <document>`

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


#### `document` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Document to be created in specified Collection  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"document-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! mongodb create-document --site "localhost" "test-database-1" "test-collection-1" "{\"sample\":\"document\", \"sample2\":1}"
```
Output: 
```
{DocumentId: testcollection1}
```

