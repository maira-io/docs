# kafka get-broker-config

Get Config for kafka broker Resource

## Description

Get one or more config for a given kafka broker resource

## Synopsis

`kafka get-broker-config [--site  <site>] [--cluster  <cluster>] --broker_id  <broker_id> [<config_names>]`

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


#### `broker_id` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; broker id  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `--broker_id  "broker_id-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _required_  


#### `config_names` - (string)

&nbsp;&nbsp;&nbsp;&nbsp; Name of config. If not provided, returns all the config of the resource  

&nbsp;&nbsp;&nbsp;&nbsp; Example:  `"config_names-1"`

&nbsp;&nbsp;&nbsp;&nbsp; Default: `_None_`
&nbsp;&nbsp;&nbsp;&nbsp; Attributes: _optional, multiple allowed_  



## Examples

Input: 
```
x = ! kafka get-brokerresource-config --broker_id "1" --site "localhost"
```
Output: 
```
{
"advertised.listeners": {
  "default": true,
  "source": "Default"
},
"alter.config.policy.class.name": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"alter.log.dirs.replication.quota.window.num": {
  "value": "11",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"alter.log.dirs.replication.quota.window.size.seconds": {
  "value": "1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"authorizer.class.name": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"auto.create.topics.enable": {
  "value": "true",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"auto.leader.rebalance.enable": {
  "value": "true",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"background.threads": {
  "value": "10",
  "default": true,
  "source": "Default"
},
"broker.heartbeat.interval.ms": {
  "value": "2000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"broker.id": {
  "value": "1",
  "read_only": true,
  "source": "StaticBroker"
},
"broker.id.generation.enable": {
  "value": "true",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"broker.rack": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"broker.session.timeout.ms": {
  "value": "9000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"client.quota.callback.class": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"compression.type": {
  "value": "producer",
  "default": true,
  "source": "Default"
},
"connection.failed.authentication.delay.ms": {
  "value": "100",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"connections.max.idle.ms": {
  "value": "600000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"connections.max.reauth.ms": {
  "value": "0",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"control.plane.listener.name": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controlled.shutdown.enable": {
  "value": "true",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controlled.shutdown.max.retries": {
  "value": "3",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controlled.shutdown.retry.backoff.ms": {
  "value": "5000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.listener.names": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.quorum.append.linger.ms": {
  "value": "25",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.quorum.election.backoff.max.ms": {
  "value": "1000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.quorum.election.timeout.ms": {
  "value": "1000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.quorum.fetch.timeout.ms": {
  "value": "2000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.quorum.request.timeout.ms": {
  "value": "2000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.quorum.retry.backoff.ms": {
  "value": "20",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.quorum.voters": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.quota.window.num": {
  "value": "11",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.quota.window.size.seconds": {
  "value": "1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"controller.socket.timeout.ms": {
  "value": "30000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"create.topic.policy.class.name": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"default.replication.factor": {
  "value": "1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"delegation.token.expiry.check.interval.ms": {
  "value": "3600000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"delegation.token.expiry.time.ms": {
  "value": "86400000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"delegation.token.master.key": {
  "read_only": true,
  "default": true,
  "source": "Default",
  "sensitive": true
},
"delegation.token.max.lifetime.ms": {
  "value": "604800000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"delegation.token.secret.key": {
  "read_only": true,
  "default": true,
  "source": "Default",
  "sensitive": true
},
"delete.records.purgatory.purge.interval.requests": {
  "value": "1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"delete.topic.enable": {
  "value": "true",
  "read_only": true,
  "source": "StaticBroker"
},
"fetch.max.bytes": {
  "value": "57671680",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"fetch.purgatory.purge.interval.requests": {
  "value": "1000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"group.initial.rebalance.delay.ms": {
  "value": "0",
  "read_only": true,
  "source": "StaticBroker"
},
"group.max.session.timeout.ms": {
  "value": "1800000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"group.max.size": {
  "value": "2147483647",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"group.min.session.timeout.ms": {
  "value": "6000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"initial.broker.registration.timeout.ms": {
  "value": "60000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"inter.broker.listener.name": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"inter.broker.protocol.version": {
  "value": "3.0-IV1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"kafka.metrics.polling.interval.secs": {
  "value": "10",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"kafka.metrics.reporters": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"leader.imbalance.check.interval.seconds": {
  "value": "300",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"leader.imbalance.per.broker.percentage": {
  "value": "10",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"listener.security.protocol.map": {
  "value": "PLAINTEXT:PLAINTEXT,SSL:SSL,SASL_PLAINTEXT:SASL_PLAINTEXT,SASL_SSL:SASL_SSL",
  "default": true,
  "source": "Default"
},
"listeners": {
  "value": "PLAINTEXT://:9093",
  "source": "StaticBroker"
},
"log.cleaner.backoff.ms": {
  "value": "15000",
  "default": true,
  "source": "Default"
},
"log.cleaner.dedupe.buffer.size": {
  "value": "134217728",
  "default": true,
  "source": "Default"
},
"log.cleaner.delete.retention.ms": {
  "value": "86400000",
  "default": true,
  "source": "Default"
},
"log.cleaner.enable": {
  "value": "true",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"log.cleaner.io.buffer.load.factor": {
  "value": "0.9",
  "default": true,
  "source": "Default"
},
"log.cleaner.io.buffer.size": {
  "value": "524288",
  "default": true,
  "source": "Default"
},
"log.cleaner.io.max.bytes.per.second": {
  "value": "1.7976931348623157E308",
  "default": true,
  "source": "Default"
},
"log.cleaner.max.compaction.lag.ms": {
  "value": "9223372036854775807",
  "default": true,
  "source": "Default"
},
"log.cleaner.min.cleanable.ratio": {
  "value": "0.5",
  "default": true,
  "source": "Default"
},
"log.cleaner.min.compaction.lag.ms": {
  "value": "0",
  "default": true,
  "source": "Default"
},
"log.cleaner.threads": {
  "value": "1",
  "default": true,
  "source": "Default"
},
"log.cleanup.policy": {
  "value": "delete",
  "default": true,
  "source": "Default"
},
"log.dir": {
  "value": "/tmp/kafka-logs",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"log.dirs": {
  "value": "/home/beehyv/kafka/logs1",
  "read_only": true,
  "source": "StaticBroker"
},
"log.flush.interval.messages": {
  "value": "9223372036854775807",
  "default": true,
  "source": "Default"
},
"log.flush.interval.ms": {
  "default": true,
  "source": "Default"
},
"log.flush.offset.checkpoint.interval.ms": {
  "value": "60000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"log.flush.scheduler.interval.ms": {
  "value": "9223372036854775807",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"log.flush.start.offset.checkpoint.interval.ms": {
  "value": "60000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"log.index.interval.bytes": {
  "value": "4096",
  "default": true,
  "source": "Default"
},
"log.index.size.max.bytes": {
  "value": "10485760",
  "default": true,
  "source": "Default"
},
"log.message.downconversion.enable": {
  "value": "true",
  "default": true,
  "source": "Default"
},
"log.message.format.version": {
  "value": "3.0-IV1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"log.message.timestamp.difference.max.ms": {
  "value": "9223372036854775807",
  "default": true,
  "source": "Default"
},
"log.message.timestamp.type": {
  "value": "CreateTime",
  "default": true,
  "source": "Default"
},
"log.preallocate": {
  "value": "false",
  "default": true,
  "source": "Default"
},
"log.retention.bytes": {
  "value": "-1",
  "default": true,
  "source": "Default"
},
"log.retention.check.interval.ms": {
  "value": "300000",
  "read_only": true,
  "source": "StaticBroker"
},
"log.retention.hours": {
  "value": "168",
  "read_only": true,
  "source": "StaticBroker"
},
"log.retention.minutes": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"log.retention.ms": {
  "default": true,
  "source": "Default"
},
"log.roll.hours": {
  "value": "168",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"log.roll.jitter.hours": {
  "value": "0",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"log.roll.jitter.ms": {
  "default": true,
  "source": "Default"
},
"log.roll.ms": {
  "default": true,
  "source": "Default"
},
"log.segment.bytes": {
  "value": "1073741824",
  "source": "StaticBroker"
},
"log.segment.delete.delay.ms": {
  "value": "60000",
  "default": true,
  "source": "Default"
},
"max.connection.creation.rate": {
  "value": "2147483647",
  "default": true,
  "source": "Default"
},
"max.connections": {
  "value": "2147483647",
  "default": true,
  "source": "Default"
},
"max.connections.per.ip": {
  "value": "2147483647",
  "default": true,
  "source": "Default"
},
"max.connections.per.ip.overrides": {
  "default": true,
  "source": "Default"
},
"max.incremental.fetch.session.cache.slots": {
  "value": "1000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"message.max.bytes": {
  "value": "1048588",
  "default": true,
  "source": "Default"
},
"metadata.log.dir": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"metadata.log.max.record.bytes.between.snapshots": {
  "value": "20971520",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"metadata.log.segment.bytes": {
  "value": "1073741824",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"metadata.log.segment.ms": {
  "value": "604800000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"metadata.max.retention.bytes": {
  "value": "-1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"metadata.max.retention.ms": {
  "value": "604800000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"metric.reporters": {
  "default": true,
  "source": "Default"
},
"metrics.num.samples": {
  "value": "2",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"metrics.recording.level": {
  "value": "INFO",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"metrics.sample.window.ms": {
  "value": "30000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"min.insync.replicas": {
  "value": "1",
  "default": true,
  "source": "Default"
},
"node.id": {
  "value": "-1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"num.io.threads": {
  "value": "8",
  "source": "StaticBroker"
},
"num.network.threads": {
  "value": "3",
  "source": "StaticBroker"
},
"num.partitions": {
  "value": "1",
  "read_only": true,
  "source": "StaticBroker"
},
"num.recovery.threads.per.data.dir": {
  "value": "1",
  "source": "StaticBroker"
},
"num.replica.alter.log.dirs.threads": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"num.replica.fetchers": {
  "value": "1",
  "default": true,
  "source": "Default"
},
"offset.metadata.max.bytes": {
  "value": "4096",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"offsets.commit.required.acks": {
  "value": "-1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"offsets.commit.timeout.ms": {
  "value": "5000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"offsets.load.buffer.size": {
  "value": "5242880",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"offsets.retention.check.interval.ms": {
  "value": "600000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"offsets.retention.minutes": {
  "value": "10080",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"offsets.topic.compression.codec": {
  "value": "0",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"offsets.topic.num.partitions": {
  "value": "50",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"offsets.topic.replication.factor": {
  "value": "1",
  "read_only": true,
  "source": "StaticBroker"
},
"offsets.topic.segment.bytes": {
  "value": "104857600",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"password.encoder.cipher.algorithm": {
  "value": "AES/CBC/PKCS5Padding",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"password.encoder.iterations": {
  "value": "4096",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"password.encoder.key.length": {
  "value": "128",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"password.encoder.keyfactory.algorithm": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"password.encoder.old.secret": {
  "read_only": true,
  "default": true,
  "source": "Default",
  "sensitive": true
},
"password.encoder.secret": {
  "read_only": true,
  "default": true,
  "source": "Default",
  "sensitive": true
},
"principal.builder.class": {
  "value": "org.apache.kafka.common.security.authenticator.DefaultKafkaPrincipalBuilder",
  "default": true,
  "source": "Default"
},
"process.roles": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"producer.purgatory.purge.interval.requests": {
  "value": "1000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"queued.max.request.bytes": {
  "value": "-1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"queued.max.requests": {
  "value": "500",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"quota.window.num": {
  "value": "11",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"quota.window.size.seconds": {
  "value": "1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.fetch.backoff.ms": {
  "value": "1000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.fetch.max.bytes": {
  "value": "1048576",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.fetch.min.bytes": {
  "value": "1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.fetch.response.max.bytes": {
  "value": "10485760",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.fetch.wait.max.ms": {
  "value": "500",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.high.watermark.checkpoint.interval.ms": {
  "value": "5000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.lag.time.max.ms": {
  "value": "30000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.selector.class": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.socket.receive.buffer.bytes": {
  "value": "65536",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replica.socket.timeout.ms": {
  "value": "30000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replication.quota.window.num": {
  "value": "11",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"replication.quota.window.size.seconds": {
  "value": "1",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"request.timeout.ms": {
  "value": "30000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"reserved.broker.max.id": {
  "value": "1000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"sasl.client.callback.handler.class": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"sasl.enabled.mechanisms": {
  "value": "GSSAPI",
  "default": true,
  "source": "Default"
},
"sasl.jaas.config": {
  "default": true,
  "source": "Default",
  "sensitive": true
},
"sasl.kerberos.kinit.cmd": {
  "value": "/usr/bin/kinit",
  "default": true,
  "source": "Default"
},
"sasl.kerberos.min.time.before.relogin": {
  "value": "60000",
  "default": true,
  "source": "Default"
},
"sasl.kerberos.principal.to.local.rules": {
  "value": "DEFAULT",
  "default": true,
  "source": "Default"
},
"sasl.kerberos.service.name": {
  "default": true,
  "source": "Default"
},
"sasl.kerberos.ticket.renew.jitter": {
  "value": "0.05",
  "default": true,
  "source": "Default"
},
"sasl.kerberos.ticket.renew.window.factor": {
  "value": "0.8",
  "default": true,
  "source": "Default"
},
"sasl.login.callback.handler.class": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"sasl.login.class": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"sasl.login.refresh.buffer.seconds": {
  "value": "300",
  "default": true,
  "source": "Default"
},
"sasl.login.refresh.min.period.seconds": {
  "value": "60",
  "default": true,
  "source": "Default"
},
"sasl.login.refresh.window.factor": {
  "value": "0.8",
  "default": true,
  "source": "Default"
},
"sasl.login.refresh.window.jitter": {
  "value": "0.05",
  "default": true,
  "source": "Default"
},
"sasl.mechanism.controller.protocol": {
  "value": "GSSAPI",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"sasl.mechanism.inter.broker.protocol": {
  "value": "GSSAPI",
  "default": true,
  "source": "Default"
},
"sasl.server.callback.handler.class": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"security.inter.broker.protocol": {
  "value": "PLAINTEXT",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"security.providers": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"socket.connection.setup.timeout.max.ms": {
  "value": "30000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"socket.connection.setup.timeout.ms": {
  "value": "10000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"socket.receive.buffer.bytes": {
  "value": "102400",
  "read_only": true,
  "source": "StaticBroker"
},
"socket.request.max.bytes": {
  "value": "104857600",
  "read_only": true,
  "source": "StaticBroker"
},
"socket.send.buffer.bytes": {
  "value": "102400",
  "read_only": true,
  "source": "StaticBroker"
},
"ssl.cipher.suites": {
  "default": true,
  "source": "Default"
},
"ssl.client.auth": {
  "value": "none",
  "default": true,
  "source": "Default"
},
"ssl.enabled.protocols": {
  "value": "TLSv1.2,TLSv1.3",
  "default": true,
  "source": "Default"
},
"ssl.endpoint.identification.algorithm": {
  "value": "https",
  "default": true,
  "source": "Default"
},
"ssl.engine.factory.class": {
  "default": true,
  "source": "Default"
},
"ssl.key.password": {
  "default": true,
  "source": "Default",
  "sensitive": true
},
"ssl.keymanager.algorithm": {
  "value": "SunX509",
  "default": true,
  "source": "Default"
},
"ssl.keystore.certificate.chain": {
  "default": true,
  "source": "Default",
  "sensitive": true
},
"ssl.keystore.key": {
  "default": true,
  "source": "Default",
  "sensitive": true
},
"ssl.keystore.location": {
  "default": true,
  "source": "Default"
},
"ssl.keystore.password": {
  "default": true,
  "source": "Default",
  "sensitive": true
},
"ssl.keystore.type": {
  "value": "JKS",
  "default": true,
  "source": "Default"
},
"ssl.principal.mapping.rules": {
  "value": "DEFAULT",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"ssl.protocol": {
  "value": "TLSv1.3",
  "default": true,
  "source": "Default"
},
"ssl.provider": {
  "default": true,
  "source": "Default"
},
"ssl.secure.random.implementation": {
  "default": true,
  "source": "Default"
},
"ssl.trustmanager.algorithm": {
  "value": "PKIX",
  "default": true,
  "source": "Default"
},
"ssl.truststore.certificates": {
  "default": true,
  "source": "Default",
  "sensitive": true
},
"ssl.truststore.location": {
  "default": true,
  "source": "Default"
},
"ssl.truststore.password": {
  "default": true,
  "source": "Default",
  "sensitive": true
},
"ssl.truststore.type": {
  "value": "JKS",
  "default": true,
  "source": "Default"
},
"transaction.abort.timed.out.transaction.cleanup.interval.ms": {
  "value": "10000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"transaction.max.timeout.ms": {
  "value": "900000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"transaction.remove.expired.transaction.cleanup.interval.ms": {
  "value": "3600000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"transaction.state.log.load.buffer.size": {
  "value": "5242880",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"transaction.state.log.min.isr": {
  "value": "1",
  "read_only": true,
  "source": "StaticBroker"
},
"transaction.state.log.num.partitions": {
  "value": "50",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"transaction.state.log.replication.factor": {
  "value": "1",
  "read_only": true,
  "source": "StaticBroker"
},
"transaction.state.log.segment.bytes": {
  "value": "104857600",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"transactional.id.expiration.ms": {
  "value": "604800000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"unclean.leader.election.enable": {
  "value": "false",
  "default": true,
  "source": "Default"
},
"zookeeper.clientCnxnSocket": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.connect": {
  "value": "localhost:2181",
  "read_only": true,
  "source": "StaticBroker"
},
"zookeeper.connection.timeout.ms": {
  "value": "18000",
  "read_only": true,
  "source": "StaticBroker"
},
"zookeeper.max.in.flight.requests": {
  "value": "10",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.session.timeout.ms": {
  "value": "18000",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.set.acl": {
  "value": "false",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.cipher.suites": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.client.enable": {
  "value": "false",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.crl.enable": {
  "value": "false",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.enabled.protocols": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.endpoint.identification.algorithm": {
  "value": "HTTPS",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.keystore.location": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.keystore.password": {
  "read_only": true,
  "default": true,
  "source": "Default",
  "sensitive": true
},
"zookeeper.ssl.keystore.type": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.ocsp.enable": {
  "value": "false",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.protocol": {
  "value": "TLSv1.2",
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.truststore.location": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.ssl.truststore.password": {
  "read_only": true,
  "default": true,
  "source": "Default",
  "sensitive": true
},
"zookeeper.ssl.truststore.type": {
  "read_only": true,
  "default": true,
  "source": "Default"
},
"zookeeper.sync.time.ms": {
  "value": "2000",
  "read_only": true,
  "default": true,
  "source": "Default"
}
```

