# s3 list-objects

List all objects of S3 bucket in the argument

## Description

List all objects of S3 bucket in the argument

## Synopsis

`s3 list-objects [--site  <site>] [--region  <region>] <bucket> [--prefix  <prefix>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; S3 bucket from which object keys are to be listed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"bucket-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `prefix` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Keys which start with the prefix argument  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--prefix  "maira"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! s3 list-objects --site "localhost:8081" "test1" --prefix "m"
```
Output: 
```
CREATION_TIME                   SIZE                   NAME
2022-09-20 05:31:12 +0000 UTC      0              maira123/
2022-09-20 16:21:35 +0000 UTC     16    maira123/myfile.txt
2022-09-20 16:17:11 +0000 UTC      0         maira123/test/
2022-09-20 05:31:24 +0000 UTC      0              maira456/
```

