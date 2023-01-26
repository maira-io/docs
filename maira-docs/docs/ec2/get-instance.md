# ec2 get-instance

Get instance

## Description

get aws ec2 instance

## Synopsis

`ec2 get-instance [--site  <site>] [--region  <region>] <instance_id> --region  <region>`

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


#### `instance_id` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Instance ID  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"instance_id-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `region` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Region for this instance  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--region  "region-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

ec2 get-instance --region  "region-1" "instance_id-1" --region  "region-1"
