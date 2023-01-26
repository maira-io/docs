# appd get-applications

List applications

## Description

List applications

## Synopsis

`appd get-applications [--site  <site>] [--cluster  <cluster>]`

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



## Examples

Input: 
```
! appd get-applications --site "localhost:8081"
```
Output: 
```
ID    	NAME     	DESCRIPTION                
892053	bookinfo 	                          	
864040	test-app 	My test application       	
887570	Cluster1 	                          	
864111	test app2	My second test application
```

