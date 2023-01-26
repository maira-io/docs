# cloudwatch get-metric-stats

Retrieve CloudWatch metric statistics

## Description

retrieve CloudWatch metric statistics

## Synopsis

`cloudwatch get-metric-stats [--site  <site>] [--region  <region>] <name> --start_time  <start_time> --end_time  <end_time> --namespace  <namespace> [--unit  <unit>] [--stat  <stat>] [--extended_stat  <extended_stat>] --period  <period> [--dimensions  <dimensions>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; The name of the metric, with or without spaces  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `start_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; The time stamp indicating the earliest data to be returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--start_time  "60m ago"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `60m ago`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `end_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; The time stamp indicating the latest data to be returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--end_time  "now"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `now`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `namespace` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The namespace of the metric, with or without spaces  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--namespace  "namespace-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `unit` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The unit for a given metric  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--unit  "unit-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `stat` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The unit for a given metric  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--stat  "SampleCount, Average, Sum, Maximum, Minimum"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `extended_stat` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The unit for a given metric  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--extended_stat  "extended_stat-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `period` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; The granularity, in seconds, of the returned data points  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--period  * Start time between 3 hours and 15 days ago - Use a multiple of 60 seconds (1 minute).
* Start time between 15 and 63 days ago - Use a multiple of 300 seconds (5 minutes).
* Start time greater than 63 days ago - Use a multiple of 3600 seconds (1 hour).
`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `dimensions` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The dimensions to filter against  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--dimensions  "{"Class":"None"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

cloudwatch get-metric-stats --region  "region-1" "name-1" --namespace  "namespace-1" --unit  "unit-1" --stat  "SampleCount, Average, Sum, Maximum, Minimum" --extended_stat  "extended_stat-1" --period  * Start time between 3 hours and 15 days ago - Use a multiple of 60 seconds (1 minute).
* Start time between 15 and 63 days ago - Use a multiple of 300 seconds (5 minutes).
* Start time greater than 63 days ago - Use a multiple of 3600 seconds (1 hour).
 --dimensions  "{"Class":"None"}"
