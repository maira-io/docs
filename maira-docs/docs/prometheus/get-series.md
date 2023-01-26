# prometheus get-series

Get time series for given label set

## Description

Returns the list of time series that match a certain label set.

## Synopsis

`prometheus get-series [--site  <site>] [--cluster  <cluster>] [--start_time  <start_time>] [--end_time  <end_time>] <series_selector>`

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


#### `series_selector` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Series selector argument that selects the series to return.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"up"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required, multiple allowed_  



## Examples

Input: 
```
! prometheus get-series --site "localhost" --start_time "2021-10-19T20:10:51.781Z" --end_time "2021-10-20T20:10:51.781Z" "up"
```
Output: 
```
{
  "status": "success",
  "data": [
    {
      "__name__": "up",
      "instance": "localhost:1234",
      "job": "kafka"
    },
    {
      "__name__": "up",
      "instance": "localhost:9090",
      "job": "prometheus"
    },
    {
      "__name__": "up",
      "instance": "localhost:9100",
      "job": "node-exporter"
    }
  ]
}
```

