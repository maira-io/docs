# kafka describe-topics

Describe kafka topics

## Description

Describing topics gives more information like leader of partition, in sync replicas etc

## Synopsis

`kafka describe-topics [--site  <site>] [--cluster  <cluster>] [<topics>]`

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


#### `topics` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; topic name that need to be fetched from cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"topics-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! kafka describe-topics --site "localhost" "test-topic-3"
```
Output: 
```
[
    {
        "name": "test-topic-3",
        "partitions": [
            {
                "leader": 2,
                "replicas": [2],
                "isr": [2]
            }
        ]
    }
]
```

