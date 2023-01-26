# kafka get-resource-config

Get Config for kafka Resource

## Description

Get one or more config for a given kafka resource

## Synopsis

`kafka get-resource-config [--site  <site>] [--cluster  <cluster>] --resource_name  <resource_name> --resource_type  <resource_type> [<config_names>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example: `--site  "site-1"`  
&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of kafka cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example: `--cluster  "cluster-1"`  
&nbsp;&nbsp;&nbsp;&nbsp; Default: `default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `resource_name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Can be topic name or broker id  

&nbsp;&nbsp;&nbsp;&nbsp; Example: `--resource_name  "resource_name-1"`  
&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `resource_type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Values can either be "TopicResource" or "BrokerResource"  

&nbsp;&nbsp;&nbsp;&nbsp; Example: `--resource_type  "resource_type-1"`  
&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `config_names` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of config. If not provided, returns all the config of the resource  

&nbsp;&nbsp;&nbsp;&nbsp; Example: `"config_names-1"`  
&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! kafka get-resource-config --site "localhost" --resource_type "TopicResource" --resource_name "test-topic-1" "file.delete.delay.ms"
```
Output: 
```
{
  "file.delete.delay.ms": {
    "value": "60000",
    "default": true,
    "source": "Default"
  }
}
```

