# elasticsearch get-node-stats

Get elastic search node statistics

## Description

Returns node statistics.

## Synopsis

`elasticsearch get-node-stats [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] [--nodes  <nodes>] [--metrics  <metrics>] [--index_metric  <index_metric>]`

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


#### `nodes` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of node IDs or names used to limit returned information.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--nodes  "node-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `metrics` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Limits the information returned to the specific metrics  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--metrics  "breaker, discovery, http, indices, ingest"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `index_metric` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Limit the information returned for indices metric to the specific index metrics. It can be used only if indices (or all) metric is specified  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--index_metric  "completion, docs, fielddata, flush, indexing, merge, query_cache"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! elasticsearch get-node-stats --site "localhost" --cluster "elastic-1"
```
Output: 
```
{
  "_nodes": {
    "total": 1,
    "successful": 1,
    "failed": 0
  },
  "cluster_name": "default",
  "nodes": {
    "qQmwEJbkTdW-ixce6PgTFQ": {
      "timestamp": 1635261092445,
      "name": "beehyv-H81M-S",
      "transport_address": "127.0.0.1:9300",
      "host": "localhost",
      "ip": "127.0.0.1:9300",
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
      ],
      "attributes": {
        "ml.machine_memory": "16714612736",
        "xpack.installed": "true",
        "transform.node": "true",
        "ml.max_open_jobs": "512",
        "ml.max_jvm_size": "1073741824"
      },
      "indices": {
        "docs": {
          "count": 46,
          "deleted": 4
        },
        "store": {
          "size_in_bytes": 98832154,
          "total_data_set_size_in_bytes": 98832154,
          "reserved_in_bytes": 0
        },
        "indexing": {
          "index_total": 41,
          "index_time_in_millis": 1154,
          "index_current": 0,
          "index_failed": 0,
          "delete_total": 41,
          "delete_time_in_millis": 36,
          "delete_current": 0,
          "noop_update_total": 0,
          "is_throttled": false,
          "throttle_time_in_millis": 0
        },
        "get": {
          "total": 0,
          "time_in_millis": 0,
          "exists_total": 0,
          "exists_time_in_millis": 0,
          "missing_total": 0,
          "missing_time_in_millis": 0,
          "current": 0
        },
        "search": {
          "open_contexts": 0,
          "query_total": 88,
          "query_time_in_millis": 109,
          "query_current": 0,
          "fetch_total": 88,
          "fetch_time_in_millis": 3436,
          "fetch_current": 0,
          "scroll_total": 3,
          "scroll_time_in_millis": 621,
          "scroll_current": 0,
          "suggest_total": 0,
          "suggest_time_in_millis": 0,
          "suggest_current": 0
        },
        "merges": {
          "current": 0,
          "current_docs": 0,
          "current_size_in_bytes": 0,
          "total": 3,
          "total_time_in_millis": 3103,
          "total_docs": 244,
          "total_size_in_bytes": 160809120,
          "total_stopped_time_in_millis": 0,
          "total_throttled_time_in_millis": 0,
          "total_auto_throttle_in_bytes": 272629760
        },
        "refresh": {
          "total": 52,
          "total_time_in_millis": 620,
          "external_total": 49,
          "external_total_time_in_millis": 637,
          "listeners": 0
        },
        "flush": {
          "total": 3,
          "periodic": 0,
          "total_time_in_millis": 6694
        },
        "warmer": {
          "current": 0,
          "total": 33,
          "total_time_in_millis": 1
        },
        "query_cache": {
          "memory_size_in_bytes": 0,
          "total_count": 0,
          "hit_count": 0,
          "miss_count": 0,
          "cache_size": 0,
          "cache_count": 0,
          "evictions": 0
        },
        "fielddata": {
          "memory_size_in_bytes": 0,
          "evictions": 0
        },
        "completion": {
          "size_in_bytes": 0
        },
        "segments": {
          "count": 5,
          "memory_in_bytes": 9172,
          "terms_memory_in_bytes": 5248,
          "stored_fields_memory_in_bytes": 2472,
          "term_vectors_memory_in_bytes": 0,
          "norms_memory_in_bytes": 512,
          "points_memory_in_bytes": 0,
          "doc_values_memory_in_bytes": 940,
          "index_writer_memory_in_bytes": 0,
          "version_map_memory_in_bytes": 6558,
          "fixed_bit_set_memory_in_bytes": 0,
          "max_unsafe_auto_id_timestamp": -1,
          "file_sizes": {}
        },
        "translog": {
          "operations": 3,
          "size_in_bytes": 1081,
          "uncommitted_operations": 3,
          "uncommitted_size_in_bytes": 1081,
          "earliest_last_modified_age": 6309
        },
        "request_cache": {
          "memory_size_in_bytes": 0,
          "evictions": 0,
          "hit_count": 0,
          "miss_count": 0
        },
        "recovery": {
          "current_as_source": 0,
          "current_as_target": 0,
          "throttle_time_in_millis": 0
        }
      },
      "os": {
        "timestamp": 1635261092577,
        "cpu": {
          "percent": 27,
          "load_average": {
            "1m": 2.88,
            "5m": 1.73,
            "15m": 1.51
          }
        },
        "mem": {
          "total_in_bytes": 16714612736,
          "free_in_bytes": 314003456,
          "used_in_bytes": 16400609280,
          "free_percent": 2,
          "used_percent": 98
        },
        "swap": {
          "total_in_bytes": 2147479552,
          "free_in_bytes": 479977472,
          "used_in_bytes": 1667502080
        },
        "cgroup": {
          "cpuacct": {
            "control_group": "/system.slice/elasticsearch.service",
            "usage_nanos": 60245077356
          },
          "cpu": {
            "control_group": "/system.slice/elasticsearch.service",
            "cfs_period_micros": 100000,
            "cfs_quota_micros": -1,
            "stat": {
              "number_of_elapsed_periods": 0,
              "number_of_times_throttled": 0,
              "time_throttled_nanos": 0
            }
          },
          "memory": {
            "control_group": "/system.slice/elasticsearch.service",
            "limit_in_bytes": "9223372036854771712",
            "usage_in_bytes": "1637425152"
          }
        }
      },
      "process": {
        "timestamp": 1635261092578,
        "open_file_descriptors": 340,
        "max_file_descriptors": 65535,
        "cpu": {
          "percent": 17,
          "total_in_millis": 57810
        },
        "mem": {
          "total_virtual_in_bytes": 5000441856
        }
      },
      "jvm": {
        "timestamp": 1635261092578,
        "uptime_in_millis": 63203,
        "mem": {
          "heap_used_in_bytes": 274752072,
          "heap_used_percent": 25,
          "heap_committed_in_bytes": 1073741824,
          "heap_max_in_bytes": 1073741824,
          "non_heap_used_in_bytes": 146866488,
          "non_heap_committed_in_bytes": 152174592,
          "pools": {
            "young": {
              "used_in_bytes": 155189248,
              "max_in_bytes": 0,
              "peak_used_in_bytes": 629145600,
              "peak_max_in_bytes": 0
            },
            "old": {
              "used_in_bytes": 71699456,
              "max_in_bytes": 1073741824,
              "peak_used_in_bytes": 71699456,
              "peak_max_in_bytes": 1073741824
            },
            "survivor": {
              "used_in_bytes": 47863368,
              "max_in_bytes": 0,
              "peak_used_in_bytes": 54525952,
              "peak_max_in_bytes": 0
            }
          }
        },
        "threads": {
          "count": 60,
          "peak_count": 63
        },
        "gc": {
          "collectors": {
            "young": {
              "collection_count": 13,
              "collection_time_in_millis": 310
            },
            "old": {
              "collection_count": 0,
              "collection_time_in_millis": 0
            }
          }
        },
        "buffer_pools": {
          "mapped": {
            "count": 11,
            "used_in_bytes": 17212,
            "total_capacity_in_bytes": 17212
          },
          "direct": {
            "count": 37,
            "used_in_bytes": 1467246,
            "total_capacity_in_bytes": 1467245
          },
          "mapped - 'non-volatile memory'": {
            "count": 0,
            "used_in_bytes": 0,
            "total_capacity_in_bytes": 0
          }
        },
        "classes": {
          "current_loaded_count": 23326,
          "total_loaded_count": 23326,
          "total_unloaded_count": 0
        }
      },
      "thread_pool": {
        "analyze": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "auto_complete": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "ccr": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "fetch_shard_started": {
          "threads": 8,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 8,
          "completed": 13
        },
        "fetch_shard_store": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "flush": {
          "threads": 1,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 1,
          "completed": 3
        },
        "force_merge": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "generic": {
          "threads": 9,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 9,
          "completed": 508
        },
        "get": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "listener": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "management": {
          "threads": 4,
          "queue": 0,
          "active": 1,
          "rejected": 0,
          "largest": 4,
          "completed": 71
        },
        "ml_datafeed": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "ml_job_comms": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "ml_utility": {
          "threads": 1,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 1,
          "completed": 41
        },
        "refresh": {
          "threads": 2,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 2,
          "completed": 240
        },
        "rollup_indexing": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "search": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "search_throttled": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "searchable_snapshots_cache_fetch_async": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "searchable_snapshots_cache_prewarming": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "security-crypto": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "security-token-key": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "snapshot": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "snapshot_meta": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "system_critical_read": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "system_critical_write": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "system_read": {
          "threads": 2,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 2,
          "completed": 173
        },
        "system_write": {
          "threads": 2,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 2,
          "completed": 44
        },
        "transform_indexing": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "warmer": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "watcher": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "write": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        }
      },
      "fs": {
        "timestamp": 1635261092580,
        "total": {
          "total_in_bytes": 982900588544,
          "free_in_bytes": 810727624704,
          "available_in_bytes": 760727539712
        },
        "data": [
          {
            "path": "/var/lib/elasticsearch/nodes/0",
            "mount": "/ (/dev/sda2)",
            "type": "ext4",
            "total_in_bytes": 982900588544,
            "free_in_bytes": 810727624704,
            "available_in_bytes": 760727539712
          }
        ],
        "io_stats": {
          "devices": [
            {
              "device_name": "sda2",
              "operations": 5772,
              "read_operations": 2537,
              "write_operations": 3235,
              "read_kilobytes": 135748,
              "write_kilobytes": 440196,
              "io_time_in_millis": 34376
            }
          ],
          "total": {
            "operations": 5772,
            "read_operations": 2537,
            "write_operations": 3235,
            "read_kilobytes": 135748,
            "write_kilobytes": 440196,
            "io_time_in_millis": 34376
          }
        }
      },
      "transport": {
        "server_open": 0,
        "total_outbound_connections": 0,
        "rx_count": 0,
        "rx_size_in_bytes": 0,
        "tx_count": 0,
        "tx_size_in_bytes": 0
      },
      "http": {
        "current_open": 1,
        "total_opened": 1,
        "clients": [
          {
            "id": 407079349,
            "agent": "Go-http-client/1.1",
            "local_address": "127.0.0.1:9200",
            "remote_address": "127.0.0.1:55736",
            "last_uri": "/_nodes/stats",
            "opened_time_millis": 1635261092315,
            "last_request_time_millis": 1635261092315,
            "request_count": 1,
            "request_size_bytes": 0
          }
        ]
      },
      "breakers": {
        "request": {
          "limit_size_in_bytes": 644245094,
          "limit_size": "614.3mb",
          "estimated_size_in_bytes": 0,
          "estimated_size": "0b",
          "overhead": 1.0,
          "tripped": 0
        },
        "fielddata": {
          "limit_size_in_bytes": 429496729,
          "limit_size": "409.5mb",
          "estimated_size_in_bytes": 0,
          "estimated_size": "0b",
          "overhead": 1.03,
          "tripped": 0
        },
        "in_flight_requests": {
          "limit_size_in_bytes": 1073741824,
          "limit_size": "1gb",
          "estimated_size_in_bytes": 0,
          "estimated_size": "0b",
          "overhead": 2.0,
          "tripped": 0
        },
        "model_inference": {
          "limit_size_in_bytes": 536870912,
          "limit_size": "512mb",
          "estimated_size_in_bytes": 0,
          "estimated_size": "0b",
          "overhead": 1.0,
          "tripped": 0
        },
        "eql_sequence": {
          "limit_size_in_bytes": 536870912,
          "limit_size": "512mb",
          "estimated_size_in_bytes": 0,
          "estimated_size": "0b",
          "overhead": 1.0,
          "tripped": 0
        },
        "accounting": {
          "limit_size_in_bytes": 1073741824,
          "limit_size": "1gb",
          "estimated_size_in_bytes": 9172,
          "estimated_size": "8.9kb",
          "overhead": 1.0,
          "tripped": 0
        },
        "parent": {
          "limit_size_in_bytes": 1020054732,
          "limit_size": "972.7mb",
          "estimated_size_in_bytes": 274752072,
          "estimated_size": "262mb",
          "overhead": 1.0,
          "tripped": 0
        }
      },
      "script": {
        "compilations": 1,
        "cache_evictions": 0,
        "compilation_limit_triggered": 0
      },
      "discovery": {
        "cluster_state_queue": {
          "total": 0,
          "pending": 0,
          "committed": 0
        },
        "published_cluster_states": {
          "full_states": 2,
          "incompatible_diffs": 0,
          "compatible_diffs": 24
        }
      },
      "ingest": {
        "total": {
          "count": 0,
          "time_in_millis": 0,
          "current": 0,
          "failed": 0
        },
        "pipelines": {
          "xpack_monitoring_6": {
            "count": 0,
            "time_in_millis": 0,
            "current": 0,
            "failed": 0,
            "processors": [
              {
                "script": {
                  "type": "script",
                  "stats": {
                    "count": 0,
                    "time_in_millis": 0,
                    "current": 0,
                    "failed": 0
                  }
                }
              },
              {
                "gsub": {
                  "type": "gsub",
                  "stats": {
                    "count": 0,
                    "time_in_millis": 0,
                    "current": 0,
                    "failed": 0
                  }
                }
              }
            ]
          },
          "xpack_monitoring_7": {
            "count": 0,
            "time_in_millis": 0,
            "current": 0,
            "failed": 0,
            "processors": []
          }
        }
      },
      "adaptive_selection": {},
      "script_cache": {
        "sum": {
          "compilations": 1,
          "cache_evictions": 0,
          "compilation_limit_triggered": 0
        },
        "contexts": [
          {
            "context": "aggregation_selector",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "aggs",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "aggs_combine",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "aggs_init",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "aggs_map",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "aggs_reduce",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "analysis",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "boolean_field",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "bucket_aggregation",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "date_field",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "double_field",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "field",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "filter",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "geo_point_field",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "ingest",
            "compilations": 1,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "ingest_template",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "interval",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "ip_field",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "keyword_field",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "long_field",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "moving-function",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "number_sort",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "painless_test",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "processor_conditional",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "score",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "script_heuristic",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "similarity",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "similarity_weight",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "string_sort",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "template",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "terms_set",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "update",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "watcher_condition",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "watcher_transform",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          },
          {
            "context": "xpack_template",
            "compilations": 0,
            "cache_evictions": 0,
            "compilation_limit_triggered": 0
          }
        ]
      },
      "indexing_pressure": {
        "memory": {
          "current": {
            "combined_coordinating_and_primary_in_bytes": 0,
            "coordinating_in_bytes": 0,
            "primary_in_bytes": 0,
            "replica_in_bytes": 0,
            "all_in_bytes": 0
          },
          "total": {
            "combined_coordinating_and_primary_in_bytes": 42108662,
            "coordinating_in_bytes": 42108662,
            "primary_in_bytes": 42109926,
            "replica_in_bytes": 0,
            "all_in_bytes": 42108662,
            "coordinating_rejections": 0,
            "primary_rejections": 0,
            "replica_rejections": 0
          },
          "limit_in_bytes": 107374182
        }
      }
    }
  }
}
```

