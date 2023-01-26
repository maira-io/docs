# prometheus get-runtime-info

Get various runtime information

## Description

Returns various runtime information properties about the Prometheus server

## Synopsis

`prometheus get-runtime-info [--site  <site>] [--cluster  <cluster>]`

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
! prometheus get-runtime-info --site "localhost"
```
Output: 
```
{
  "status": "success",
  "data": {
    "startTime": "2021-10-13T12:05:08.252857048Z",
    "CWD": "/",
    "reloadConfigSuccess": true,
    "lastConfigTime": "2021-10-13T12:05:09Z",
    "corruptionCount": 0,
    "goroutineCount": 31,
    "GOMAXPROCS": 4,
    "GOGC": "",
    "GODEBUG": "",
    "storageRetention": "15d"
  }
}
```

