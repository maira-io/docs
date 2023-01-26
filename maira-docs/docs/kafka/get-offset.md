# kafka get-offset

Get offset for Kafka topic partition

## Description

Get offset of a particular partition of kafka topic

## Synopsis

`kafka get-offset [--site  <site>] [--cluster  <cluster>] --topic  <topic> <partition> [--offset_option  <offset_option>] [--time  <time>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--topic  "topic-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `partition` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Partition number  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `offset_option` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Takes one of the 2 values, "OLDEST" or "NEWEST". Provide either offset_option or time  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--offset_option  "offset_option-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `time` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; This gives the offset value at that point in time. Input should follow protobuf timestamp pattern (eg. 2020-09-01T21:46:43Z)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--time  "time-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! kafka get-offset --site "localhost" --topic "test-topic-1" 1 --offset_option "NEWEST"
```
Output: 
```
1
```

