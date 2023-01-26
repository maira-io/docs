# elasticsearch get-document

get document from index

## Description

Retrieves the specified JSON document from an index

## Synopsis

`elasticsearch get-document [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] --index  <index> <id>`

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


#### `filter_path` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Path in the response. Filters and returns only parts of response matching this path  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--filter_path  "index.mappings.properties.*"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the index that contains the document.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--index  "students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `id` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Unique identifier of the document.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! elasticsearch get-document --site "localhost" --cluster "elastic-1" --index "my-index-000001" 1
```
Output: 
```
{
  "_index": "my-index-000001",
  "_type": "_doc",
  "_id": "1",
  "_version": 1,
  "_seq_no": 0,
  "_primary_term": 1,
  "found": true,
  "_source": {
    "@timestamp": "2099-11-15T13:12:00",
    "message": "GET /search HTTP/1.1 200 1070000",
    "user": {
      "id": "kimchy"
    }
  }
}
```

