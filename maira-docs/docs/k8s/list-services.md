# k8s list-services

List services

## Description

Lists all available k8s service in given namespace

## Synopsis

`k8s list-services [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] [<selector>]`

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


#### `selector` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Selector to select the services  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"app=k8s,!env"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
!k8s  list-services --site "localhost"
```
Output: 
```
[
    {
        "cluster": "minikube",
        "namespace": "default",
        "name": "kafka1-service",
        "type": 4,
        "labels": {
            "app": "kafka1-service"
        },
        "selector": {
            "app": "kafka1"
        }
    },
    {
        "cluster": "minikube",
        "namespace": "default",
        "name": "kafka2-service",
        "type": 4,
        "labels": {
            "app": "kafka2-service"
        },
        "selector": {
            "app": "kafka2"
        }
     }
]
```

