# gcp get-log-resources

Get resource descriptors

## Description

get resource descriptors for gcp logs

## Synopsis

`gcp get-log-resources [--site  <site>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
x=! gcp get-log-resources --site "localhost:8081"
```
Output: 
```
[
  {
    "type": "aiplatform.googleapis.com/Endpoint",
    "display_name": "Vertex AI Endpoint",
    "description": "A Vertex AI API Endpoint where Models are deployed into it.",
    "labels": [
      {
        "key": "resource_container",
        "description": "The identifier of the GCP Project owning the Endpoint."
      },
      {
        "key": "location",
        "description": "The region in which the service is running."
      },
      {
        "key": "endpoint_id",
        "description": "The ID of the Endpoint."
      }
    ],
    "launch_stage": 3
  }
]
```

