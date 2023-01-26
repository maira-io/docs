# appd get-tiers

Get all tiers for a business application

## Description

Get all tiers for a business application

## Synopsis

`appd get-tiers [--site  <site>] [--cluster  <cluster>] <application> [--tier  <tier>]`

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


#### `tier` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; tier name or ID  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--tier  "tier-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x = ! appd get-business-transactions --site "localhost:8081" "bookinfo" --tier "details"
```
Output: 
```
ID    	NAME       	TYPE              	AGENT-TYPE      	NODES 
939728	details    	C/C++ SDK         	NATIVE_SDK      	    1	
```

