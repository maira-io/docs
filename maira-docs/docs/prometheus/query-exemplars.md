# prometheus query-exemplars

Get exemplars for given PromQL query and time range

## Description

Returns a list of exemplars for a valid PromQL query for a specific time range

## Synopsis

`prometheus query-exemplars [--site  <site>] [--cluster  <cluster>] <query> [--start_time  <start_time>] [--end_time  <end_time>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of prometheus cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "prometheus-default"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `prometheus-default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `query` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Valid PromQL query  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"test_exemplar_metric_total"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `start_time` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Start timestamp (inclusive, Should be provided in rfc3339 format)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--start_time  "2015-07-01T20:10:51.781Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `end_time` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; End timestamp (inclusive, Should be provided in rfc3339 format)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--end_time  "2015-07-01T20:10:51.781Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! prometheus query-exemplars --site "localhost" --start_time "2021-10-19T20:10:51.781Z" --end_time "" "test_exemplar_metric_total"
```
Output: 
```
{
  "status": "success",
  "data": null
}
```

