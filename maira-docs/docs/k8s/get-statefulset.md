# k8s get-statefulset

describe statefulset

## Description

describes k8s statefulset

## Synopsis

`k8s get-statefulset [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] <name>`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the statefulset  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
x = !k8s get-statefulset "temporaltest-cassandra"
```
Output: 
```
{
    "cluster": "in-cluster",
    "namespace": "default",
    "name": "temporaltest-cassandra",
    "labels": {
        "app": "cassandra",
        "app.kubernetes.io/managed-by": "Helm",
        "chart": "cassandra-0.14.3",
        "heritage": "Helm",
        "release": "temporaltest"
    },
    "annotations": {
        "autopilot.gke.io/resource-adjustment": "{\"input\":{\"containers\":[{\"name\":\"temporaltest-cassandra\"}]},\"output\":{\"containers\":[{\"limits\":{\"cpu\":\"500m\",\"ephemeral-storage\":\"1Gi\",\"memory\":\"2Gi\"},\"requests\":{\"cpu\":\"500m\",\"ephemeral-storage\":\"1Gi\",\"memory\":\"2Gi\"},\"name\":\"temporaltest-cassandra\"}]},\"modified\":true}",
        "meta.helm.sh/release-name": "temporaltest",
        "meta.helm.sh/release-namespace": "default"
    },
    "status": {
        "replicas": 1,
        "updated_replicas": 1,
        "ready_replicas": 1
    },
    "creation_time_millis": 1632240106000
}
```

