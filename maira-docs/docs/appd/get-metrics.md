# appd get-metrics

Get metric hierarchy

## Description

Get metric hierarchy

## Synopsis

`appd get-metrics [--site  <site>] [--cluster  <cluster>] <application> [<metric_path>]`

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


#### `metric_path` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; path to the metric in the metric hierarchy.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"Overall Application Performance"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x = ! appd get-metrics --site "localhost:8081" "bookinfo"
```
Output: 
```
TYPE  	NAME                                   
folder	Overall Application Performance       	
folder	Business Transaction Performance      	
folder	Information Points                    	
folder	Application Infrastructure Performance	
folder	Errors                                	
folder	End User Experience                   	
folder	Mobile                                	
folder	Backends                              	
folder	Service Endpoints                     	
```

