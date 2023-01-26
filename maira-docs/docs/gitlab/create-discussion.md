# gitlab create-discussion

Create new discussion for gitlab issue

## Description

Create new discussion for gitlab issue

## Synopsis

`gitlab create-discussion [--site  <site>] --project  <project> --issue  <issue> <body> [--created_at  <created_at>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `project` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The path of the project  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--project  "saiprakash_maira/testing"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `issue` - (integer)

&nbsp;&nbsp;&nbsp;&nbsp; The IID of an issue  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--issue  "1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `body` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The content of the thread  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"discussion_here"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `created_at` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; Date time string, ISO 8601 formatted  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--created_at  "2016-03-11T03:45:40Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x = ! gitlab create-discussion --site "localhost" --project "saiprakash_maira/testing" --issue 4 "hello discussion 2"
```
Output: 
```
{
  "id": "d92e0d28ff5fe5a407effb55aa6fe977b66d9e77",
  "individual_note": false,
  "notes": [
    {
      "id": 734515843,
      "type": "DiscussionNote",
      "body": "hello discussion 2",
      "attachment": null,
      "author": {
        "id": 10215385,
        "name": "Sai Prakash V",
        "username": "saiprakash_maira",
        "state": "active",
        "avatar_url": "https://secure.gravatar.com/avatar/21de8fa15da66206e835a51bf9b68c32?s=80\u0026d=identicon",
        "web_url": "https://gitlab.com/saiprakash_maira"
      },
      "created_at": "2021-11-16T12:52:11.405Z",
      "updated_at": "2021-11-16T12:52:11.405Z",
      "system": false,
      "noteable_id": 97399269,
      "noteable_type": "Issue",
      "resolvable": false,
      "confidential": false,
      "noteable_iid": 4,
      "commands_changes": {}
    }
  ]
}
```

