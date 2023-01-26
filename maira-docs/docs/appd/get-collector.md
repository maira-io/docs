# appd get-collector

Retrieve collector

## Description

Returns collector if ID is provided. Else, returns all collectors

## Synopsis

`appd get-collector [--site  <site>] [--cluster  <cluster>] [<id>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; configuration-id(s) for retrieving the collector(s)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `562`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! appd get-collector --site "localhost:8081" 562
```
Output: 
```
ID 	VERSION	NAME            	TYPE 	HOSTNAME	USERNAME	PASSWORD	PORT	DATABASE NAME 
562	      0	test            	MONGO	        	master  	        	   0	             	
```

