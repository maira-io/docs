# appd get-events

Get events in an application

## Description

Get events of a specific event type occurred in an application in a specific time period

## Synopsis

`appd get-events [--site  <site>] [--cluster  <cluster>] <application> [--type  <type>] [--severity  <severity>] [--start_time  <start_time>] [--end_time  <end_time>]`

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


#### `application` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; application name or ID  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"app-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Type of the events.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--type  "CUSTOM"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `APPLICATION_CONFIG_CHANGE`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `severity` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Provides the severity levels.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--severity  "INFO"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `INFO`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `start_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; time from which to get metrics data  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--start_time  "15m ago"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `15m ago`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `end_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; time to which to get metrics data  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--end_time  "now"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `now`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

appd get-events "app-1"
