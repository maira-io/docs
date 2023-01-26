# k8s list-pods

Get k8s pods

## Description

Get all the kubernetes pods

## Synopsis

`k8s list-pods [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] [<selector>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Selector to select the pods to be listed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"app=pod1,!label1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
!k8s list-pods --site "localhost" --namespace "default" "app=kafka1,!label1"
```
Output: 
```
[
    {
        "cluster": "minikube",
        "namespace": "default",
        "name": "kafka1-7cb95d9db6-xdkx4",
        "containers": [
            {
                "name": "kafka1",
                "image": "wurstmeister/kafka:latest",
                "state": 4,
                "restart_count": 66,
                "StateDetails": {
                    "StateWaiting": {
                        "reason": "CrashLoopBackOff",
                        "message": "back-off 5m0s restarting failed container=kafka1 pod=kafka1-7cb95d9db6-xdkx4_default(01cf06a3-31cf-4044-ad75-cc36f074b741)"
                    }
                }
            }
        ],
        "node_ip": "192.168.49.2",
        "pod_ip": "172.17.0.3",
        "state": 3,
        "labels": {
            "app": "kafka1",
            "pod-template-hash": "7cb95d9db6"
        },
        "pod_conditions": [
            {
                "type": 2,
                "status": 1
            },
            {
                "type": 3,
                "status": 2
            },
            {
                "type": 1,
                "status": 2
            },
            {
                "type": 4,
                "status": 1
            }
        ],
        "status": "CrashLoopBackOff",
        "start_time_millis": 1629190297000
    }
]
```

