# icinga create-check-command

Creates a new check command object

## Description

Creates a new check command object

## Synopsis

`icinga create-check-command [--site  <site>] [--cluster  <cluster>] --name  <name> [--templates  <templates>] [--ignore_on_err ] --commands  <commands> [--arguments  <arguments>] [--vars  <vars>]`

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


#### `templates` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Import existing configuration templates for this object type. These templates must either be statically configured or provided in config packages  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--templates  "plugin-check-command"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `ignore_on_err` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, Ignore object creation errors and return an HTTP 200 status instead.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--ignore_on_err  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


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
! icinga create-check-command --site "localhost" --name "ccmd1" --commands "/usr/local/sbin/check_http"
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

