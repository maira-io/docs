# k8s list-configmaps

Get k8s configmap

## Description

Get the kubernetes configmaps that satisfies the provided arguments

## Synopsis

`k8s list-configmaps [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] [<selector>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Selector to select the configmaps  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"env=prod,!partition"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! k8s list-configmaps
```
Output: 
```
[
    {
      "cluster": "in-cluster",
      "namespace": "default",
      "name": "dynamic-config",
      "annotations": {
          "meta.helm.sh/release-name": "test",
          "meta.helm.sh/release-namespace": "default"
      },
      "labels": {
          "app.kubernetes.io/instance": "test",
          "app.kubernetes.io/managed-by": "Helm",
          "app.kubernetes.io/version": "1.12.0",
          "helm.sh/chart": "test-0.12.0"
      },
      "data": {
          "dynamic_config.yaml": ""
      }
    }
]
```

