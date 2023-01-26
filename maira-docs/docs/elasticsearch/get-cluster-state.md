# elasticsearch get-cluster-state

Get elastic search cluster state

## Description

Get metadata about state of cluster including nodes, indices in the cluster, cluster-level settings, locations of all shards

## Synopsis

`elasticsearch get-cluster-state [--site  <site>] [--cluster  <cluster>] [--target  <target>] [--filter_path  <filter_path>] [--metrics  <metrics>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of elastic search cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "elastic-default"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `elastic-default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `target` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of data streams, indices, and index aliases used to limit the request  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--target  "students or _all"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `filter_path` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Path in the response. Filters and returns only parts of response matching this path  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--filter_path  "index.mappings.properties.*"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `metrics` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of the options- blocks, master_node, metadata, nodes, routing_nodes, routing_table, version  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--metrics  "blocks"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! elasticsearch get-cluster-state --site "localhost" --cluster "elastic-1" --target "class" --target "test" --metrics "routing_table" --metrics "nodes"
```
Output: 
```
{
  "cluster_name": "default",
  "cluster_uuid": "PAlTKPcxSE6oX_tNaTBhgw",
  "nodes": {
    "qQmwEJbkTdW-ixce6PgTFQ": {
      "name": "beehyv-H81M-S",
      "ephemeral_id": "CAc2g-pVSj-ZS1dtoZs1Jw",
      "transport_address": "127.0.0.1:9300",
      "attributes": {
        "ml.machine_memory": "16714620928",
        "xpack.installed": "true",
        "transform.node": "true",
        "ml.max_open_jobs": "512",
        "ml.max_jvm_size": "1073741824"
      },
      "roles": [
        "data",
        "data_cold",
        "data_content",
        "data_frozen",
        "data_hot",
        "data_warm",
        "ingest",
        "master",
        "ml",
        "remote_cluster_client",
        "transform"
      ]
    }
  },
  "routing_table": {
    "indices": {
      "test": {
        "shards": {
          "0": [
            {
              "state": "STARTED",
              "primary": true,
              "node": "qQmwEJbkTdW-ixce6PgTFQ",
              "relocating_node": null,
              "shard": 0,
              "index": "test",
              "allocation_id": {
                "id": "a3T0tMbdRCaQwpKtgI8Kbw"
              }
            },
            {
              "state": "UNASSIGNED",
              "primary": false,
              "node": null,
              "relocating_node": null,
              "shard": 0,
              "index": "test",
              "recovery_source": {
                "type": "PEER"
              },
              "unassigned_info": {
                "reason": "INDEX_CREATED",
                "at": "2021-10-01T04:14:36.262Z",
                "delayed": false,
                "allocation_status": "no_attempt"
              }
            }
          ]
        }
      },
      "class": {
        "shards": {
          "1": [
            {
              "state": "STARTED",
              "primary": true,
              "node": "qQmwEJbkTdW-ixce6PgTFQ",
              "relocating_node": null,
              "shard": 1,
              "index": "class",
              "allocation_id": {
                "id": "WaBvZ0CVR8KKn01HYH1MGw"
              }
            },
            {
              "state": "UNASSIGNED",
              "primary": false,
              "node": null,
              "relocating_node": null,
              "shard": 1,
              "index": "class",
              "recovery_source": {
                "type": "PEER"
              },
              "unassigned_info": {
                "reason": "INDEX_CREATED",
                "at": "2021-09-28T10:10:43.225Z",
                "delayed": false,
                "allocation_status": "no_attempt"
              }
            }
          ],
          "0": [
            {
              "state": "STARTED",
              "primary": true,
              "node": "qQmwEJbkTdW-ixce6PgTFQ",
              "relocating_node": null,
              "shard": 0,
              "index": "class",
              "allocation_id": {
                "id": "J-fqcFhIRciuCZmUjAnycA"
              }
            },
            {
              "state": "UNASSIGNED",
              "primary": false,
              "node": null,
              "relocating_node": null,
              "shard": 0,
              "index": "class",
              "recovery_source": {
                "type": "PEER"
              },
              "unassigned_info": {
                "reason": "INDEX_CREATED",
                "at": "2021-09-28T10:10:43.225Z",
                "delayed": false,
                "allocation_status": "no_attempt"
              }
            }
          ]
        }
      }
    }
  }
}
```

