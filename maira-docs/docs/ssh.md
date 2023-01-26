# ssh

Run a command on a remote server using ssh

## Description

Run a command on a remote server using ssh

## Synopsis

`ssh [--site  <site>] <server> <command> [--insecure_ignore_host_key ]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site to run the command  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `server` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; server to run the command on  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"server1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `command` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; command to run  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"ls -l"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `insecure_ignore_host_key` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; ignore host key verification (insecure)  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--insecure_ignore_host_key  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

ssh "server1" "ls -l" --insecure_ignore_host_key  
