# confirm

Confirm with an authorized user

## Description

Confirm with an authorized user if we should proceed with the next action in the playbook

## Synopsis

`confirm <message> [--approver  <approver>]`

## Arguments


#### `message` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; message to be sent for confirmation  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"message-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `approver` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; user, group or group:role who needs to approve  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--approver  "kafka:admin"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `default:admin`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

confirm "message-1"
