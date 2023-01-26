# prometheus get-config

Get currently loaded configuration

## Description

Returns currently loaded configuration file

## Synopsis

`prometheus get-config [--site  <site>] [--cluster  <cluster>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of prometheus cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "prometheus-default"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `prometheus-default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

Input: 
```
! prometheus get-config --site "localhost"
```
Output: 
```
{
  "status": "success",
  "data": {
    "yaml": "global:\n  scrape_interval: 15s\n  scrape_timeout: 10s\n  evaluation_interval: 15s\nalerting:\n  alertmanagers:\n  - follow_redirects: true\n    scheme: http\n    timeout: 10s\n    api_version: v2\n    static_configs:\n    - targets:\n      - localhost:9093\nrule_files:\n- /usr/local/bin/prometheus/prometheus_rules.yml\nscrape_configs:\n- job_name: prometheus\n  honor_timestamps: true\n  scrape_interval: 15s\n  scrape_timeout: 10s\n  metrics_path: /metrics\n  scheme: http\n  follow_redirects: true\n  static_configs:\n  - targets:\n    - localhost:9090\n- job_name: node-exporter\n  honor_timestamps: true\n  scrape_interval: 15s\n  scrape_timeout: 10s\n  metrics_path: /metrics\n  scheme: http\n  follow_redirects: true\n  static_configs:\n  - targets:\n    - localhost:9100\n- job_name: kafka\n  honor_timestamps: true\n  scrape_interval: 15s\n  scrape_timeout: 10s\n  metrics_path: /metrics\n  scheme: http\n  follow_redirects: true\n  static_configs:\n  - targets:\n    - localhost:1234\n"
  }
}
```

