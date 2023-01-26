# elasticsearch cancel-tasks

Cancel a long-running task

## Description

If a long-running task supports cancellation, it can be cancelled with the cancel-tasks

## Synopsis

`elasticsearch cancel-tasks [--site  <site>] [--cluster  <cluster>] [--task_id  <task_id>] [--actions  <actions>] [--node_id  <node_id>]`

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


#### `task_id` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Task Id to cancel  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--task_id  "oTUltX4IQMOUUVeiohTt8A:124"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `actions` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Used to cancel multiple tasks of same action at a time  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--actions  "cluster:*, *search"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `node_id` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Node IDs or names used to limit tasks to cancel  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--node_id  "node-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

elasticsearch cancel-tasks --task_id  "oTUltX4IQMOUUVeiohTt8A:124" --actions  "cluster:*, *search" --node_id  "node-1"
