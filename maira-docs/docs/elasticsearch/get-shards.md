# elasticsearch get-shards

Get detailed view of what nodes contain which shards

## Description

Get detailed view of what nodes contain which shards along with info whether it’s a primary or replica, the number of docs, the bytes it takes on disk, and the node where it’s located.

## Synopsis

`elasticsearch get-shards [--site  <site>] [--cluster  <cluster>] [--target  <target>] [--filter_path  <filter_path>] [--sort  <sort>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Column names or column aliases used to sort the response.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--sort  "index,health"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `index`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
x = ! elasticsearch get-shards --site "localhost" --cluster "elastic-1" --target "test" --target "class"
```
Output: 
```
[
  {
    "index": "class",
    "shard": "1",
    "prirep": "p",
    "state": "STARTED",
    "docs": "0",
    "store": "208b",
    "ip": "127.0.0.1",
    "node": "beehyv-H81M-S"
  },
  {
    "index": "class",
    "shard": "1",
    "prirep": "r",
    "state": "UNASSIGNED",
    "docs": null,
    "store": null,
    "ip": null,
    "node": null
  },
  {
    "index": "class",
    "shard": "0",
    "prirep": "p",
    "state": "STARTED",
    "docs": "0",
    "store": "208b",
    "ip": "127.0.0.1",
    "node": "beehyv-H81M-S"
  },
  {
    "index": "class",
    "shard": "0",
    "prirep": "r",
    "state": "UNASSIGNED",
    "docs": null,
    "store": null,
    "ip": null,
    "node": null
  },
  {
    "index": "test",
    "shard": "0",
    "prirep": "p",
    "state": "STARTED",
    "docs": "0",
    "store": "230b",
    "ip": "127.0.0.1",
    "node": "beehyv-H81M-S"
  },
  {
    "index": "test",
    "shard": "0",
    "prirep": "r",
    "state": "UNASSIGNED",
    "docs": null,
    "store": null,
    "ip": null,
    "node": null
  }
]
```

