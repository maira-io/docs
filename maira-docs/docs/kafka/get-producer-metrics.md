# kafka get-producer-metrics

Get kafka producer metrics

## Description

Get metrics of a producer for a given topic or all topics combined

## Synopsis

`kafka get-producer-metrics [--site  <site>] [--cluster  <cluster>] [<topic>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of kafka cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "cluster-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `topic` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of kafka topic  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"topic-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

kafka get-producer-metrics "topic-1"
