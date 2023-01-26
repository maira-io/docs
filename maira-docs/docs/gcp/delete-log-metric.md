# gcp delete-log-metric

Update log metric

## Description

update log metric for gcp logs

## Synopsis

`gcp delete-log-metric [--site  <site>] [<name>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Metric name to delete metric  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"CassandraMetric"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

gcp delete-log-metric "CassandraMetric"
