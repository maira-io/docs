# icinga modify-host

modifies a host object

## Description

modifies a host object

## Synopsis

`icinga modify-host [--site  <site>] [--cluster  <cluster>] --name  <name> [--address  <address>] [--groups  <groups>] [--vars  <vars>] [--check_command  <check_command>]`

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
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! icinga modify-host --site "localhost" --name "host1" --check_command "hostalive"  --vars `{"os":"unix"}`
```
Output: 
```
{
  "results": [
    {
      "code": 200,
      "name": "host1",
      "status": "Attributes updated.",
      "type": "Host"
    }
  ]
}
```

