# elasticsearch update-cluster-settings

update elastic search cluster settings

## Description

Update cluster-wide settings. Settings can be persistent(apply across restarts) or transient(don’t survive a full cluster restart)

## Synopsis

`elasticsearch update-cluster-settings [--site  <site>] [--cluster  <cluster>] [--transient  <transient>] [--persistent  <persistent>]`

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


#### `transient` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Transient settings to be updated  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--transient  "{"indices.recovery.max_bytes_per_sec" : "20mb"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `persistent` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Persistent settings to be updated  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--persistent  "{"indices.recovery.max_bytes_per_sec" : "20mb"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch update-cluster-settings --site "localhost" --cluster "elastic-1" --persistent `{"indices.recovery.max_bytes_per_sec" : "20mb"}`
```
Output: 
```
{
  "acknowledged": true,
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

