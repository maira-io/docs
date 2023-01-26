# kafka get-broker-metrics

Get Kafka Broker metrics

## Description

Get metrics for one broker or all kafka brokers combined

## Synopsis

`kafka get-broker-metrics [--site  <site>] [--cluster  <cluster>] [--broker  <broker>] [<beans>]`

## Arguments


#### `site` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Site where this command will be executed  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--site  "site-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `input.site`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `cluster` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of kafka cluster  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--cluster  "cluster-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `default`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `broker` - (int)

&nbsp;&nbsp;&nbsp;&nbsp; broker id  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--broker  1`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `-1`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional_  


#### `beans` - ()

&nbsp;&nbsp;&nbsp;&nbsp; list of beans  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"kafka.server"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
! kafka get-broker-metrics --site "localhost" --broker 1
```
Output: 
```
BROKER_ID	LINUX-DISK-READ-BYTES	LINUX-DISK-WRITE-BYTES	ISR_EXPANDS	AT_MIN_ISR_PARTITION_COUNT	PARTITION_COUNT	LEADER_COUNT	PURGATORY_SIZE	TOTAL_FETCH_REQUESTS	REQUEST_HANDLER_AVG_IDLE_PERCENT_MEAN_RATE	HEAP_MEMORY_USED	NON_HEAP_MEMORY_USED	
        1	             22421504	              12103680	          0	                         1	              3	           2	           507	                   1	                                         0	       182999488	            71166816
```

