# elasticsearch delete-by-query

Delete documents that match the specific query

## Description

Delete documents that match the specific query

## Synopsis

`elasticsearch delete-by-query [--site  <site>] [--cluster  <cluster>] --query  <query> [<index>] [--async ]`

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


#### `query` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Query string you wish to parse and use for filtering. Defined using Lucene query string syntax.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--query  "(new york city) OR (big apple)"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Indices to close  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `async` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Create a task instead of waiting for completion  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--async  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! elasticsearch delete-by-query --site "localhost" --query "user.id:kimchy"
```
Output: 
```
{
  "took": 258,
  "timed_out": false,
  "total": 1,
  "deleted": 1,
  "batches": 1,
  "version_conflicts": 0,
  "noops": 0,
  "retries": {
    "bulk": 0,
    "search": 0
  },
  "throttled_millis": 0,
  "requests_per_second": -1.0,
  "throttled_until_millis": 0,
  "failures": []
}
```

