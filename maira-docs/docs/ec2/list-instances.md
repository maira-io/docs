# ec2 list-instances

List ec2 instances

## Description

List ec2 instances

## Synopsis

`ec2 list-instances [--site  <site>] [--region  <region>] [--tags  <tags>] [--availability_zone  <availability_zone>] [--owner  <owner>]`

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


#### `tags` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Tags assigned to the instance. Just specify the key to filter all resources with that tag key. Specity key=[value1,value2,...] to filter by acceptable values  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--tags  "creator_id=[user1,user2]"`
 ,  `--tags  "prod_env"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `availability_zone` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Availability Zone of the instance  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--availability_zone  "us-west-1c"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `owner` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; ID of Owner of the instance  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--owner  "248495481387"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

ec2 list-instances --region  "region-1" --tags  "creator_id=[user1,user2]" --availability_zone  "us-west-1c" --owner  "248495481387"
