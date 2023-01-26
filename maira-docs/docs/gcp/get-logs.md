# gcp get-logs

Get logs

## Description

get gcp logs

## Synopsis

`gcp get-logs [--site  <site>] [--resource_type  <resource_type>] [--resource_labels  <resource_labels>] [--log_name  <log_name>] [--severity  <severity>] [--count  <count>] [--filter  <filter>] [--start_time  <start_time>] [--end_time  <end_time>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `resource_type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Types of one or more parent resources from which to retrieve log entries  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--resource_type  "gae_app"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `resource_labels` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; labels for this resource  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--resource_labels  "resource_labels-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `log_name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; log names to filter out the logs  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--log_name  "projects/macro-context-293714/logs/GCEGuestAgent"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `severity` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Severity of the logs  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--severity  "INFO"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `count` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; The maximum number of results to return from this request  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--count  100`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `filter` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; A filter that chooses which log entries to return  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--filter  "resource.type=gae_app AND severity>=ERROR"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `start_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; The time stamp indicating the earliest data to be returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--start_time  "60m ago"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `60m ago`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `end_time` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; The time stamp indicating the latest data to be returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--end_time  "now"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `now`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x=! gcp get-logs --site "localhost:8081" --severity "INFO"
```
Output: 
```
{
    "entries": [
        {
            "logName": "projects/macro-context-293714/logs/stdout",
            "resource": {
                "type": "k8s_container",
                "labels": {
                    "cluster_name": "cluster-2",
                    "container_name": "cassandra",
                    "location": "us-central1-c",
                    "namespace_name": "cass-operator",
                    "pod_name": "cluster1-dc1-default-sts-0",
                    "project_id": "macro-context-293714"
                }
            },
            "textPayload": "INFO  [nioEventLoopGroup-2-2] 2022-10-14 18:33:07,313 Cli.java:617 - address=/10.44.2.1:41262 url=/api/v0/probes/readiness status=200 OK",
            "timestamp": "2022-10-14T18:33:07.313548359Z",
            "receiveTimestamp": "2022-10-14T18:33:10.224351324Z",
            "severity": "INFO",
            "insertId": "0opzd7uiep1lo5ao",
            "labels": {
                "compute.googleapis.com/resource_name": "gke-cluster-2-default-pool-58fe26d9-qil9",
                "k8s-pod/app_kubernetes_io/managed-by": "cass-operator",
                "k8s-pod/cassandra_datastax_com/cluster": "cluster1",
                "k8s-pod/cassandra_datastax_com/datacenter": "dc1",
                "k8s-pod/cassandra_datastax_com/node-state": "Started",
                "k8s-pod/cassandra_datastax_com/rack": "default",
                "k8s-pod/cassandra_datastax_com/seed-node": "true",
                "k8s-pod/controller-revision-hash": "cluster1-dc1-default-sts-59c94c67cc",
                "k8s-pod/statefulset_kubernetes_io/pod-name": "cluster1-dc1-default-sts-0"
            }
        }
    ]
}
```

