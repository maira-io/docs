# elasticsearch cluster-reroute

Change the allocation of shards in a cluster.

## Description

Change the allocation of shards in a cluster.

## Synopsis

`elasticsearch cluster-reroute [--site  <site>] [--cluster  <cluster>] [--dry_run ] [--explain ] [--retry_failed ] --command  <command> --index  <index> --shard  <shard> [--node  <node>] [--to_node  <to_node>] [--allow_primary ]`

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


#### `dry_run` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, then the request simulates the operation and returns the resulting state, but will not actually perform the requested changes.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--dry_run  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `explain` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, then the response contains an explanation of why the commands can or cannot be executed.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--explain  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `retry_failed` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, then retries allocation of shards that are blocked due to too many subsequent allocation failures  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--retry_failed  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `command` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Command to perform. Can be either move/ cancel/ allocate_replica  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--command  "move"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Index on which command is performed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--index  "students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `shard` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Shard on which command is performed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--shard  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `node` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Node from which shard should move or node from which shared allocation should be cancelled or node to which unassigned replica should be allocated  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--node  "node2"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `to_node` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; In case of move command, node to which shard is moved  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--to_node  "node2"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `allow_primary` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; By default only replica shard allocations can be cancelled. If it is necessary to cancel the allocation of a primary shard then the allow_primary flag must be set  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--allow_primary  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch cluster-reroute --site "localhost" --cluster "elastic-1" --shard 1 --index "test1" --command "cancel" --node "beehyv-H81M-S"
```
Output: 
```
{
  "acknowledged": true,
  "state": {
    "cluster_uuid": "PAlTKPcxSE6oX_tNaTBhgw",
    "version": 205,
    "state_uuid": "X1sf1LsgR32BBBuZA8k98w",
    "master_node": "qQmwEJbkTdW-ixce6PgTFQ",
    "blocks": {
      "indices": {
        "my_source_index": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_91": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_6": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_7": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_1": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_90": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_2": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_3": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_18": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_9": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_4": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "split_index_1": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        },
        "my_target_index_17": {
          "8": {
            "description": "index write (api)",
            "retryable": false,
            "levels": ["write"]
          }
        }
      }
    },
    "nodes": {
      "qQmwEJbkTdW-ixce6PgTFQ": {
        "name": "beehyv-H81M-S",
        "ephemeral_id": "bBC23TPMRceAV4tbM4P-1Q",
        "transport_address": "127.0.0.1:9300",
        "attributes": {
          "ml.machine_memory": "16714612736",
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
        "my_target_index_18": {
          "shards": {
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_18",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_18",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_18",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_18",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_target_index_9": {
          "shards": {
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_9",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_9",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_target_index_2": {
          "shards": {
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_2",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_2",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_target_index": {
          "shards": {
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_target_index_17": {
          "shards": {
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_17",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_17",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_17",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_17",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_target_index_7": {
          "shards": {
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_7",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.787Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_7",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.787Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_7",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.787Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_7",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.787Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_target_index_1": {
          "shards": {
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_1",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_source_index": {
          "shards": {
            "3": [
              {
                "state": "STARTED",
                "primary": true,
                "node": "qQmwEJbkTdW-ixce6PgTFQ",
                "relocating_node": null,
                "shard": 3,
                "index": "my_source_index",
                "allocation_id": {
                  "id": "8VUKYE24Si2H3Xipiq313g"
                }
              }
            ],
            "2": [
              {
                "state": "STARTED",
                "primary": true,
                "node": "qQmwEJbkTdW-ixce6PgTFQ",
                "relocating_node": null,
                "shard": 2,
                "index": "my_source_index",
                "allocation_id": {
                  "id": "2S0iGAliTC6HdE1zmXZqBA"
                }
              }
            ],
            "1": [
              {
                "state": "STARTED",
                "primary": true,
                "node": "qQmwEJbkTdW-ixce6PgTFQ",
                "relocating_node": null,
                "shard": 1,
                "index": "my_source_index",
                "allocation_id": {
                  "id": "GGi3hGSdRDCtDbYmxAiNpA"
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
                "index": "my_source_index",
                "allocation_id": {
                  "id": "7iCzE1RKQ5O3BSf_ztL6aA"
                }
              }
            ]
          }
        },
        "my_target_index_90": {
          "shards": {
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_90",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_90",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_90",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my-index-000001": {
          "shards": {
            "0": [
              {
                "state": "STARTED",
                "primary": true,
                "node": "qQmwEJbkTdW-ixce6PgTFQ",
                "relocating_node": null,
                "shard": 0,
                "index": "my-index-000001",
                "allocation_id": {
                  "id": "yGUGaycRTBStQU-SdZwIqQ"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my-index-000001",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
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
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
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
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              }
            ]
          }
        },
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
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              }
            ]
          }
        },
        "split_index_1": {
          "shards": {
            "5": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 5,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 5,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "2": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 2,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 2,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "7": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 7,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 7,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "3": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 3,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 3,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "6": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 6,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 6,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "4": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 4,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 4,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "split_index_1",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_target_index_91": {
          "shards": {
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_91",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_91",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_91",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_91",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_target_index_6": {
          "shards": {
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_6",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_6",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_6",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_6",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        ".geoip_databases": {
          "shards": {
            "0": [
              {
                "state": "STARTED",
                "primary": true,
                "node": "qQmwEJbkTdW-ixce6PgTFQ",
                "relocating_node": null,
                "shard": 0,
                "index": ".geoip_databases",
                "allocation_id": {
                  "id": "ti9-aNLqQbq9-12w20KG1Q"
                }
              }
            ]
          }
        },
        "test1": {
          "shards": {
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "test1",
                "recovery_source": {
                  "type": "EXISTING_STORE",
                  "bootstrap_new_history_uuid": false
                },
                "unassigned_info": {
                  "reason": "REROUTE_CANCELLED",
                  "at": "2021-10-18T14:46:59.806Z",
                  "delayed": false,
                  "allocation_status": "fetching_shard_data"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "test1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "test1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "test1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              }
            ],
            "2": [
              {
                "state": "STARTED",
                "primary": true,
                "node": "qQmwEJbkTdW-ixce6PgTFQ",
                "relocating_node": null,
                "shard": 2,
                "index": "test1",
                "allocation_id": {
                  "id": "FJ3-uA9hTv2TM8ryBwDEew"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 2,
                "index": "test1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 2,
                "index": "test1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 2,
                "index": "test1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
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
                "index": "test1",
                "allocation_id": {
                  "id": "lYTGC06PRBCep416kGKMQQ"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "test1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "test1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "test1",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.792Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              }
            ]
          }
        },
        "my_target_index_4": {
          "shards": {
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_4",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_4",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_4",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_4",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        },
        "my_target_index_3": {
          "shards": {
            "1": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_3",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 1,
                "index": "my_target_index_3",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ],
            "0": [
              {
                "state": "UNASSIGNED",
                "primary": false,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_3",
                "recovery_source": {
                  "type": "PEER"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "no_attempt"
                }
              },
              {
                "state": "UNASSIGNED",
                "primary": true,
                "node": null,
                "relocating_node": null,
                "shard": 0,
                "index": "my_target_index_3",
                "recovery_source": {
                  "type": "LOCAL_SHARDS"
                },
                "unassigned_info": {
                  "reason": "CLUSTER_RECOVERED",
                  "at": "2021-10-14T11:40:34.791Z",
                  "delayed": false,
                  "allocation_status": "deciders_no"
                }
              }
            ]
          }
        }
      }
    },
    "routing_nodes": {
      "unassigned": [
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_18",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_18",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_18",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_18",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_9",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_9",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_2",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_2",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_17",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_17",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_17",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_17",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_7",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.787Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_7",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.787Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_7",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.787Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_7",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.787Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_1",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_90",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_90",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_90",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my-index-000001",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
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
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
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
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
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
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 5,
          "index": "split_index_1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 5,
          "index": "split_index_1",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 2,
          "index": "split_index_1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 2,
          "index": "split_index_1",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 7,
          "index": "split_index_1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 7,
          "index": "split_index_1",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 3,
          "index": "split_index_1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 3,
          "index": "split_index_1",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 6,
          "index": "split_index_1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 6,
          "index": "split_index_1",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "split_index_1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "split_index_1",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 4,
          "index": "split_index_1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 4,
          "index": "split_index_1",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "split_index_1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "split_index_1",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_91",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_91",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_91",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_91",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_6",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_6",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_6",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_6",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "test1",
          "recovery_source": {
            "type": "EXISTING_STORE",
            "bootstrap_new_history_uuid": false
          },
          "unassigned_info": {
            "reason": "REROUTE_CANCELLED",
            "at": "2021-10-18T14:46:59.806Z",
            "delayed": false,
            "allocation_status": "fetching_shard_data"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "test1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "test1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "test1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 2,
          "index": "test1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 2,
          "index": "test1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 2,
          "index": "test1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "test1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "test1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "test1",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.792Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_4",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_4",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_4",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_4",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_3",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 1,
          "index": "my_target_index_3",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": false,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_3",
          "recovery_source": {
            "type": "PEER"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "no_attempt"
          }
        },
        {
          "state": "UNASSIGNED",
          "primary": true,
          "node": null,
          "relocating_node": null,
          "shard": 0,
          "index": "my_target_index_3",
          "recovery_source": {
            "type": "LOCAL_SHARDS"
          },
          "unassigned_info": {
            "reason": "CLUSTER_RECOVERED",
            "at": "2021-10-14T11:40:34.791Z",
            "delayed": false,
            "allocation_status": "deciders_no"
          }
        }
      ],
      "nodes": {
        "qQmwEJbkTdW-ixce6PgTFQ": [
          {
            "state": "STARTED",
            "primary": true,
            "node": "qQmwEJbkTdW-ixce6PgTFQ",
            "relocating_node": null,
            "shard": 3,
            "index": "my_source_index",
            "allocation_id": {
              "id": "8VUKYE24Si2H3Xipiq313g"
            }
          },
          {
            "state": "STARTED",
            "primary": true,
            "node": "qQmwEJbkTdW-ixce6PgTFQ",
            "relocating_node": null,
            "shard": 2,
            "index": "my_source_index",
            "allocation_id": {
              "id": "2S0iGAliTC6HdE1zmXZqBA"
            }
          },
          {
            "state": "STARTED",
            "primary": true,
            "node": "qQmwEJbkTdW-ixce6PgTFQ",
            "relocating_node": null,
            "shard": 1,
            "index": "my_source_index",
            "allocation_id": {
              "id": "GGi3hGSdRDCtDbYmxAiNpA"
            }
          },
          {
            "state": "STARTED",
            "primary": true,
            "node": "qQmwEJbkTdW-ixce6PgTFQ",
            "relocating_node": null,
            "shard": 0,
            "index": "my_source_index",
            "allocation_id": {
              "id": "7iCzE1RKQ5O3BSf_ztL6aA"
            }
          },
          {
            "state": "STARTED",
            "primary": true,
            "node": "qQmwEJbkTdW-ixce6PgTFQ",
            "relocating_node": null,
            "shard": 0,
            "index": "my-index-000001",
            "allocation_id": {
              "id": "yGUGaycRTBStQU-SdZwIqQ"
            }
          },
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
            "state": "STARTED",
            "primary": true,
            "node": "qQmwEJbkTdW-ixce6PgTFQ",
            "relocating_node": null,
            "shard": 0,
            "index": ".geoip_databases",
            "allocation_id": {
              "id": "ti9-aNLqQbq9-12w20KG1Q"
            }
          },
          {
            "state": "STARTED",
            "primary": true,
            "node": "qQmwEJbkTdW-ixce6PgTFQ",
            "relocating_node": null,
            "shard": 2,
            "index": "test1",
            "allocation_id": {
              "id": "FJ3-uA9hTv2TM8ryBwDEew"
            }
          },
          {
            "state": "STARTED",
            "primary": true,
            "node": "qQmwEJbkTdW-ixce6PgTFQ",
            "relocating_node": null,
            "shard": 0,
            "index": "test1",
            "allocation_id": {
              "id": "lYTGC06PRBCep416kGKMQQ"
            }
          }
        ]
      }
    },
    "security_tokens": {}
  }
}
```

