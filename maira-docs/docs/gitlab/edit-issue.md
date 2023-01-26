# gitlab edit-issue

Edit gitlab issue

## Description

Edit gitlab issue

## Synopsis

`gitlab edit-issue [--site  <site>] [--title  <title>] --assignee  <assignee> [--confidential ] [--created_at  <created_at>] [--description  <description>] [--due_date  <due_date>] [--epic_id  <epic_id>] --project  <project> --issue  <issue> [--issue_type  <issue_type>] [--labels  <labels>] [--remove_labels  <remove_labels>] [--state_event  <state_event>] [--weight  <weight>] [--updated_at  <updated_at>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `title` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Title of an issue  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--title  "issue1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `assignee` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The username of the user to assign the issue to.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--assignee  "1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `confidential` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Set an issue to be confidential. Default is false.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--confidential  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `created_at` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; When the issue was created. Date time string, ISO 8601 formatted, for example 2016-03-11T03:45:40Z. Requires administrator or project/group owner rights.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--created_at  "2016-03-11T03:45:40Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `description` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The description of an issue. Limited to 1,048,576 characters.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--description  "issue1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `due_date` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The due date. Date time string in the format YYYY-MM-DD  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--due_date  "2016-03-11"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `epic_id` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; ID of the epic to add the issue to. Valid values are greater than or equal to 0.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--epic_id  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `project` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The path of the project  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--project  "saiprakash_maira/testing"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `issue` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; The internal ID of a project’s issue  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--issue  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `issue_type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The type of issue. One of issue, incident, or test_case. Default is issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--issue_type  "issue"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `labels` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Label names for an issue  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--labels  "label"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `remove_labels` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Label names to remove from an issue.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--remove_labels  "label"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `state_event` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; The state event of an issue. Set close to close the issue and reopen to reopen it  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--state_event  "reopen"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `weight` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; The weight of the issue. Valid values are greater than or equal to 0.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--weight  10`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `updated_at` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; When the issue was updated.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--updated_at  "2016-03-11T03:45:40Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! gitlab edit-issue --site "localhost" --project "saiprakash_maira/testing" --issue 4 --description "hello new edit"
```
Output: 
```
{
  "id": 97399269,
  "iid": 4,
  "project_id": 31319463,
  "title": "issue3",
  "description": "hello new edit",
  "state": "opened",
  "created_at": "2016-03-11T03:45:40.000Z",
  "updated_at": "2021-11-16T12:35:51.270Z",
  "closed_at": null,
  "closed_by": null,
  "labels": [ "new_label" ],
  "milestone": null,
  "assignees": [
    {
      "id": 10215385,
      "name": "Sai Prakash V",
      "username": "saiprakash_maira",
      "state": "active",
      "avatar_url": "https://secure.gravatar.com/avatar/21de8fa15da66206e835a51bf9b68c32?s=80\u0026d=identicon",
      "web_url": "https://gitlab.com/saiprakash_maira"
    }
  ],
  "author": {
    "id": 10215385,
    "name": "Sai Prakash V",
    "username": "saiprakash_maira",
    "state": "active",
    "avatar_url": "https://secure.gravatar.com/avatar/21de8fa15da66206e835a51bf9b68c32?s=80\u0026d=identicon",
    "web_url": "https://gitlab.com/saiprakash_maira"
  },
  "type": "ISSUE",
  "assignee": {
    "id": 10215385,
    "name": "Sai Prakash V",
    "username": "saiprakash_maira",
    "state": "active",
    "avatar_url": "https://secure.gravatar.com/avatar/21de8fa15da66206e835a51bf9b68c32?s=80\u0026d=identicon",
    "web_url": "https://gitlab.com/saiprakash_maira"
  },
  "user_notes_count": 2,
  "merge_requests_count": 0,
  "upvotes": 0,
  "downvotes": 0,
  "due_date": "2021-12-12",
  "confidential": false,
  "discussion_locked": false,
  "issue_type": "issue",
  "web_url": "https://gitlab.com/saiprakash_maira/testing/-/issues/4",
  "time_stats": {
    "time_estimate": 0,
    "total_time_spent": 0,
    "human_time_estimate": null,
    "human_total_time_spent": null
  },
  "task_completion_status": {
    "count": 0,
    "completed_count": 0
  },
  "weight": 10,
  "blocking_issues_count": 0,
  "has_tasks": false,
  "_links": {
    "self": "https://gitlab.com/api/v4/projects/31319463/issues/4",
    "notes": "https://gitlab.com/api/v4/projects/31319463/issues/4/notes",
    "award_emoji": "https://gitlab.com/api/v4/projects/31319463/issues/4/award_emoji",
    "project": "https://gitlab.com/api/v4/projects/31319463"
  },
  "references": {
    "short": "#4",
    "relative": "#4",
    "full": "saiprakash_maira/testing#4"
  },
  "subscribed": true,
  "moved_to_id": null,
  "service_desk_reply_to": null,
  "health_status": null
}
```

