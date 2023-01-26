# prometheus get-labels

Get list of label names

## Description

Returns a list of label names

## Synopsis

`prometheus get-labels [--site  <site>] [--cluster  <cluster>] [<series_selector>] [--start_time  <start_time>] [--end_time  <end_time>]`

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


#### `series_selector` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Series selector argument that selects the series from which labels are read  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"up"`

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
! prometheus get-labels --site "localhost" --start_time "2021-10-19T20:10:51.781Z" --end_time "2021-10-20T20:10:51.781Z" "up"
```
Output: 
```
{
  "status": "success",
  "data": ["__name__", "instance", "job"]
}
```

