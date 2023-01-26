# maira list-resources

List resources

## Description

List resources in the DB

## Synopsis

`maira list-resources --namespace  <namespace> --resource_monitor  <resource_monitor>`

## Arguments


#### `namespace` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Namespace on which list resources is executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--namespace  "namespace-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.namespace`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `resource_monitor` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Resource monitor to which the resources are attached to  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--resource_monitor  "resource_monitor-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

maira list-resources --resource_monitor  "resource_monitor-1"
