# icinga delete-object

deletes object

## Description

deletes object

## Synopsis

`icinga delete-object [--site  <site>] [--cluster  <cluster>] --name  <name> --object_type  <object_type> [--cascade  <cascade>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the object. Must be unique on a per-host basis  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--name  "checkcommand1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `object_type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Type of object  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--object_type  "hosts, services"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `cascade` - (boolean)

&nbsp;&nbsp;&nbsp;&nbsp; Delete objects depending on the deleted objects (e.g. services on a host).  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cascade  "cascade-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

icinga delete-object --name  "checkcommand1" --object_type  "hosts, services" --cascade  "cascade-1"
