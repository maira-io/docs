# gcp get-log-metric

Get log metric

## Description

Get log metric for gcp logs

## Synopsis

`gcp get-log-metric [--site  <site>] [<name>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Metric name to fetch metric  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"CassandraMetric"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! gcp get-log-metric --site "localhost:8081" "testMetric1"
```
Output: 
```
NAME            DESCRIPTION     FILTER                                                  TYPE                                    
testMetric1     testDesc2       severity="INFO"                                         logging.googleapis.com/user/testMetric1
```

