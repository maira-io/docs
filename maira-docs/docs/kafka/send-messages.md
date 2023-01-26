# kafka send-messages

Send messages to kafka topic

## Description

Send Messages to topic. Can be sent either synchronously or asynchronously

## Synopsis

`kafka send-messages [--site  <site>] [--cluster  <cluster>] --topic  <topic> [--asynchronous ] [--messages  <messages>]`

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



## Examples

Input: 
```
! kafka send-messages --site "localhost" --asynchronous --topic "test-topic-1" `{"key":"1","value":"hi"}` `{"key":"2","value":"hello"}`
```
Output: 
```
[
  {
    "topic_name":"test-topic-1",
    "offset":3,
    "time":{
      "seconds":1653551998,
      "nanos":115288839
    }
  },
  {
    "topic_name":"test-topic-1",
    "offset":4,
    "time":{
      "seconds":1653551998,
      "nanos":115590761
    }
  }
]
```

