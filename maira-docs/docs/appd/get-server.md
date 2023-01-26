# appd get-server

Retrieve database server

## Description

Returns single DB server if ID is provided. Else, returns all servers

## Synopsis

`appd get-server [--site  <site>] [--cluster  <cluster>] [<id>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; DB server ID for retrieving the DB Server  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `418`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! appd get-server --site "localhost:8081"
```
Output: 
```
ID 	VERSION	NAME	CONFIG ID	ROLE      	TYPE 	CREATED ON          	MODIFIED ON         	HOST     	PORT 	IP ADDRESS                                       
418	      0	test	      562	STANDALONE	MONGO	2022-09-29T11:41:38Z	2022-09-30T07:49:42Z	LOCALHOST	27017	192.168.49.1|172.18.0.1|172.17.0.1|192.168.5.159                                       
418	      0	test            	      562	STANDALONE	MONGO	2022-09-29T11:41:38Z	2022-09-30T07:49:42Z	LOCALHOST	27017	192.168.49.1|172.18.0.1|172.17.0.1|192.168.5.159	
420	      0	s               	      564	STANDALONE	MONGO	2022-09-29T12:28:10Z	2022-09-30T07:49:42Z	LOCALHOST	27017	192.168.49.1|172.18.0.1|172.17.0.1|192.168.5.159	
421	      0	b               	      565	STANDALONE	MONGO	2022-09-29T12:28:20Z	2022-09-30T07:49:42Z	LOCALHOST	27017	192.168.49.1|172.18.0.1|172.17.0.1|192.168.5.159	
426	      0	sample collector	      571	STANDALONE	MONGO	2022-09-30T09:01:00Z	2022-09-30T09:01:00Z	LOCALHOST	27017	192.168.49.1|172.18.0.1|172.17.0.1|192.168.5.159
```

