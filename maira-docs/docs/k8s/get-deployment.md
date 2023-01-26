# k8s get-deployment

describe deployment

## Description

describes k8s deployment

## Synopsis

`k8s get-deployment [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] <name>`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the deployment  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! k8s get-deployment --site "localhost" --namespace "default" "kafka1"
```
Output: 
```
{
    "cluster": "minikube",
    "namespace": "default",
    "name": "kafka1",
    "labels": {
        "app": "kafka1"
    },
    "annotations": {
        "deployment.kubernetes.io/revision": "1"
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
                "message": "ReplicaSet \"kafka1-7cb95d9db6\" has successfully progressed."
            },
            {
                "type": 2,
                "condition_status": 2,
                "reason": "MinimumReplicasUnavailable",
                "message": "Deployment does not have minimum availability."
            }
        ]
    },
    "creation_time_millis": 1629190296000
}
```

