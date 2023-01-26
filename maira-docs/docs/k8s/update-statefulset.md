# k8s update-statefulset

Update k8s statefulset

## Description

Update k8s statefulset with the provided labels and annotations

## Synopsis

`k8s update-statefulset [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] <name> [--labels  <labels>] [--annotations  <annotations>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the statefulset to be updated  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `labels` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; labels that are to be merged with existing ones  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--labels  "{"label1":"val1"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `annotations` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; annotations that are to be merged with existing ones  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--annotations  "{"ann1":"val1"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! k8s update-statefulset --site "localhost" --namespace "default" --labels `{"abc":"defg","zxc":"vbn"}` --annotations `{"ccc":"fsfd"}` "web"
```
Output: 
```
statefulset web updated
```

