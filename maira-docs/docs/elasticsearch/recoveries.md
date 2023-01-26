# elasticsearch recoveries

Get shard recoveries

## Description

Get information about ongoing and completed shard recoveries

## Synopsis

`elasticsearch recoveries [--site  <site>] [--cluster  <cluster>] [--target  <target>] [--filter_path  <filter_path>] [--sort  <sort>] [--active_only ] [--detailed ]`

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


#### `sort` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of column names or column aliases used to sort the response.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--sort  "index,health"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `active_only` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, response only includes ongoing shard recoveries.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--active_only  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `detailed` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, response only includes detailed info of shard recoveries.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--detailed  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch recoveries --site "localhost" --cluster "elastic-1" --target "my-index-000001" --target "class" --detailed --sort "index" --sort "files"
```
Output: 
```
[
  {
    "index": "class",
    "shard": "0",
    "time": "859ms",
    "type": "existing_store",
    "stage": "done",
    "source_host": "n/a",
    "source_node": "n/a",
    "target_host": "localhost",
    "target_node": "beehyv-H81M-S",
    "repository": "n/a",
    "snapshot": "n/a",
    "files": "0",
    "files_recovered": "0",
    "files_percent": "100.0%",
    "files_total": "1",
    "bytes": "0",
    "bytes_recovered": "0",
    "bytes_percent": "100.0%",
    "bytes_total": "208",
    "translog_ops": "0",
    "translog_ops_recovered": "0",
    "translog_ops_percent": "100.0%"
  },
  {
    "index": "class",
    "shard": "1",
    "time": "669ms",
    "type": "existing_store",
    "stage": "done",
    "source_host": "n/a",
    "source_node": "n/a",
    "target_host": "localhost",
    "target_node": "beehyv-H81M-S",
    "repository": "n/a",
    "snapshot": "n/a",
    "files": "0",
    "files_recovered": "0",
    "files_percent": "100.0%",
    "files_total": "1",
    "bytes": "0",
    "bytes_recovered": "0",
    "bytes_percent": "100.0%",
    "bytes_total": "208",
    "translog_ops": "0",
    "translog_ops_recovered": "0",
    "translog_ops_percent": "100.0%"
  },
  {
    "index": "my-index-000001",
    "shard": "0",
    "time": "2.1s",
    "type": "existing_store",
    "stage": "done",
    "source_host": "n/a",
    "source_node": "n/a",
    "target_host": "localhost",
    "target_node": "beehyv-H81M-S",
    "repository": "n/a",
    "snapshot": "n/a",
    "files": "0",
    "files_recovered": "0",
    "files_percent": "100.0%",
    "files_total": "13",
    "bytes": "0",
    "bytes_recovered": "0",
    "bytes_percent": "100.0%",
    "bytes_total": "15934",
    "translog_ops": "0",
    "translog_ops_recovered": "0",
    "translog_ops_percent": "100.0%"
  }
]
```

