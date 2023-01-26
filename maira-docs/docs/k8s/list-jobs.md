# k8s list-jobs

Get k8s jobs

## Description

Get the kubernetes jobs that satisfies the provided arguments

## Synopsis

`k8s list-jobs [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] [<selector>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Selector to select the jobs  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"app=k8s,!env"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! k8s list-jobs
```
Output: 
```
[
    {
        "cluster": "in-cluster",
        "namespace": "default",
        "name": "hello-1632241680",
        "labels": {
            "controller-uid": "6047d70f-9154-4e7e-8dcd-1d0bad4a5059",
            "job-name": "hello-1632241680"
        }
    },
    {
        "cluster": "in-cluster",
        "namespace": "default",
        "name": "hello-1632241740",
        "labels": {
            "controller-uid": "2dad1e38-aa75-4e85-9495-f1bd8f50d823",
            "job-name": "hello-1632241740"
        }
    },
    {
        "cluster": "in-cluster",
        "namespace": "default",
        "name": "hello-1632241800",
        "labels": {
            "controller-uid": "c5195f3c-71d1-46f2-bfc0-e56de80314cf",
            "job-name": "hello-1632241800"
        }
    }
]
```

