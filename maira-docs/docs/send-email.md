# send-email

Send an email

## Description

Send an email with an existing template or with a subject and body

## Synopsis

`send-email --to  <to> [--cc  <cc>] [--subject  <subject>] [--body  <body>] [--template  <template>] [<template_args>]`

## Arguments


#### `to` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Addressses to which email is to be sent  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--to  "to-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required, multiple allowed_  


#### `cc` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Addressses to which email is Cc-ed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cc  "cc-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `subject` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Subject of the email  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--subject  "subject-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `body` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Body of the email  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--body  "body-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `template` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Template to be used for sending the email. Template/args are mutually exclusive with subject/body  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--template  "template-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `template_args` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Template arguments(json) to be used for sending the email  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"template_args-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

send-email --to  "to-1" --cc  "cc-1" --subject  "subject-1" --body  "body-1" --template  "template-1" "template_args-1"
