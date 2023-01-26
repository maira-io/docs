# elasticsearch get-index-segments

Get elastic search index segments

## Description

Get low-level information about the Lucene segments in index shards

## Synopsis

`elasticsearch get-index-segments [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] [--raw ] [--verbose ] [<index>]`

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


#### `verbose` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Get additional debugging information  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--verbose  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the index  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! elasticsearch get-index-segments --site "localhost"
```
Output: 
```
[
  {
    "index": ".geoip_databases",
    "shard": "0",
    "prirep": "p",
    "ip": "127.0.0.1",
    "segment": "_ah",
    "generation": "377",
    "docs.count": "42",
    "docs.deleted": "3",
    "size": "41.1mb",
    "size.memory": "1460",
    "committed": "true",
    "searchable": "true",
    "version": "8.9.0",
    "compound": "false"
  },
  {
    "index": "my-index-000001",
    "shard": "0",
    "prirep": "p",
    "ip": "127.0.0.1",
    "segment": "_3",
    "generation": "3",
    "docs.count": "2",
    "docs.deleted": "1",
    "size": "6.2kb",
    "size.memory": "2092",
    "committed": "true",
    "searchable": "true",
    "version": "8.9.0",
    "compound": "true"
  },
  {
    "index": "my-index-000001",
    "shard": "0",
    "prirep": "p",
    "ip": "127.0.0.1",
    "segment": "_7",
    "generation": "7",
    "docs.count": "1",
    "docs.deleted": "0",
    "size": "4.6kb",
    "size.memory": "1876",
    "committed": "true",
    "searchable": "true",
    "version": "8.9.0",
    "compound": "true"
  },
  {
    "index": "my-index-000001",
    "shard": "0",
    "prirep": "p",
    "ip": "127.0.0.1",
    "segment": "_8",
    "generation": "8",
    "docs.count": "1",
    "docs.deleted": "0",
    "size": "4.6kb",
    "size.memory": "1876",
    "committed": "true",
    "searchable": "true",
    "version": "8.9.0",
    "compound": "true"
  },
  {
    "index": "my-index-000001",
    "shard": "0",
    "prirep": "p",
    "ip": "127.0.0.1",
    "segment": "_9",
    "generation": "9",
    "docs.count": "1",
    "docs.deleted": "0",
    "size": "4.1kb",
    "size.memory": "1876",
    "committed": "true",
    "searchable": "true",
    "version": "8.9.0",
    "compound": "true"
  }
]
```

