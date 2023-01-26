# kafka get-brokers

get kafka brokers

## Description

Get brokers in the cluster

## Synopsis

`kafka get-brokers [--site  <site>] [--cluster  <cluster>] [<broker_ids>]`

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


#### `broker_ids` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; ids for the brokers that are to be fetched  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! kafka get-brokers --site "localhost" 1 2
```
Output: 
```
[
  {
    "id": 1,
    "addr": "site-H81M-S:9093"
  },
  {
    "id": 2,
    "addr": "site-H81M-S:9094"
  }
]
```

