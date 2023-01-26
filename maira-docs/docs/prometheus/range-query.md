# prometheus range-query

Run a prometheus query over a range of time

## Description

Evaluates an expression query over a range of time

## Synopsis

`prometheus range-query [--site  <site>] [--cluster  <cluster>] <query> --start_time  <start_time> --end_time  <end_time> --step  <step>`

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

&nbsp;&nbsp;&nbsp;&nbsp; Query  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"up"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `start_time` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Start timestamp (inclusive, Should be provided in rfc3339 format)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--start_time  "2015-07-01T20:10:51.781Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `end_time` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; End timestamp (inclusive, Should be provided in rfc3339 format)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--end_time  "2015-07-01T20:10:51.781Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `step` - (float)

&nbsp;&nbsp;&nbsp;&nbsp; Query resolution step width in seconds  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--step  15.5`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! prometheus range-query --site "localhost" --start_time "2021-10-19T20:10:51.781Z" --end_time "2021-10-20T20:10:51.781Z" "up" --step 30000
```
Output: 
```
{
  "status": "success",
  "data": {
    "resultType": "matrix",
    "result": [
      {
        "metric": {
          "__name__": "up",
          "instance": "localhost:1234",
          "job": "kafka"
        },
        "values": [
          [1634674251.781, "0"],
          [1634704251.781, "0"],
          [1634734251.781, "0"]
        ]
      },
      {
        "metric": {
          "__name__": "up",
          "instance": "localhost:9090",
          "job": "prometheus"
        },
        "values": [
          [1634674251.781, "1"],
          [1634704251.781, "1"],
          [1634734251.781, "1"]
        ]
      },
      {
        "metric": {
          "__name__": "up",
          "instance": "localhost:9100",
          "job": "node-exporter"
        },
        "values": [
          [1634674251.781, "0"],
          [1634704251.781, "0"],
          [1634734251.781, "0"]
        ]
      }
    ]
  }
}
```

