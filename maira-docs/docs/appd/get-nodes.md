# appd get-nodes

Get all nodes in a business application

## Description

Get node information for all nodes in a business application

## Synopsis

`appd get-nodes [--site  <site>] [--cluster  <cluster>] <application> [--node  <node>] [--tier  <tier>]`

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


#### `application` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; application name or ID  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"app-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `node` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; node name or ID  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--node  "node-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `tier` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; tier name or ID  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--tier  "tier-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x = ! appd get-nodes --site "localhost:8081" "bookinfo" --tier "ratings-v1-0"
```
Output: 
```
ID    	NAME        	TYPE 	TIER   	MACHINE-NAME	MACHINE-OS	MACHINE-AGT-VERSION	APP-AGT-VERSION                 	AGENT-TYPE       
919391	ratings-v1-0	Other	ratings	ac0015dd9de4	Linux     	                   	21.4.0.0 compatible with 4.4.1.0	NODEJS_APP_AGENT	
919393	ratings-v1-0	Other	ratings	3105e31e1b26	Linux     	                   	21.4.0.0 compatible with 4.4.1.0	NODEJS_APP_AGENT
```

