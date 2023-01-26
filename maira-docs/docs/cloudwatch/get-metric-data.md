# cloudwatch get-metric-data

Retrieve CloudWatch metric values

## Description

retrieve CloudWatch metric values

## Synopsis

`cloudwatch get-metric-data [--site  <site>] [--region  <region>] --start_time  <start_time> --end_time  <end_time> --query  <query>`

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


#### `start_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; The time stamp indicating the earliest data to be returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--start_time  "2016-03-11T03:45:40Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `now-60m`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `end_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; The time stamp indicating the latest data to be returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--end_time  "now"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `now`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `query` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The metric queries to be returned. A single GetMetricData call can include as many as 500 MetricDataQuery structures  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--query  "query-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

cloudwatch get-metric-data --region  "region-1" --query  "query-1"
