# elasticsearch nodes-info

Get cluster nodes information

## Description

Retrieves one or more (or all) of the cluster nodes information.

## Synopsis

`elasticsearch nodes-info [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] [--raw ] [<nodes>] [--metric  <metric>]`

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


#### `nodes` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Nodes  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"node1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `metric` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Limits the information returned to the specific metrics  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--metric  "http, ingest, jvm"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! elasticsearch nodes-info --site "localhost"
```
Output: 
```
[
  {
    "ip": "127.0.0.1",
    "heap.percent": "55",
    "ram.percent": "94",
    "cpu": "19",
    "load_1m": "0.75",
    "load_5m": "1.06",
    "load_15m": "1.35",
    "node.role": "cdfhilmrstw",
    "master": "*",
    "name": "beehyv-H81M-S"
  }
]
```

