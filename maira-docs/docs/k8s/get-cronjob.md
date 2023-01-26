# k8s get-cronjob

describe cronjob

## Description

describes k8s cronjob

## Synopsis

`k8s get-cronjob [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] <name>`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the cronjob  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
!k8s get-cronjobs "hello"
```
Output: 
```
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
```

