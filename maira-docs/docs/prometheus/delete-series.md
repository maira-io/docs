# prometheus delete-series

Delete data for a selection of series in a time range

## Description

Delete data for a selection of series in a time range. The actual data still exists on disk and is cleaned up in future compactions or can be explicitly cleaned up by hitting the Clean Tombstones endpoint.

## Synopsis

`prometheus delete-series [--site  <site>] [--cluster  <cluster>] [--start_time  <start_time>] [--end_time  <end_time>] <series_selector>`

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

&nbsp;&nbsp;&nbsp;&nbsp; label matcher argument that selects the series to delete. At least one series_selector argument must be provided.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"up"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required, multiple allowed_  



## Examples

Input: 
```
! prometheus delete-series --site "localhost" "up"
```
Output: 
```

```

