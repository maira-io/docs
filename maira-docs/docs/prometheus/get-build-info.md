# prometheus get-build-info

Get build information properties

## Description

Returns various build information properties about the Prometheus server

## Synopsis

`prometheus get-build-info [--site  <site>] [--cluster  <cluster>]`

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
! prometheus get-build-info --site "localhost"
```
Output: 
```
{
  "status": "success",
  "data": {
    "version": "2.26.0",
    "revision": "3cafc58827d1ebd1a67749f88be4218f0bab3d8d",
    "branch": "HEAD",
    "buildUser": "root@a67cafebe6d0",
    "buildDate": "20210331-11:56:23",
    "goVersion": "go1.16.2"
  }
}
```

