# icinga modify-check-command

modifies a check command object

## Description

modifies a check command object

## Synopsis

`icinga modify-check-command [--site  <site>] [--cluster  <cluster>] --name  <name> --commands  <commands> [--arguments  <arguments>] [--vars  <vars>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the check command object. Must be unique on a per-host basis  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--name  "checkcommand1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `commands` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The command. This can either be an array of individual command arguments. Alternatively a string can be specified in which case the shell interpreter (usually /bin/sh) takes care of parsing the command.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--commands  "/usr/local/sbin/check_http"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required, multiple allowed_  


#### `arguments` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A dictionary of command arguments  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--arguments  "{ "-I": "$mytest_iparam$" }"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `vars` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A dictionary containing custom variables that are specific to this command.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--vars  "{http_ssl: false, http_sni: false}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! icinga modify-check-command --site "localhost" --name "ccmd1" --commands "/usr/local/sbin/check_https"
```
Output: 
```
{
  "results": [
    {
      "code": 200,
      "name": "ccmd1",
      "status": "Attributes updated.",
      "type": "CheckCommand"
    }
  ]
}
```

