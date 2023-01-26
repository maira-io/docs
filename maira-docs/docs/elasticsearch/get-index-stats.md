# elasticsearch get-index-stats

get stats of indices

## Description

Returns statistics for one or more indices. For data streams, the API retrieves statistics for the stream’s backing indices.

## Synopsis

`elasticsearch get-index-stats [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] <index> [--metrics  <metrics>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of indices (supports wild card expression)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students"`
 ,  `"class"`
 ,  `"_all"`
 ,  `"log_2099_*"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required, multiple allowed_  


#### `metrics` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of metrics used to limit the request. Possible values are completion, docs, field data, flush, get, indexing, merge, query_cache, refresh, request_cache, search, segments, store, suggest, translog, warmer  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--metrics  "fielddata,flush,refresh"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! elasticsearch get-index-stats --site "localhost" --cluster "elastic-1" "my-index-000001" "test" --metrics "refresh" --metrics "flush"
```
Output: 
```
{
  "_shards": {
    "total": 4,
    "successful": 2,
    "failed": 0
  },
  "_all": {
    "primaries": {
      "refresh": {
        "total": 7,
        "total_time_in_millis": 30,
        "external_total": 6,
        "external_total_time_in_millis": 31,
        "listeners": 0
      },
      "flush": {
        "total": 3,
        "periodic": 0,
        "total_time_in_millis": 200
      }
    },
    "total": {
      "refresh": {
        "total": 7,
        "total_time_in_millis": 30,
        "external_total": 6,
        "external_total_time_in_millis": 31,
        "listeners": 0
      },
      "flush": {
        "total": 3,
        "periodic": 0,
        "total_time_in_millis": 200
      }
    }
  },
  "indices": {
    "test": {
      "uuid": "vpk0XgasRcGGTfpvF7R-SQ",
      "primaries": {
        "refresh": {
          "total": 2,
          "total_time_in_millis": 0,
          "external_total": 2,
          "external_total_time_in_millis": 0,
          "listeners": 0
        },
        "flush": {
          "total": 1,
          "periodic": 0,
          "total_time_in_millis": 0
        }
      },
      "total": {
        "refresh": {
          "total": 2,
          "total_time_in_millis": 0,
          "external_total": 2,
          "external_total_time_in_millis": 0,
          "listeners": 0
        },
        "flush": {
          "total": 1,
          "periodic": 0,
          "total_time_in_millis": 0
        }
      }
    },
    "my-index-000001": {
      "uuid": "I340__lCQhmYTtKOQkRYUA",
      "primaries": {
        "refresh": {
          "total": 5,
          "total_time_in_millis": 30,
          "external_total": 4,
          "external_total_time_in_millis": 31,
          "listeners": 0
        },
        "flush": {
          "total": 2,
          "periodic": 0,
          "total_time_in_millis": 200
        }
      },
      "total": {
        "refresh": {
          "total": 5,
          "total_time_in_millis": 30,
          "external_total": 4,
          "external_total_time_in_millis": 31,
          "listeners": 0
        },
        "flush": {
          "total": 2,
          "periodic": 0,
          "total_time_in_millis": 200
        }
      }
    }
  }
}
```

