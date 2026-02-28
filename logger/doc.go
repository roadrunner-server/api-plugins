// Package logger defines logging interfaces for RoadRunner plugins.
// The Log interface wraps go.uber.org/zap, exposing leveled methods from Debug
// through Fatal. The Named interface provides NamedLogger for obtaining a
// logger scoped to a specific plugin or component name.
package logger
