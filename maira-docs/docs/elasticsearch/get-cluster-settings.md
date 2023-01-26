# elasticsearch get-cluster-settings

Get elastic search cluster wide settings

## Description

Get elastic search cluster wide settings

## Synopsis

`elasticsearch get-cluster-settings [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] [--include_defaults ]`

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


#### `include_defaults` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Parameter to decide whether default settings are to be retrieved  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--include_defaults  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch get-cluster-settings --site "localhost" --cluster "elastic-1"
```
Output: 
```
{
  "persistent": {
    "indices": {
      "recovery": {
        "max_bytes_per_sec": "20mb"
      }
    }
  },
  "transient": {}
}
```

