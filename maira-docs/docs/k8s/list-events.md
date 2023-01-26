# k8s list-events

List events

## Description

List k8s events in given namespace

## Synopsis

`k8s list-events [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] [--object_name  <object_name>] [--object_kind  <object_kind>]`

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


#### `object_name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; name of the object i.e pod name etc  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--object_name  "pod-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `object_kind` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; kind of the object  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--object_kind  "POD"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! k8s list-events --site "localhost" --object_name "kafka1-7cb95d9db6-xdkx4" --object_kind "pod"
```
Output: 
```
[
    {
        "name": "kafka1-7cb95d9db6-xdkx4.16a7c5487527f0b3",
        "reason": "SandboxChanged",
        "description": "Pod sandbox changed, it will be killed and re-created.",
        "cluster": "minikube",
        "namespace": "default",
        "involved_object_name": "kafka1-7cb95d9db6-xdkx4",
        "kind": 2
    },
    {
        "name": "kafka1-7cb95d9db6-xdkx4.16a7c54d938a1fa0",
        "reason": "Pulled",
        "description": "Container image \"wurstmeister/kafka\" already present on machine",
        "cluster": "minikube",
        "namespace": "default",
        "involved_object_name": "kafka1-7cb95d9db6-xdkx4",
        "kind": 2
    },
    {
        "name": "kafka1-7cb95d9db6-xdkx4.16a7c54f49eeea9b",
        "reason": "Created",
        "description": "Created container kafka1",
        "cluster": "minikube",
        "namespace": "default",
        "involved_object_name": "kafka1-7cb95d9db6-xdkx4",
        "kind": 2
    },
    {
        "name": "kafka1-7cb95d9db6-xdkx4.16a7c551bae8fa9a",
        "reason": "Started",
        "description": "Started container kafka1",
        "cluster": "minikube",
        "namespace": "default",
        "involved_object_name": "kafka1-7cb95d9db6-xdkx4",
        "kind": 2
    },
    {
        "name": "kafka1-7cb95d9db6-xdkx4.16a7c56325610ab2",
        "reason": "BackOff",
        "description": "Back-off restarting failed container",
        "cluster": "minikube",
        "namespace": "default",
        "involved_object_name": "kafka1-7cb95d9db6-xdkx4",
        "kind": 2
    }
]
```

