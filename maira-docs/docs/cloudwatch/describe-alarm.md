# cloudwatch describe-alarm

Describe cloudwatch alarms

## Description

show details about cloudwatch alarms

## Synopsis

`cloudwatch describe-alarm [--site  <site>] [--region  <region>] [<names>] [--action_prefix  <action_prefix>] [--name_prefix  <name_prefix>] [--state  <state>]`

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


#### `names` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Specific alarm names to look for  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"names-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `action_prefix` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; To only filter those alarms that use a certain alarm action  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--action_prefix  "action_prefix-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `name_prefix` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Prefix for the alarm name  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--name_prefix  "name_prefix-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `state` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Filter alarms that are currently in the specified state  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--state  "state-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! cloudwatch describe-alarm --site "localhost:8081"
```
Output: 
```
NAME                                                                            STATE   THRESHOLD 
awsec2-i-0bc9464a6e9ba87b5-GreaterThanOrEqualToThreshold-CPUUtilization         ALARM       0.001
awsec2-i-0bc9464a6e9ba87b5-GreaterThanOrEqualToThreshold-CPUUtilization2        ALARM      0.0001
awsec2-i-0bc9464a6e9ba87b5-GreaterThanOrEqualToThreshold-CPUUtilization4        ALARM       0.001
awsec2-i-0bc9464a6e9ba87b5-GreaterThanOrEqualToThreshold-CPUUtilization5        ALARM       0.001
awsec2-i-0bc9464a6e9ba87b5-GreaterThanOrEqualToThreshold-DiskReadOps            OK           0.99
awsec2-i-0bc9464a6e9ba87b5-GreaterThanOrEqualToThreshold-DiskWriteBytes         OK           0.99
```

