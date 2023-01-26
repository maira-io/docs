# maira update-resources

Update resources

## Description

Update resources in the DB

## Synopsis

`maira update-resources --namespace  <namespace> --resource_monitor  <resource_monitor> [<resources>]`

## Arguments


#### `namespace` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Namespace on which update resources is executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--namespace  "namespace-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.namespace`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `resource_monitor` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Resource monitor to which the resources are attached to  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--resource_monitor  "resource_monitor-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `resources` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of resources (json list) which  are getting updated  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"resources-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

maira update-resources --resource_monitor  "resource_monitor-1" "resources-1"
