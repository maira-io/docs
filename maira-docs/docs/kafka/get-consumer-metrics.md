# kafka get-consumer-metrics

Get kafka consumer metrics

## Description

Get metrics for one or more consumer for a kafka consumer

## Synopsis

`kafka get-consumer-metrics <metric> <consumer>`

## Arguments


#### `metric` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Specify metric as string or Id  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"metric-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `consumer` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Specify consumer as string or Id  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"consumer-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

kafka get-consumer-metrics "metric-1" "consumer-1"
