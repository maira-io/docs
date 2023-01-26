# prometheus query

Run a prometheus query at a single point in time

## Description

Evaluates an instant query at a single point in time

## Synopsis

`prometheus query [--site  <site>] [--cluster  <cluster>] <query> [--time  <time>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Query expression in promQL format  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"up"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; Evaluation timestamp. Query is evaluated at this time (Should be provided in rfc3339 format)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--time  "2015-07-01T20:10:51.781Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! prometheus query --site "localhost" "up"
```
Output: 
```
{
  "status": "success",
  "data": {
    "resultType": "vector",
    "result": [
      {
        "metric": {
          "__name__": "up",
          "instance": "localhost:1234",
          "job": "kafka"
        },
        "value": [1634749948.86, "0"]
      },
      {
        "metric": {
          "__name__": "up",
          "instance": "localhost:9090",
          "job": "prometheus"
        },
        "value": [1634749948.86, "1"]
      },
      {
        "metric": {
          "__name__": "up",
          "instance": "localhost:9100",
          "job": "node-exporter"
        },
        "value": [1634749948.86, "0"]
      }
    ]
  }
}
```

