# k8s update-deployment

Update k8s deployment

## Description

Update k8s deployment with the provided labels and annotations

## Synopsis

`k8s update-deployment [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] <name> [--labels  <labels>] [--annotations  <annotations>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the deployment to be updated  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `labels` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; labels that are to be merged with existing ones  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--labels  "{"key1":"value1"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `annotations` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; annotations that are to be merged with existing ones  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--annotations  "{"key1":"value1"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! k8s update-deployment --site "localhost" --namespace "default" --labels `{"abc":"def"}` --annotations `{"ghi":"jkl"}` "kafka3"
```
Output: 
```
deployment kafka3 updated
```

