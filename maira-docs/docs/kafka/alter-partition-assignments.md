# kafka alter-partition-assignments

Alter partition assignments

## Description

define brokers in which replicas of each partition reside, partitions will be reassigned

## Synopsis

`kafka alter-partition-assignments [--site  <site>] [--cluster  <cluster>] --topic  <topic> [<partition_assignment>]`

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


#### `partition_assignment` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Partition assignment  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"partition_assignment-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! kafka alter-partition-assignments --site "localhost" --topic "test-topic-1" "[[1,2],[2,3]]"
```
Output: 
```
Partitions of topic test-topic-1 reassigned
```

