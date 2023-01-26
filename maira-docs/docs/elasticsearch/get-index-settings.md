# elasticsearch get-index-settings

Get elastic search index settings

## Description

Get settings info for one or more indices

## Synopsis

`elasticsearch get-index-settings [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] <index> [--settings  <settings>]`

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


#### `index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of indices (also supports wild card expression)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students,class / _all / log_2099_*"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required, multiple allowed_  


#### `settings` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List or wildcard expression of setting names to filter the request  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--settings  "index.number_*"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! elasticsearch get-index-settings --site "localhost" --cluster "elastic-1" "my_source_index" --settings "index.number_of_shards" --settings "index.number_of_replicas"
```
Output: 
```
{
  "my_source_index": {
    "settings": {
      "index": {
        "number_of_shards": "4",
        "number_of_replicas": "0"
      }
    }
  }
}
```

