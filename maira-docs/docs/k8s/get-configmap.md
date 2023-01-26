# k8s get-configmap

describe configmap

## Description

describes k8s configmap

## Synopsis

`k8s get-configmap [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] <name>`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the configmap  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! k8s get-configmaps "dynamic-config"
```
Output: 
```
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
```

