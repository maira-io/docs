# ec2 resize-autoscaling-group

Resize autoscaling group

## Description

Resize autoscaling group

## Synopsis

`ec2 resize-autoscaling-group [--site  <site>] [--region  <region>] <name> --region  <region> [--min  <min>] [--max  <max>] [--desired_size  <desired_size>]`

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


#### `name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the autoscaling group  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `region` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Region of the autoscaling group to be resized  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--region  "region-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `min` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; minimum size of autoscaling group  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--min  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `max` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; maximum size of autoscaling group  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--max  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `desired_size` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; desired size of autoscaling group  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--desired_size  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

ec2 resize-autoscaling-group --region  "region-1" "name-1" --region  "region-1" --min  1 --max  1 --desired_size  1
