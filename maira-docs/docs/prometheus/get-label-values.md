# prometheus get-label-values

Get label values for a provided label name

## Description

Returns a list of label values for a provided label name

## Synopsis

`prometheus get-label-values [--site  <site>] [--cluster  <cluster>] <label> [--series_selector  <series_selector>] [--start_time  <start_time>] [--end_time  <end_time>]`

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


#### `label` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the label for which values are returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"job"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `series_selector` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Series selector argument that selects the series from which label values are read  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--series_selector  "up"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


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
! prometheus get-label-values --site "localhost" --start_time "2021-10-19T20:10:51.781Z" --end_time "2021-10-20T20:10:51.781Z" --series_selector "up" "instance"
```
Output: 
```
{
  "status": "success",
  "data": ["localhost:1234", "localhost:9090", "localhost:9100"]
}
```

