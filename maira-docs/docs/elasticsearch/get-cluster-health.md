# elasticsearch get-cluster-health

Get elastic search cluster health

## Description

Get the health status of a cluster

## Synopsis

`elasticsearch get-cluster-health [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] [--raw ] [--target  <target>] [--level  <level>] [--local ]`

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


#### `raw` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If false, returns compact and aligned text output  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--raw  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `target` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of data streams, indices, and index aliases used to limit the request  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--target  "students or _all"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `level` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Can be either cluster/indices/shards. Controls the details level of the health information returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--level  "cluster"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `cluster`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `local` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, the request retrieves information from the local node only. By default information is retrieved from the master node.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--local  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch get-cluster-health --site "localhost"
```
Output: 
```
[
  {
    "epoch": "1639500455",
    "timestamp": "16:47:35",
    "cluster": "default",
    "status": "red",
    "node.total": "1",
    "node.data": "1",
    "shards": "20",
    "pri": "20",
    "relo": "0",
    "init": "0",
    "unassign": "81",
    "pending_tasks": "0",
    "max_task_wait_time": "-",
    "active_shards_percent": "19.8%"
  }
]
```

