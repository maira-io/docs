# elasticsearch open-index

Open closed index

## Description

Opens a closed index. For data streams, this opens any closed backing indices.

## Synopsis

`elasticsearch open-index [--site  <site>] [--cluster  <cluster>] <index>`

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

&nbsp;&nbsp;&nbsp;&nbsp; Indices to open  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students,books"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required, multiple allowed_  



## Examples

Input: 
```
! elasticsearch open-index --site "localhost" --cluster "elastic-1" "test" "test1"
```
Output: 
```
{
  "acknowledged": true,
  "shards_acknowledged": true
}
```

