# kafka send-receive-messages

Send messages to kafka topic and read with a comsumer

## Description

Send Messages to topic and read them with a consumer.

## Synopsis

`kafka send-receive-messages [--site  <site>] [--cluster  <cluster>] --topic  <topic> [--asynchronous ] [--messages  <messages>] [--all_partitions ] [--partitions  <partitions>] [--offset_option  <offset_option>] [--offset_value  <offset_value>]`

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


#### `asynchronous` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Send messages asynchronously  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--asynchronous  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `messages` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Message to be sent  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--messages  "messages-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


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
! kafka send-receive-messages --site "localhost" --asynchronous --topic "test-topic-1" `{"key":"1","value":"hi"}, {"key":"2","value":"hello"}` --partitions 1 --offset_option "OLDEST"
```
Output: 
```
[
  {
    "topic_name":"test-topic-1",
    "partition":1,
    "messages":[
      {
        "key":"1",
        "value":"hi"
      }
    ],
    "time":{
        "seconds":1653550892,
        "nanos":468828435
      }
  }
]
```

