# kafka get-topics

Get Kafka topics

## Description

List the topics available in the cluster with the default options or get certain topic if topic name is provided

## Synopsis

`kafka get-topics [--site  <site>] [--cluster  <cluster>] [<topics>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of kafka cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "cluster-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `topics` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; topic name that need to be fetched from cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"topics-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! kafka get-topics --site "localhost" "test-topic-1" "test-topic-3"
```
Output: 
```
[
  {
    "test-topic-1": {
        "num_partitions": 1,
        "replication_factor": 1,
        "replica_assignment": {
            "0": {
                "replica": [3]
            }
        },
        "config_entries": {
            "segment.bytes": "1073741824"
        }
    },
    "test-topic-3": {
        "num_partitions": 1,
        "replication_factor": 1,
        "replica_assignment": {
            "0": {
                "replica": [2]
            }
        },
        "config_entries": {
            "segment.bytes": "1073741824"
        }
    }
  }
]
```

