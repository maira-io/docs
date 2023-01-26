# icinga modify-service

modifies a service object

## Description

modifies a service object

## Synopsis

`icinga modify-service [--site  <site>] [--cluster  <cluster>] --name  <name> --host_name  <host_name> [--groups  <groups>] [--vars  <vars>] [--check_command  <check_command>]`

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
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! icinga modify-service --site "localhost" --name "service1" --host_name "host1" --check_command "hostalive" --vars `{"snmp_community":"private"}`
```
Output: 
```
{
  "results": [
    {
      "code": 200,
      "name": "host1!service1",
      "status": "Attributes updated.",
      "type": "Service"
    }
  ]
}
```

