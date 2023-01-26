# s3 add-content

Add content into bucket

## Description

Add file named "key" in bucket

## Synopsis

`s3 add-content [--site  <site>] --key  <key> --bucket  <bucket> <content>`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
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


#### `content` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Content to be added  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"This is test2 file"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
!  s3 add-content "Hello world" --bucket "mairatestbucket10" --key "firstContent" --site "localhost"
```
Output: 
```
{
  "BucketKeyEnabled": false,
  "ETag": "\"3e25960a79dbc69b674cd4ec67a72c62\"",
  "Expiration": null,
  "RequestCharged": "",
  "SSECustomerAlgorithm": null,
  "SSECustomerKeyMD5": null,
  "SSEKMSEncryptionContext": null,
  "SSEKMSKeyId": null,
  "ServerSideEncryption": "",
  "VersionId": null,
  "ResultMetadata": {}
}
```

