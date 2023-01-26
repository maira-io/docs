# elasticsearch get-tasks

Get tasks currently executing in the cluster.

## Description

Returns information about the tasks currently executing on one or more nodes in the cluster.

## Synopsis

`elasticsearch get-tasks [--site  <site>] [--cluster  <cluster>] [--filter_path  <filter_path>] [--raw ] [--task  <task>] [--actions  <actions>] [--detailed ] [--node  <node>]`

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


#### `task` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Task Id  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--task  "oTUltX4IQMOUUVeiohTt8A:124"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `actions` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Used to limit the list of actions returned  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--actions  "cluster:*, *search"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `detailed` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If set, response returns detailed information of the task  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--detailed  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `node` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Node IDs or names used to limit returned information.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--node  "node-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

elasticsearch get-tasks --filter_path  "index.mappings.properties.*" --raw   --task  "oTUltX4IQMOUUVeiohTt8A:124" --actions  "cluster:*, *search" --detailed   --node  "node-1"
