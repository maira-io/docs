# cloudwatch set-alarm-state

Temporarily sets the state of an alarm for testing purposes

## Description

temporarily sets the state of an alarm for testing purposes

## Synopsis

`cloudwatch set-alarm-state [--site  <site>] [--region  <region>] <name> --reason  <reason> [--reason_data  <reason_data>] --state  <state>`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `region` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; AWS region for the service  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--region  "region-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The name of the alarm  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `reason` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The reason that this alarm is set to this specific state, in text format  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--reason  "reason-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `reason_data` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The reason that this alarm is set to this specific state, in JSON format  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--reason_data  "reason_data-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `state` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The value of the state  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--state  "state-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

cloudwatch set-alarm-state --region  "region-1" "name-1" --reason  "reason-1" --reason_data  "reason_data-1" --state  "state-1"
