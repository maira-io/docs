# elasticsearch get-cluster-statistics

Get elastic search cluster statistics

## Description

Returns cluster statistics.

## Synopsis

`elasticsearch get-cluster-statistics [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] [--node_filters  <node_filters>]`

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


#### `node_filters` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Used to limit the information in response  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--node_filters  "_local or _all or _master or _node_name_* or 10.0.0.3,10.0.0.4 or 10.0.0.* or _all,master:false"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch get-cluster-statistics --site "localhost" --cluster "elastic-1"
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
  "cluster_uuid": "PAlTKPcxSE6oX_tNaTBhgw",
  "timestamp": 1633094291141,
  "status": "red",
  "indices": {
    "count": 6,
    "shards": {
      "total": 12,
      "primaries": 12,
      "replication": 0.0,
      "index": {
        "shards": {
          "min": 1,
          "max": 4,
          "avg": 2.0
        },
        "primaries": {
          "min": 1,
          "max": 4,
          "avg": 2.0
        },
        "replication": {
          "min": 0.0,
          "max": 0.0,
          "avg": 0.0
        }
      }
    },
    "docs": {
      "count": 46,
      "deleted": 44
    },
    "store": {
      "size_in_bytes": 43159566,
      "total_data_set_size_in_bytes": 43159566,
      "reserved_in_bytes": 0
    },
    "fielddata": {
      "memory_size_in_bytes": 0,
      "evictions": 0
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
    "completion": {
      "size_in_bytes": 0
    },
    "segments": {
      "count": 3,
      "memory_in_bytes": 5076,
      "terms_memory_in_bytes": 2880,
      "stored_fields_memory_in_bytes": 1496,
      "term_vectors_memory_in_bytes": 0,
      "norms_memory_in_bytes": 256,
      "points_memory_in_bytes": 0,
      "doc_values_memory_in_bytes": 444,
      "index_writer_memory_in_bytes": 0,
      "version_map_memory_in_bytes": 0,
      "fixed_bit_set_memory_in_bytes": 0,
      "max_unsafe_auto_id_timestamp": -1,
      "file_sizes": {}
    },
    "mappings": {
      "field_types": [
        {
          "name": "date",
          "count": 1,
          "index_count": 1,
          "script_count": 0
        },
        {
          "name": "keyword",
          "count": 3,
          "index_count": 1,
          "script_count": 0
        },
        {
          "name": "object",
          "count": 1,
          "index_count": 1,
          "script_count": 0
        },
        {
          "name": "text",
          "count": 4,
          "index_count": 2,
          "script_count": 0
        }
      ],
      "runtime_field_types": []
    },
    "analysis": {
      "char_filter_types": [],
      "tokenizer_types": [],
      "filter_types": [],
      "analyzer_types": [],
      "built_in_char_filters": [],
      "built_in_tokenizers": [],
      "built_in_filters": [],
      "built_in_analyzers": []
    },
    "versions": [
      {
        "version": "7.14.2",
        "index_count": 19,
        "primary_shard_count": 40,
        "total_primary_bytes": 43159566
      }
    ]
  },
  "nodes": {
    "count": {
      "total": 1,
      "coordinating_only": 0,
      "data": 1,
      "data_cold": 1,
      "data_content": 1,
      "data_frozen": 1,
      "data_hot": 1,
      "data_warm": 1,
      "ingest": 1,
      "master": 1,
      "ml": 1,
      "remote_cluster_client": 1,
      "transform": 1,
      "voting_only": 0
    },
    "versions": ["7.14.2"],
    "os": {
      "available_processors": 4,
      "allocated_processors": 4,
      "names": [
        {
          "name": "Linux",
          "count": 1
        }
      ],
      "pretty_names": [
        {
          "pretty_name": "Ubuntu 18.04.6 LTS",
          "count": 1
        }
      ],
      "architectures": [
        {
          "arch": "amd64",
          "count": 1
        }
      ],
      "mem": {
        "total_in_bytes": 16714620928,
        "free_in_bytes": 190173184,
        "used_in_bytes": 16524447744,
        "free_percent": 1,
        "used_percent": 99
      }
    },
    "process": {
      "cpu": {
        "percent": 0
      },
      "open_file_descriptors": {
        "min": 338,
        "max": 338,
        "avg": 338
      }
    },
    "jvm": {
      "max_uptime_in_millis": 345092725,
      "versions": [
        {
          "version": "16.0.2",
          "vm_name": "OpenJDK 64-Bit Server VM",
          "vm_version": "16.0.2+7",
          "vm_vendor": "Eclipse Foundation",
          "bundled_jdk": true,
          "using_bundled_jdk": true,
          "count": 1
        }
      ],
      "mem": {
        "heap_used_in_bytes": 210448752,
        "heap_max_in_bytes": 1073741824
      },
      "threads": 55
    },
    "fs": {
      "total_in_bytes": 982900588544,
      "free_in_bytes": 835309924352,
      "available_in_bytes": 785309839360
    },
    "plugins": [],
    "network_types": {
      "transport_types": {
        "security4": 1
      },
      "http_types": {
        "security4": 1
      }
    },
    "discovery_types": {
      "zen": 1
    },
    "packaging_types": [
      {
        "flavor": "default",
        "type": "deb",
        "count": 1
      }
    ],
    "ingest": {
      "number_of_pipelines": 1,
      "processor_stats": {
        "gsub": {
          "count": 0,
          "failed": 0,
          "current": 0,
          "time_in_millis": 0
        },
        "script": {
          "count": 0,
          "failed": 0,
          "current": 0,
          "time_in_millis": 0
        }
      }
    }
  }
}
```

