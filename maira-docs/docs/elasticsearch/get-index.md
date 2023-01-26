# elasticsearch get-index

Get elastic search index

## Description

Get information about elastic search index

## Synopsis

`elasticsearch get-index [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] [--raw ] [<index>]`

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


#### `raw` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If false, returns compact and aligned text output  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--raw  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `index` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of the index  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"students"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_all`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! elasticsearch get-index --site "localhost"
```
Output: 
```
[
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_17",
    "uuid": "nHfj7WewRE2tQTMA5DTNcA",
    "pri": "2",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "green",
    "status": "open",
    "index": "my_source_index",
    "uuid": "tkkO4G2UQ_KDXfOUgZVuGQ",
    "pri": "4",
    "rep": "0",
    "docs.count": "0",
    "docs.deleted": "0",
    "store.size": "832b",
    "pri.store.size": "832b"
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_18",
    "uuid": "9whKu-_zSretubh_P05hyw",
    "pri": "2",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "yellow",
    "status": "open",
    "index": "test",
    "uuid": "vpk0XgasRcGGTfpvF7R-SQ",
    "pri": "1",
    "rep": "1",
    "docs.count": "0",
    "docs.deleted": "0",
    "store.size": "230b",
    "pri.store.size": "230b"
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_9",
    "uuid": "7j7EmPYKRcufspTWs_OtIw",
    "pri": "1",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "yellow",
    "status": "open",
    "index": "t-test1",
    "uuid": "WosIvQUYR-KmIydtOGUCBg",
    "pri": "1",
    "rep": "1",
    "docs.count": "0",
    "docs.deleted": "0",
    "store.size": "208b",
    "pri.store.size": "208b"
  },
  {
    "health": "red",
    "status": "open",
    "index": "split_index_1",
    "uuid": "h5jhRdxsRQibcXu64EHXoQ",
    "pri": "8",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "red",
    "status": "open",
    "index": "target-test1",
    "uuid": "ETmgcyngQUCKBzEQodk3qQ",
    "pri": "1",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_91",
    "uuid": "AxRVBmPJSYOVkKJNa9_QhA",
    "pri": "2",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "yellow",
    "status": "open",
    "index": "test1",
    "uuid": "07QzkcVMRVORd5zxHYLQ0Q",
    "pri": "3",
    "rep": "3",
    "docs.count": "0",
    "docs.deleted": "0",
    "store.size": "690b",
    "pri.store.size": "690b"
  },
  {
    "health": "green",
    "status": "open",
    "index": ".geoip_databases",
    "uuid": "wbVFSqdfRNKUTjeZ-AXx1A",
    "pri": "1",
    "rep": "0",
    "docs.count": "42",
    "docs.deleted": "3",
    "store.size": "41.1mb",
    "pri.store.size": "41.1mb"
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index",
    "uuid": "2Uduz_34SU6jQuSKQ1PQdQ",
    "pri": "2",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_90",
    "uuid": "ThDt-IeVTH-Ii6_N7M3gKA",
    "pri": "1",
    "rep": "2",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_6",
    "uuid": "-PpVy-DlRHC5pKMjs4tnxw",
    "pri": "2",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_7",
    "uuid": "VW8SG6hoRDe_jQPHbcZYfA",
    "pri": "2",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "yellow",
    "status": "open",
    "index": "sdfa",
    "uuid": "KYONjP3aTTCi9wxsE-EbKw",
    "pri": "1",
    "rep": "2",
    "docs.count": "0",
    "docs.deleted": "0",
    "store.size": "208b",
    "pri.store.size": "208b"
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_4",
    "uuid": "_ZE6FERzRO27o-H9Y9vPbQ",
    "pri": "2",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_2",
    "uuid": "RPp3p_ucSR6NYW4TlijnOQ",
    "pri": "1",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "yellow",
    "status": "open",
    "index": "class",
    "uuid": "HfkB3ftvSNSFf5HNXeKUbA",
    "pri": "2",
    "rep": "1",
    "docs.count": "0",
    "docs.deleted": "0",
    "store.size": "416b",
    "pri.store.size": "416b"
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_3",
    "uuid": "AzqOzJ0wQZ-DUiJ26bFGlw",
    "pri": "2",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  },
  {
    "health": "yellow",
    "status": "open",
    "index": "my-index-000001",
    "uuid": "I340__lCQhmYTtKOQkRYUA",
    "pri": "1",
    "rep": "1",
    "docs.count": "5",
    "docs.deleted": "1",
    "store.size": "20.2kb",
    "pri.store.size": "20.2kb"
  },
  {
    "health": "red",
    "status": "open",
    "index": "my_target_index_1",
    "uuid": "K6oB-VNsTk-2_-kBFYWr1Q",
    "pri": "1",
    "rep": "1",
    "docs.count": null,
    "docs.deleted": null,
    "store.size": null,
    "pri.store.size": null
  }
]
```

