# elasticsearch shrink-index

Shrink an existing index

## Description

Shrinks an existing index into a new index with fewer primary shards

## Synopsis

`elasticsearch shrink-index [--site  <site>] [--cluster  <cluster>] <index> --target_index  <target_index> [--shards  <shards>] [--replicas  <replicas>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the index to be shrunk  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `target_index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name for the shrunk index. This shouldn't already exist in the cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--target_index  "students-shrunk"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `shards` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Primary shard count in target index. Must be lesser and a factor of source index primary shards  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--shards  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `1`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `replicas` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Replica count in target index  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--replicas  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `1`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x = ! elasticsearch shrink-index "testingindex" --target_index "shrunktestingindex" --site "localhost" --shards 1
```
Output: 
```
{
  "acknowledged": true,
  "shards_acknowledged": true,
  "index": "shrunktestingindex"
}
```

