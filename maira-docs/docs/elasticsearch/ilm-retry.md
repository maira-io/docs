# elasticsearch ilm-retry

Retry executing the policy for an index in ERROR step

## Description

Set the policy back to the step where the error occurred and execute the step. Use ilm-explain to determine if an index is in the ERROR step.

## Synopsis

`elasticsearch ilm-retry [--site  <site>] [--cluster  <cluster>] --index  <index>`

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

&nbsp;&nbsp;&nbsp;&nbsp; Name of the index  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--index  "students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! elasticsearch ilm-retry --site "localhost" --cluster "elastic-1" --index "test"
```
Output: 
```
{
  "acknowledged": true
}
```

