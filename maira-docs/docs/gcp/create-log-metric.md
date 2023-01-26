# gcp create-log-metric

Create log metric

## Description

create log metric for gcp logs

## Synopsis

`gcp create-log-metric [--site  <site>] [<name>] [--description  <description>] [--filter  <filter>] [--disabled ] [--log_name  <log_name>] [--severity  <severity>] [--filter  <filter>] [--start_time  <start_time>] [--end_time  <end_time>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Metric name to create metric  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"CassandraMetric"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `description` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Description for the metric  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--description  "Metric details for Cassandra"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `filter` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A filter that chooses which log entries to create metric for  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--filter  "resource.type=gae_app AND severity>=ERROR"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `disabled` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set to True, then this metric is disabled and it does not generate any points  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--disabled  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `log_name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; log names to filter out the logs  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--log_name  "projects/macro-context-293714/logs/GCEGuestAgent"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `severity` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Severity of the logs  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--severity  "INFO"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `filter` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A filter that chooses which log entries to return  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--filter  "resource.type=gae_app AND severity>=ERROR"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `start_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; The time stamp indicating the earliest data to be returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--start_time  "60m ago"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `end_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; The time stamp indicating the latest data to be returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--end_time  "now"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! gcp create-log-metric --site "localhost:8081" "testMetric5" --description "testDesc" --filter "protoPayload.status >= 400 AND sample(insertId, 0.1)"
```
Output: 
```
NAME            DESCRIPTION     FILTER 
testMetric2     testDesc        severity="INFO"
```

