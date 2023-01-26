# prometheus get-wal-replay-stats

Get statistics for wal replay

## Description

Returns statistics for Write-Ahead Log (wal) replay

## Synopsis

`prometheus get-wal-replay-stats [--site  <site>] [--cluster  <cluster>]`

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



## Examples

Input: 
```
! prometheus get-wal-replay-stats --site "localhost"
```
Output: 
```
{
  "status": "success",
  "data": {
    "min": 444,
    "max": 449,
    "current": 449
  }
}
```

