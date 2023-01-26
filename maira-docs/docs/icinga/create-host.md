# icinga create-host

Creates a new host object

## Description

Creates a new host object

## Synopsis

`icinga create-host [--site  <site>] [--cluster  <cluster>] --name  <name> [--templates  <templates>] [--ignore_on_err ] [--address  <address>] [--groups  <groups>] [--vars  <vars>] --check_command  <check_command>`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the host object  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--name  "example.localdomain"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `templates` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Import existing configuration templates for this object type. These templates must either be statically configured or provided in config packages  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--templates  "generic-host"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `ignore_on_err` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, Ignore object creation errors and return an HTTP 200 status instead.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--ignore_on_err  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `address` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The host’s IPv4 address. Available as command runtime macro $address$ if set.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--address  "192.168.1.1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `groups` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A list of host groups this host belongs to  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--groups  "group1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `vars` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A dictionary containing custom variables that are specific to this host.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--vars  "{"os":"linux"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `check_command` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The name of the check command.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--check_command  "hostalive"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! icinga create-host --site "localhost" --cluster "icinga-default" --name "host1" --templates "generic-host" --ignore_on_err --address "192.168.5.96" --check_command "hostalive" --vars `{"os":"linux"}`
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

