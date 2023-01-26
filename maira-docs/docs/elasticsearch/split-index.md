# elasticsearch split-index

Split an existing index

## Description

Splits an existing index into a new index with more primary shards

## Synopsis

`elasticsearch split-index [--site  <site>] [--cluster  <cluster>] <index> --target_index  <target_index> --shards  <shards>`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the index to be split  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `target_index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name for the split index. This shouldn't already exist in the cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--target_index  "students-split"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `shards` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Primary shard count in target index. It must be higher and a multiple of source index primary shards  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--shards  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! elasticsearch split-index "my_source_index" --site "localhost" --shards 2 --target_index "split_index_1"
```
Output: 
```
{
  "acknowledged": true,
  "shards_acknowledged": true,
  "index": "splittestingindex"
}
```

