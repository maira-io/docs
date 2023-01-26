# elasticsearch create-index

Create a new elastic search index

## Description

Create a new elastic search index

## Synopsis

`elasticsearch create-index [--site  <site>] [--cluster  <cluster>] <index> [--shards  <shards>] [--replicas  <replicas>]`

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


#### `index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the index to be created  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `shards` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Number of primary shards that an index should have  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--shards  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `1`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `replicas` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Number of replicas  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--replicas  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `1`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch create-index --site "localhost" --cluster "elastic-1" --shards 4 --replicas 1 "my_source_index"
```
Output: 
```
{
  "acknowledged": true,
  "shards_acknowledged": true,
  "index": "class"
}
```

