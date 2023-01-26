# kafka create-topic

create Kafka topic

## Description

Create topic with the provided name & config

## Synopsis

`kafka create-topic [--site  <site>] [--cluster  <cluster>] <topic> [--validate_only ] [--replication_factor  <replication_factor>] [--num_partitions  <num_partitions>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the topic  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"topic-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `validate_only` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Only check if a topic can be created with provided config and not create topic in cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--validate_only  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `replication_factor` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Replication factor  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--replication_factor  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `num_partitions` - (int)

&nbsp;&nbsp;&nbsp;&nbsp;   

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--num_partitions  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! kafka create-topic --site "localhost" --cluster "default" "test-topic-1" --replication_factor 2 --num_partitions 3
```
Output: 
```
Topic "test-topic-1" is created
```

