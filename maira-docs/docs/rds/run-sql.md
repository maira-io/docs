# rds run-sql

Run SQL Query

## Description

run sql query in aws rds mysql instance

## Synopsis

`rds run-sql [--site  <site>] [--region  <region>] [--host  <host>] [--database  <database>] <query>`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `region` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; AWS region for the service  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--region  "region-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `host` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Host name for database instance  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--host  "host-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `database` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Database name  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--database  "database-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `query` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Sql Query  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"select * from accounts"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  



## Examples

Input: 
```
! rds run-sql --site "localhost:8081" --host "host1" "select * from movies.movies"
```
Output: 
```
title   director        genre                   release_year    
Joker   Todd Phillips   psychological thriller          2019
Movie1  Karan Johar     psychological thriller          2020 
```

