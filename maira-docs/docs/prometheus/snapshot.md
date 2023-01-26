# prometheus snapshot

Create a snapshot of all current data

## Description

Creates a snapshot of all current data into snapshots/<datetime>-<rand> under the TSDB's data directory and returns the directory as response. It will optionally skip snapshotting data that is only present in the head block, and which has not yet been compacted to disk.

## Synopsis

`prometheus snapshot [--site  <site>] [--cluster  <cluster>] [--skip_head ]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of prometheus cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "prometheus-default"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `prometheus-default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `skip_head` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Skip data present in the head block  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--skip_head  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! prometheus snapshot --site "localhost"
```
Output: 
```
{
  "status": "success",
  "data": {
    "name": "20211021T115101Z-25ffe8c87b4ff6da"
  }
}
```

