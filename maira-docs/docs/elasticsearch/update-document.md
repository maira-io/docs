# elasticsearch update-document

update elastic search document

## Description

Update a document using the specified script

## Synopsis

`elasticsearch update-document [--site  <site>] [--cluster  <cluster>] --index  <index> <id> [--script  <script>] [--params  <params>]`

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


#### `id` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Unique identifier of the document.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `script` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Script to be executed on document  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--script  "ctx._source.new_field = 'value_of_new_field' or ctx._source.counter += params.count or ctx._source.tags.add(params.tag)"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `params` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Params that are used in source  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--params  "{"count" : 4, "tag": "blue"}"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x = ! elasticsearch update-document --site "localhost" --cluster "elastic-1" --index "my-index-000001" 2 --script "ctx._source.new_field = params.count" --params `{"count" : 4}`
```
Output: 
```
{
  "_index": "my-index-000001",
  "_type": "_doc",
  "_id": "2",
  "_version": 4,
  "result": "updated",
  "_shards": {
  "total": 2,
    "s  uccessful": 1,
    "failed": 0
  },
  "_seq_no": 7,
  "_primary_term": 1
}
```

