# prometheus get-alertmanagers

Get the Prometheus alertmanager discovery

## Description

Returns an overview of the current state of the Prometheus alertmanager discovery. Both the active and dropped Alertmanagers are part of the response

## Synopsis

`prometheus get-alertmanagers [--site  <site>] [--cluster  <cluster>]`

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
! prometheus get-alertmanagers --site "localhost"
```
Output: 
```
{
  "status": "success",
  "data": {
    "activeAlertmanagers": [
      {
        "url": "http://localhost:9093/api/v2/alerts"
      }
    ],
    "droppedAlertmanagers": []
  }
}
```

