# kafka create-partitions

Create partitions for kafka topic

## Description

Add additional partitions for a topic

## Synopsis

`kafka create-partitions [--site  <site>] [--cluster  <cluster>] <topic> [--total_partition_count  <total_partition_count>] [--assignment  <assignment>] [--validate_only ]`

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
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `total_partition_count` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; this count will be the sum of existing partitions and partitions to be created  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--total_partition_count  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `assignment` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Stringified value of replica assignment array. Number of replica assignments should be equal to number of new partitions  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--assignment  "assignment-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `validate_only` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Only validate whether partition can be created with provided values  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--validate_only  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! kafka create-partitions --site "localhost" "test-topic-1" --total_partition_count 5 --assignment ‘[[1,2],[2,3]]’
```
Output: 
```
Partitions for the topic test-topic-1 created
```

