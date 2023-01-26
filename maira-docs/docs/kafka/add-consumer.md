# kafka add-consumer

Add kafka consumer

## Description

Add a kafka consumer for a given topic

## Synopsis

`kafka add-consumer <consumer> <topic>`

## Arguments


#### `consumer` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the consumer to be added  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"consumer-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `topic` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the topic for this consumer  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"topic-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

kafka add-consumer "consumer-1" "topic-1"
