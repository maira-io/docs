# s3 add-object

Add object into bucket

## Description

Add file named "key" in bucket

## Synopsis

`s3 add-object [--site  <site>] [--region  <region>] --key  <key> --bucket  <bucket> <value>`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `region` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; AWS region for the service  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--region  "region-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `key` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Key for the object being added  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--key  "test2"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `bucket` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The name of the bucket where content needs to be added  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--bucket  "mairatestbucket2"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `value` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Value to be added  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"This is test2 file"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
!  s3 add-object "Hello world" --bucket "mairatestbucket10" --key "firstContent" --site "localhost"
```
Output: 
```

```

