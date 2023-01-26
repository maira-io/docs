# s3 add-file

Uploads file into bucket

## Description

Uploads file from local system to bucket

## Synopsis

`s3 add-file [--site  <site>] --targetpath  <targetpath> --bucket  <bucket> <filepath>`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `targetpath` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The path where the file needs to be uploaded into bucket. This path is relative to bucket and should be unique for each file upload
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--targetpath  "testfolder/test2.txt"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `bucket` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the bucket where file has to be uploaded  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--bucket  "mairatestbucket2"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `filepath` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Filepath of the object in localsystem  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"/home/beehyv/Desktop/test2.txt"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! s3 add-file --bucket "mairatestbucket10" --site "localhost" "/home/beehyv/Desktop/test1.txt" --targetpath "logs/folder1/test1"
```
Output: 
```
{
  "BucketKeyEnabled": false,
  "ETag": "\"37bd3cc9edb8c9069c528fc9710a23eb\"",
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

