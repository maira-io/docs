# appd delete-collector

Delete a collector

## Description

Delete a collector with configuration ID

## Synopsis

`appd delete-collector [--site  <site>] [--cluster  <cluster>] --id  <id>`

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


#### `id` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; configuration-id for deleting the collector  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--id  562`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required, multiple allowed_  



## Examples

appd delete-collector --id  562
