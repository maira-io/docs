# kafka create-consumer

Create kafka consumer

## Description



## Synopsis

`kafka create-consumer [--site  <site>] [--cluster  <cluster>] <consuming_topic> [--all_partitions ] [--partitions  <partitions>] [--offset_option  <offset_option>] [--offset_value  <offset_value>]`

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


#### `consuming_topic` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Topic to be consumed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"consuming_topic-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `all_partitions` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; boolean value to decide whether consumer should consume from all partitions  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--all_partitions  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `partitions` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; partition number from which consumer should consume. Multiple partitions can be consumed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--partitions  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `offset_option` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Value can be either "OLDEST" or "NEWEST"  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--offset_option  "OLDEST"`
 ,  `--offset_option  "NEWEST"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `offset_value` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; offset from which messages are read in the provided partitions. Provide either offset_option or offset_value  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--offset_value  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! kafka create-consumer --site "localhost" "test-topic-1" --partitions 1 --offset_option "NEWEST"
```
Output: 
```
Consumer for the topic "test-topic-1" created
```

