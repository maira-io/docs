# gcp get-log-names

Get log names

## Description

get names for gcp logs

## Synopsis

`gcp get-log-names [--site  <site>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x=! gcp get-log-names --site "localhost:8081"
```
Output: 
```
[                                                                                                                                                                  
  "projects/macro-context-293714/logs/GCEGuestAgent",
  "projects/macro-context-293714/logs/OSConfigAgent",
  "projects/macro-context-293714/logs/apache_access"
]
```

