# icinga query-configuration-objects

Returns information about configuration objects

## Description

Returns information about configuration objects

## Synopsis

`icinga query-configuration-objects [--site  <site>] [--cluster  <cluster>] [--type  <type>] [--name  <name>] [--attrs  <attrs>] [--joins  <joins>] [--meta  <meta>] [--filters  <filters>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of Icinga cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "icinga-default"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `icinga-default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `type` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Type of configuration object. Used to limit the objects returned in the response to a certain type of object  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--type  "hosts/ services"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `name` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Object name to query for a single object  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--name  "example.localdomain"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `attrs` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Attribute list of object to limit output  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--attrs  "name/ address"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `joins` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Join related object types and their attributes specified as list  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--joins  "host"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `meta` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Enable meta information  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--meta  "meta=used_by (references from other objects), meta=location (location information) specified as list"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  


#### `filters` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Used to return objects that only match the provided criteria  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--filters  "service.acknowledgement_expiry==0, service.downtime_depth == 0.0"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! icinga query-configuration-objects --site "localhost" --type "hosts"
```
Output: 
```
{
  "results": [
    {
      "attrs": {
        "__name": "beehyv-H81M-S",
        "acknowledgement": 0,
        "acknowledgement_expiry": 0,
        "acknowledgement_last_change": 0,
        "action_url": "",
        "active": true,
        "address": "127.0.0.1",
        "address6": "::1",
        "check_attempt": 1,
        "check_command": "hostalive",
        "check_interval": 60,
        "check_period": "",
        "check_timeout": null,
        "command_endpoint": "",
        "display_name": "beehyv-H81M-S",
        "downtime_depth": 0,
        "enable_active_checks": true,
        "enable_event_handler": true,
        "enable_flapping": false,
        "enable_notifications": true,
        "enable_passive_checks": true,
        "enable_perfdata": true,
        "event_command": "",
        "executions": null,
        "flapping": false,
        "flapping_current": 0,
        "flapping_ignore_states": null,
        "flapping_last_change": 0,
        "flapping_threshold": 0,
        "flapping_threshold_high": 30,
        "flapping_threshold_low": 25,
        "force_next_check": false,
        "force_next_notification": false,
        "groups": ["linux-servers"],
        "ha_mode": 0,
        "handled": false,
        "icon_image": "",
        "icon_image_alt": "",
        "last_check": 1635938707.104217,
        "last_check_result": {
          "active": true,
          "check_source": "beehyv-H81M-S",
          "command": [
            "/usr/lib/nagios/plugins/check_ping",
            "-H",
            "127.0.0.1",
            "-c",
            "5000,100%",
            "-w",
            "3000,80%"
          ],
          "execution_end": 1635938707.104121,
          "execution_start": 1635938707.100854,
          "exit_status": 128,
          "output": "execvpe(/usr/lib/nagios/plugins/check_ping) failed: No such file or directory",
          "performance_data": [],
          "schedule_end": 1635938707.104217,
          "schedule_start": 1635938707.1000001,
          "scheduling_source": "beehyv-H81M-S",
          "state": 3,
          "ttl": 0,
          "type": "CheckResult",
          "vars_after": {
            "attempt": 1,
            "reachable": true,
            "state": 3,
            "state_type": 1
          },
          "vars_before": {
            "attempt": 1,
            "reachable": true,
            "state": 3,
            "state_type": 1
          }
        },
        "last_hard_state": 1,
        "last_hard_state_change": 1635416489.6736,
        "last_reachable": true,
        "last_state": 1,
        "last_state_change": 0,
        "last_state_down": 0,
        "last_state_type": 1,
        "last_state_unreachable": 0,
        "last_state_up": 0,
        "max_check_attempts": 3,
        "name": "beehyv-H81M-S",
        "next_check": 1635938767.1,
        "next_update": 1635938827.1082416,
        "notes": "",
        "notes_url": "",
        "original_attributes": null,
        "package": "_etc",
        "paused": false,
        "previous_state_change": 0,
        "problem": true,
        "retry_interval": 30,
        "severity": 2112,
        "source_location": {
          "first_column": 1,
          "first_line": 18,
          "last_column": 20,
          "last_line": 18,
          "path": "/etc/icinga2/conf.d/hosts.conf"
        },
        "state": 1,
        "state_type": 1,
        "templates": ["beehyv-H81M-S", "generic-host"],
        "type": "Host",
        "vars": {
          "disks": {
            "disk": {},
            "disk /": {
              "disk_partitions": "/"
            }
          },
          "http_vhosts": {
            "http": {
              "http_uri": "/"
            }
          },
          "notification": {
            "mail": {
              "groups": ["icingaadmins"]
            }
          },
          "os": "Linux"
        },
        "version": 0,
        "volatile": false,
        "zone": ""
      },
      "joins": {},
      "meta": {},
      "name": "beehyv-H81M-S",
      "type": "Host"
    }
  ]
}
```

