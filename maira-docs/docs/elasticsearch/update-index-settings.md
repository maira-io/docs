# elasticsearch update-index-settings

update elastic search index settings

## Description

Changes settings for indices

## Synopsis

`elasticsearch update-index-settings [--site  <site>] [--cluster  <cluster>] [<index>] [--settings  <settings>]`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of indices (supports wild card expression)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students,class or _all or log_2099_*"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `settings` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Map of settings to update in indices. Set a property to null to revert it to default setting  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--settings  "{ "index.refresh_interval" : null, "index.number_of_replicas" : 1 }"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch update-index-settings --site "localhost" --cluster "elastic-1" "class" "test" --settings `{"index.number_of_replicas": 1}`
```
Output: 
```
{
  "acknowledged": true
}
```

