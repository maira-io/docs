# k8s create-job

create job

## Description

create job

## Synopsis

`k8s create-job [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] [--image  <image>] [--command  <command>] [--command_args  <command_args>]`

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


#### `image` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Image name for this job  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--image  "perl"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `command` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Command for this job  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--command  "perl"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `command_args` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; arg for the command  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--command_args  "-Mbignum=bpi"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
!k8s create-job --image "busybox" --command "/bin/sh" --command_args "-c" --command_args "date; echo Hello from the Kubernetes cluster"
```
Output: 
```
{
    "cluster": "in-cluster",
    "namespace": "default",
    "name": "default",
    "annotations": {
        "autopilot.gke.io/resource-adjustment": "{\"input\":{\"containers\":[{\"name\":\"default\"}]},\"output\":{\"containers\":[{\"limits\":{\"cpu\":\"500m\",\"ephemeral-storage\":\"1Gi\",\"memory\":\"2Gi\"},\"requests\":{\"cpu\":\"500m\",\"ephemeral-storage\":\"1Gi\",\"memory\":\"2Gi\"},\"name\":\"default\"}]},\"modified\":true}"
    },
    "labels": {
        "controller-uid": "1d962783-7c37-4701-af34-3527447858aa",
        "job-name": "default"
    }
}
```

