# gitlab list-issues

List gitlab issue

## Description

List gitlab issue

## Synopsis

`gitlab list-issues [--site  <site>] [--assignee  <assignee>] [--author  <author>] [--confidential ] [--created_after  <created_after>] [--created_before  <created_before>] [--due_date  <due_date>] [--issues  <issues>] [--issue_type  <issue_type>] [--iteration_title  <iteration_title>] [--labels  <labels>] [--non_archived ] [--order_by  <order_by>] [--scope  <scope>] [--search  <search>] [--sort  <sort>] [--state  <state>] [--updated_after  <updated_after>] [--updated_before  <updated_before>] [--weight  <weight>] [--with_labels_details ]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `assignee` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues assigned to the given username. Similar to assignee_id and mutually exclusive with assignee_id. In GitLab CE, the assignee_username array should only contain a single value. Otherwise, an invalid parameter error is returned.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--assignee  "myname"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `author` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues created by the given username. Similar to author_id and mutually exclusive with author_id.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--author  "author_name"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `confidential` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Filter confidential or public issues.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--confidential  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `created_after` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues created on or after the given time. Expected in ISO 8601 format  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--created_after  "2019-03-15T08:00:00Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `created_before` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues created on or before the given time. Expected in ISO 8601 format  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--created_before  "2019-03-15T08:00:00Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `due_date` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues that have no due date, are overdue, or whose due date is this week, this month, or between two weeks ago and next month. Accepts: 0 (no due date), overdue, week, month, next_month_and_previous_two_weeks.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--due_date  "overdue"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `issues` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Return only the issues having the given iid  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--issues  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `issue_type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Filter to a given type of issue. One of issue, incident, or test_case.  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--issue_type  "test_case"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `iteration_title` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues assigned to the iteration with the given title. Similar to iteration_id and mutually exclusive with iteration_id.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--iteration_title  "title"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `labels` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Label names, issues must have all labels to be returned. None lists all issues with no labels. Any lists all issues with at least one label. Predefined names are case-insensitive.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--labels  "label"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `non_archived` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues only from non-archived projects. If false, the response returns issues from both archived and non-archived projects. Default is true.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--non_archived  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `order_by` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues ordered by created_at, due_date, label_priority, milestone_due, popularity, priority, relative_position, title, updated_at, or weight fields. Default is created_at.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--order_by  "created_at"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `scope` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues for the given scope: created_by_me, assigned_to_me or all. Defaults to created_by_me For versions before 11.0, use the now deprecated created-by-me or assigned-to-me scopes instead.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--scope  "created_by_me"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `search` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Search issues against their title and description  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--search  "title"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `sort` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues sorted in asc or desc order. Default is desc  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--sort  "asc"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `state` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Return all issues or just those that are opened or closed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--state  "opened"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `updated_after` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues updated on or after the given time. Expected in ISO 8601 format
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--updated_after  "2019-03-15T08:00:00Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `updated_before` - (time)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues updated on or before the given time. Expected in ISO 8601 format
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--updated_before  "2019-03-15T08:00:00Z"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `weight` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; Return issues with the specified weight. None returns issues with no weight assigned. Any returns issues with a weight assigned.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--weight  10`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `with_labels_details` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; If true, the response returns more details for each label in labels field: :name, :color, :description, :description_html, :text_color. Default is false.
  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--with_labels_details  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x = ! gitlab list-issues --site "localhost" --labels "new_label"
```
Output: 
```
[
  {
    "id": 97399269,
    "iid": 4,
    "project_id": 31319463,
    "title": "issue3",
    "description": "hello new edit",
    "state": "opened",
    "created_at": "2016-03-11T03:45:40.000Z",
    "updated_at": "2021-11-16T12:52:11.470Z",
    "closed_at": null,
    "closed_by": null,
    "labels": ["new_label"],
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
    "user_notes_count": 3,
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
    "moved_to_id": null,
    "service_desk_reply_to": null,
    "health_status": null
  }
]
```

