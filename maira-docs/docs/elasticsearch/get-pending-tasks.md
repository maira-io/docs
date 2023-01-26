# elasticsearch get-pending-tasks

Get elastic search cluster-level changes that have not yet been executed

## Description

Get elastic search cluster-level changes that have not yet been executed

## Synopsis

`elasticsearch get-pending-tasks [--site  <site>] [--cluster  <cluster>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example: `--site  site-1`  
&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`  
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of elastic search cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example: `--cluster  elastic-1`  
&nbsp;&nbsp;&nbsp;&nbsp; Default: `default`  
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch get-pending-tasks --site "localhost" --cluster "elastic-1"
```
Output: 
```
{
  "tasks": []
}

```
