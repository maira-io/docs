# appd get-dbagent-events

Retrieve DB agent events

## Description

Get all Database Agent Events

## Synopsis

`appd get-dbagent-events [--site  <site>] [--cluster  <cluster>] [--start_time  <start_time>] [--end_time  <end_time>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of elastic search cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "appdynamics-default"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `appdynamics-default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `start_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; time from which to get events data  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--start_time  "30m ago"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `30m ago`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `end_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; time to which to get events data  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--end_time  "now"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `now`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

appd get-dbagent-events
