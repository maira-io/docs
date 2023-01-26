# kafka get-topic-config

Get Config for kafka Topic Resource

## Description

Get one or more config for a given kafka topic resource

## Synopsis

`kafka get-topic-config [--site  <site>] [--cluster  <cluster>] --topic  <topic> [<config_names>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; topic name  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--topic  "topic-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `config_names` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of config. If not provided, returns all the config of the resource  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"config_names-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

kafka get-topic-config --topic  "topic-1" "config_names-1"
