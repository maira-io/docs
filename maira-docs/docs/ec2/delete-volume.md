# ec2 delete-volume

Delete an ec2 volume

## Description

Delete an ec2 volume

## Synopsis

`ec2 delete-volume [--site  <site>] [--region  <region>] <volume_id> --region  <region> [--dryrun ]`

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


#### `volume_id` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Volume ID to be deleted  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"volume_id-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `region` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Region of the volume to be deleted  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--region  "region-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `dryrun` - (bool)

&nbsp;&nbsp;&nbsp;&nbsp; Do a dryrun of the delete or actually delete  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--dryrun  `

&nbsp;&nbsp;&nbsp;&nbsp; Default: `false`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  



## Examples

ec2 delete-volume --region  "region-1" "volume_id-1" --region  "region-1"
