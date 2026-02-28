// Package jobs defines interfaces for asynchronous job processing in RoadRunner.
// It provides the Driver interface for queue backends (e.g. AMQP, Kafka, Beanstalk),
// the Pipeline interface for configuring named pipelines, and the Message and Job
// interfaces for producing and consuming queue items. The State struct exposes
// runtime counters such as active, delayed, and reserved job counts.
package jobs
