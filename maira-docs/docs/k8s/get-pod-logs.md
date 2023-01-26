# k8s get-pod-logs

Get k8s pod logs

## Description

Get logs from kubernetes pods

## Synopsis

`k8s get-pod-logs [--site  <site>] [--cluster  <cluster>] [--namespace  |-n  <namespace>] [<name>] [--since  <since>] [--since_time  <since_time>] [--tail_lines  <tail_lines>] [--limit_bytes  <limit_bytes>] [--label_selector  <label_selector>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; name of the pod  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `since` - (duration)

&nbsp;&nbsp;&nbsp;&nbsp; Duration for which logs are required. Only one of since or since_time can be specified  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--since  "5 seconds"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `since_time` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Timestamp from which logs required. Should follow RFC3339 standard  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--since_time  "2019-10-12T07:20:50.52Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `tail_lines` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; the number of lines from the end of the logs to show  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--tail_lines  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `limit_bytes` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; the number of bytes to read from the server before terminating the log output  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--limit_bytes  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `label_selector` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; label selector to select pods to get logs for. Only used if name is not provided  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--label_selector  "app=myapp"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! k8s get-pod-logs --site "localhost" "hello-1629338580-jw8sf"
```
Output: 
```
{
     "logs": [
          "Thu Aug 19 02:03:06 UTC 2021",
          "Hello from the Kubernetes cluster",
          ""
     ]
}
```

