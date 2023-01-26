# elasticsearch search

Search with a given query

## Description

Returns search hits that match the query in the request

## Synopsis

`elasticsearch search [--site  <site>] [--cluster  <cluster>] [--target  <target>] [--from  <from>] [--size  <size>] [--sort  <sort>] --query  <query>`

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


#### `target` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of data streams, indices, and index aliases used to limit the request  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--target  "students or _all"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `from` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Starting document offset  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--from  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `0`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `size` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Defines the number of hits to return  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--size  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `10`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `sort` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; List of <field>:<direction> pairs  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--sort  ""name":"asc","age":"desc""`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `query` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Query string you wish to parse and use for filtering. Defined using Lucene query string syntax.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--query  "(new york city) OR (big apple)"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! elasticsearch search --site "localhost" --cluster "elastic-1" --target "my-index-000001" --query "user.id:kimchy"
```
Output: 
```
{
  "took": 1,
  "timed_out": false,
  "_shards": {
    "total": 1,
    "successful": 1,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": {
      "value": 3,
      "relation": "eq"
    },
    "max_score": 0.6931471,
    "hits": [
      {
        "_index": "my-index-000001",
        "_type": "_doc",
        "_id": "5qCjO3wBAue9X_A2Q4qK",
        "_score": 0.6931471,
        "_source": {
          "@timestamp": "2099-11-15T13:12:00",
          "message": "GET /search HTTP/1.1 200 1070000",
          "user": {
            "id": "kimchy"
          }
        }
      },
      {
        "_index": "my-index-000001",
        "_type": "_doc",
        "_id": "o74aUXwBYiwTRml6JRUV",
        "_score": 0.6931471,
        "_source": {
          "@timestamp": "2099-11-15T13:12:00",
          "message": "GET /search HTTP/1.1 200 1070000",
          "user": {
            "id": "kimchy"
          }
        }
      },
      {
        "_index": "my-index-000001",
        "_type": "_doc",
        "_id": "3",
        "_score": 0.6931471,
        "_source": {
          "@timestamp": "2099-11-15T13:12:00",
          "message": "GET /search HTTP/1.1 200 1070000",
          "user": {
            "id": "kimchy"
          }
        }
      }
    ]
  }
}
```

