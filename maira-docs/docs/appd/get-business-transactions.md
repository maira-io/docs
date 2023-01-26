# appd get-business-transactions

Get business transactions for a given application

## Description

Get business transactions for a given application

## Synopsis

`appd get-business-transactions [--site  <site>] [--cluster  <cluster>] <application>`

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



## Examples

Input: 
```
x = ! appd get-business-transactions --site "localhost:8081" "bookinfo"
```
Output: 
```
ID   	NAME                    	ENTRY-POINT-TYPE	INTERNAL-NAME           	TIER-ID	TIER       	BACKGROUND 
94244	/productpage            	PYTHON_WEB      	/productpage            	 939727	productpage	false     	
94245	_APPDYNAMICS_DEFAULT_TX_	POJO            	_APPDYNAMICS_DEFAULT_TX_	 939711	reviews    	false     	
94247	/login                  	PYTHON_WEB      	/login                  	 939727	productpage	false     	
94273	_APPDYNAMICS_DEFAULT_TX_	PYTHON_WEB      	_APPDYNAMICS_DEFAULT_TX_	 939727	productpage	false     	
94723	/logout                 	PYTHON_WEB      	/logout                 	 939727	productpage	false     
```

