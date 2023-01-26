# slack send-message

Send Slack Message

## Description

Send a message in slack channel

## Synopsis

`slack send-message --name  <name> --message  <message>`

## Arguments


#### `name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; User/Channel name to post the message  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--name  "@Slack User/#slack-channel"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `message` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Message to post in the channel  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--message  "Build has been successful"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
x=! slack send-message --name "@user1" --message "test message"
```
Output: 
```
Message successfully sent to @user1
```

