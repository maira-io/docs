# s3 delete-content

Delete Object

## Description

delete object from bucket

## Synopsis

`s3 delete-content [--site  <site>] <key> --bucket  <bucket>`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
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
x = ! s3 delete-content --bucket "mairatestbucket10" "firstContent" --site "localhost"
```
Output: 
```
{
  "DeleteMarker": false,
  "RequestCharged": "",
  "VersionId": null,
  "ResultMetadata": {}
}
```

