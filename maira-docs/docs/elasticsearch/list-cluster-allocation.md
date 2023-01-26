# elasticsearch list-cluster-allocation

Get compact and aligned text explanation for shards

## Description

Get compact and aligned text explanation for shards current allocation.

## Synopsis

`elasticsearch list-cluster-allocation [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] [--nodes  <nodes>]`

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



## Examples

Input: 
```
! elasticsearch list-cluster-allocation --site "localhost"
```
Output: 
```
[
  {
    "shards": "81",
    "disk.indices": null,
    "disk.used": null,
    "disk.avail": null,
    "disk.total": null,
    "disk.percent": null,
    "host": null,
    "ip": null,
    "node": "UNASSIGNED"
  }
]
```

