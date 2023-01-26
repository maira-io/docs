# s3 delete-object

Delete Object

## Description

delete object from bucket

## Synopsis

`s3 delete-object [--site  <site>] [--region  <region>] <key> --bucket  <bucket>`

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

&nbsp;&nbsp;&nbsp;&nbsp; The key of the object to be removed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"test2"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `bucket` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The name of the bucket where object is located  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--bucket  "mairatestbucket2"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
x = ! s3 delete-object --bucket "mairatestbucket10" "firstContent" --site "localhost"
```
Output: 
```

```

