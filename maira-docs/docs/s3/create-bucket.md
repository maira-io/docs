# s3 create-bucket

Create new bucket

## Description

Create new S3 bucket

## Synopsis

`s3 create-bucket [--site  <site>] [--region  <region>] <bucket> --access  <access> --zone  <zone>`

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


#### `bucket` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; name of the bucket to be created  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"mairabucket7"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `access` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The ACL access of the bucket  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--access  "private"`
 ,  `--access  "public-read"`
 ,  `--access  "public-read-write"`
 ,  `--access  "authenticated-read"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `zone` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Zone in which, the bucket needs to be created  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--zone  "ap-south-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! s3 create-bucket --site "localhost" "mairatestbucket10" --zone "ap-south-1" --access "private"
```
Output: 
```
{
  "Location": "http://mairatestbucket11.s3.amazonaws.com/",
  "ResultMetadata": {}
}
```

