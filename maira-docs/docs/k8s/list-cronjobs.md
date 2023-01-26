# k8s list-cronjobs

Get k8s cronjob

## Description

Get the kubernetes cronjobs that satisfies the provided arguments

## Synopsis

`k8s list-cronjobs [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] [<selector>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Selector to select cronjobs  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"app=k8s,!tier"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
!k8s list-cronjobs
```
Output: 
```
[
    {
        "cluster": "in-cluster",
        "namespace": "default",
        "name": "hello",
        "schedule": "*/1 * * * *",
        "annotations": {
            "autopilot.gke.io/resource-adjustment": "{\"input\":{\"containers\":[{\"name\":\"hello\"}]},\"output\":{\"containers\":[{\"limits\":{\"cpu\":\"500m\",\"ephemeral-storage\":\"1Gi\",\"memory\":\"2Gi\"},\"requests\":{\"cpu\":\"500m\",\"ephemeral-storage\":\"1Gi\",\"memory\":\"2Gi\"},\"name\":\"hello\"}]},\"modified\":true}"
        },
        "concurrency_policy": 2
    }
]
```

