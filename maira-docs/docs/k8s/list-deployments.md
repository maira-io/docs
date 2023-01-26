# k8s list-deployments

Get k8s deployments

## Description

Get the kubernetes deployments satisfying the provided args

## Synopsis

`k8s list-deployments [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] [<selector>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Selector to select the deployments  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"app=kafka3,test"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! k8s list-deployments --site "localhost" --namespace "default"  "app=kafka3,test"
```
Output: 
```
[
    {
        "cluster": "minikube",
        "namespace": "default",
        "name": "kafka3",
        "labels": {
            "app": "kafka3",
            "test": "test"
        },
        "annotations": {
            "deployment.kubernetes.io/revision": "1",
            "testing": "testing"
        },
        "deployment_status": {
            "replicas": 1,
            "updated_replicas": 1,
            "unavailable_replicas": 1,
            "deployment_conditions": [
                {
                    "type": 3,
                    "condition_status": 1,
                    "reason": "NewReplicaSetAvailable",
                    "message": "ReplicaSet \"kafka3-77bf6975d4\" has successfully progressed."
                },
                {
                    "type": 2,
                    "condition_status": 2,
                    "reason": "MinimumReplicasUnavailable",
                    "message": "Deployment does not have minimum availability."
                }
            ]
        },
        "creation_time_millis": 1629190311000
    }
]
```

