# icinga create-service

Creates a new service object

## Description

Creates a new service object

## Synopsis

`icinga create-service [--site  <site>] [--cluster  <cluster>] --name  <name> [--templates  <templates>] [--ignore_on_err ] --host_name  <host_name> [--groups  <groups>] [--vars  <vars>] --check_command  <check_command>`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of Icinga cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "icinga-default"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `icinga-default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the service object. Must be unique on a per-host basis  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--name  "service1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `templates` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Import existing configuration templates for this object type. These templates must either be statically configured or provided in config packages  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--templates  "generic-service"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `ignore_on_err` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, Ignore object creation errors and return an HTTP 200 status instead.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--ignore_on_err  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `host_name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The host this service belongs to. There must be a Host object with that name.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--host_name  "host1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `groups` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A list of service groups this service belongs to  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--groups  "group1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `vars` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A dictionary containing custom variables that are specific to this service.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--vars  "{"snmp_community":"public", "snmp_oid":"DISMAN-EVENT-MIB::sysUpTimeInstance"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `check_command` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The name of the check command.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--check_command  "snmp"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! icinga create-service --site "localhost" --name "service1" --host_name "host1" --check_command "snmp" --vars `{"snmp_community":"public", "snmp_oid":"DISMAN-EVENT-MIB::sysUpTimeInstance"}`
```
Output: 
```
{
  "results": [
    {
      "code": 200,
      "status": "Object was created"
    }
  ]
}
```

