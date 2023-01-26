# k8s resize-hpa

Resize HPA

## Description

Resize Horizontal pod autoscaler with given values

## Synopsis

`k8s resize-hpa [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] <name> [--min  <min>] [--max  <max>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Cluster where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "in-cluster"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `in-cluster`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `namespace` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Kubernetes namespace  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--namespace  "namespace-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of HPA to be resized  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `min` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Minimum value of auto scaling  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--min  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `1`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `max` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Maximum value of auto scaling  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--max  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `1`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! k8s resize-hpa --site "localhost" --namespace "default" --min 1 --max 5 "php-apache"
```
Output: 
```
"hpa php-apache resized"
```

