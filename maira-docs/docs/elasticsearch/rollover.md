# elasticsearch rollover

Rollover an index alias or data stream

## Description

Rollover an index alias or data stream to create a new index

## Synopsis

`elasticsearch rollover [--site  <site>] [--cluster  <cluster>] <target> [--target_index  <target_index>] [--max_age  <max_age>] [--max_docs  <max_docs>] [--max_size  <max_size>] [--max_primary_shard_size  <max_primary_shard_size>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the rollover target index alias or data stream  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `target_index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name for the new index  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--target_index  "index-01"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `max_age` - (duration)

&nbsp;&nbsp;&nbsp;&nbsp; rollover after the maximum elapsed time from index creation is reached  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--max_age  "5d"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `max_docs` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; rollover after the specified maximum number of documents is reached  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--max_docs  500`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `max_size` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; rollover when the index reaches a certain size.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--max_size  "5mb"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `max_primary_shard_size` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; rollover when the largest primary shard in the index reaches a certain size  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--max_primary_shard_size  "5mb"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

elasticsearch rollover "students" --target_index  "index-01" --max_age  "5d" --max_docs  500 --max_size  "5mb" --max_primary_shard_size  "5mb"
