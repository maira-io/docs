# elasticsearch get-cluster-allocation

Get an explanation for shards current allocation

## Description

Get an explanation for shards current allocation.

## Synopsis

`elasticsearch get-cluster-allocation [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] --index  <index> <shard> [--primary ] --current_node  <current_node>`

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


#### `filter_path` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Path in the response. Filters and returns only parts of response matching this path  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--filter_path  "index.mappings.properties.*"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the index that you would like an explanation for.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--index  "students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `shard` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; ID of the shard that you would like an explanation for.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `primary` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, returns explanation for the primary shard for the given shard ID.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--primary  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `current_node` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Specifies the node ID or the name of the node to only explain a shard that is currently located on the specified node.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--current_node  "node-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! elasticsearch get-cluster-allocation --site "localhost" --cluster "elastic-1" --primary --index "class" 1 --current_node "beehyv-H81M-S"
```
Output: 
```
{
  "index": "class",
  "shard": 1,
  "primary": true,
  "current_state": "started",
  "current_node": {
    "id": "qQmwEJbkTdW-ixce6PgTFQ",
    "name": "beehyv-H81M-S",
    "transport_address": "127.0.0.1:9300",
    "attributes": {
      "ml.machine_memory": "16714620928",
      "xpack.installed": "true",
      "transform.node": "true",
      "ml.max_open_jobs": "512",
      "ml.max_jvm_size": "1073741824"
    },
    "weight_ranking": 1
  },
  "can_remain_on_current_node": "yes",
  "can_rebalance_cluster": "no",
  "can_rebalance_cluster_decisions": [
    {
      "decider": "rebalance_only_when_active",
      "decision": "NO",
      "explanation": "rebalancing is not allowed until all replicas in the cluster are active"
    },
    {
      "decider": "cluster_rebalance",
      "decision": "NO",
      "explanation": "the cluster has unassigned shards and cluster setting [cluster.routing.allocation.allow_rebalance] is set to [indices_all_active]"
    }
  ],
  "can_rebalance_to_other_node": "no",
  "rebalance_explanation": "rebalancing is not allowed"
}
```

