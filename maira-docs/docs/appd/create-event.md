# appd create-event

Create events for an application

## Description

Create event for an application

## Synopsis

`appd create-event [--site  <site>] [--cluster  <cluster>] <application> [--summary  <summary>] [--type  <type>] [--severity  <severity>] [--propertynames  <propertynames>] [--propertyvalues  <propertyvalues>]`

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


#### `summary` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Provides the summary for the event.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--summary  "sample summary"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `This is summary of a sample event when the input is not provided`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Enter the event type  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--type  "APPLICATION_ERROR"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `CUSTOM`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `severity` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Provides the severity level. Allowec values - INFO, WARN, ERROR  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--severity  "INFO"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `INFO`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `propertynames` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Input property names  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--propertynames  "key1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `propertyvalues` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Input property values  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--propertyvalues  "value1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

appd create-event "app-1" --propertynames  "key1" --propertyvalues  "value1"
