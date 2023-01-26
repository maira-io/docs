# s3 list-buckets

List all S3 buckets

## Description

List all S3 buckets

## Synopsis

`s3 list-buckets [--site  <site>] [--region  <region>]`

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



## Examples

Input: 
```
! s3 list-buckets --site "localhost"
```
Output: 
```
CREATION_TIME                          NAME
2022-09-19 10:32:34 +0000 UTC   beehyvpavan
2022-09-19 08:26:38 +0000 UTC   pavantest13
```

