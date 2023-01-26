# kafka alter-resource-config

Alter config of kafka resources

## Description

Alters config of a given kafka resource

## Synopsis

`kafka alter-resource-config [--site  <site>] [--cluster  <cluster>] --resource_name  <resource_name> --resource_type  <resource_type> <entries> [--validate_only ]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of kafka cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "cluster-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `resource_name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Can be topic name or broker id  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--resource_name  "resource_name-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `resource_type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Values can either be "TopicResource" or "BrokerResource"  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--resource_type  "resource_type-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `entries` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; stringified Map with keys as config names and values as modified config values  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"entries-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `validate_only` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp;   

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--validate_only  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! kafka alter-resource-config --site "localhost" --resource_type "TopicResource" --resource_name "test-topic-1" `{"file.delete.delay.ms":"10000", "delete.retention.ms":"1000000"}`
```
Output: 
```
Config for the resource test-topic-1 updated
```

