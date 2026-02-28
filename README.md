<a href="https://roadrunner.dev" target="_blank">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://github.com/roadrunner-server/.github/assets/8040338/e6bde856-4ec6-4a52-bd5b-bfe78736c1ff">
    <img align="center" src="https://github.com/roadrunner-server/.github/assets/8040338/040fb694-1dd3-4865-9d29-8e0748c2c8b8">
  </picture>
</a>

# API Plugins

[![GoDoc](https://pkg.go.dev/badge/github.com/roadrunner-server/api-plugins/v5)](https://pkg.go.dev/github.com/roadrunner-server/api-plugins/v5)
[![Go Report Card](https://goreportcard.com/badge/github.com/roadrunner-server/api-plugins/v5)](https://goreportcard.com/report/github.com/roadrunner-server/api-plugins/v5)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/roadrunner-server/api-plugins/blob/master/LICENSE)
[![Discord](https://img.shields.io/badge/discord-chat-magenta.svg)](https://discord.gg/TFeEmCs)

This repository defines Go interface contracts for [RoadRunner](https://roadrunner.dev) server plugins. Plugins implement these interfaces so that the [Endure](https://github.com/roadrunner-server/endure) dependency injection framework can automatically discover and wire dependencies via its plugin graph. Endure inspects each plugin's `Init()` method parameters, matches them to registered plugins that implement the required interfaces, and topologically sorts the resulting DAG to determine initialization order.

## Packages

| Package          | Description                                                                                                                 |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `jobs`           | Job queue interfaces — `Driver`, `Constructor`, `Pipeline`, `Message`, `Job`, `Queue`, and `State` for async job processing |
| `kv`             | Key-value storage — `Storage`, `Item`, and `Constructor` interfaces with TTL support                                        |
| `lock`           | Lock queue — `Item` and `Queue` interfaces for priority-based lock management                                               |
| `logger`         | Logging — `Log` and `Named` interfaces wrapping `go.uber.org/zap`                                                           |
| `priority_queue` | Base priority queue — `Item` interface with `ID`, `GroupID`, `Priority`                                                     |
| `status`         | Status reporting — `Status` struct with `Code` field                                                                        |

## Usage

A plugin declares its dependencies by accepting interfaces in its `Init()` method. Endure auto-resolves them from the plugin graph:

```go
type Plugin struct{}

func (p *Plugin) Init(log logger.Log, kv kv.Storage) error {
    log.Info("plugin initialized")
    return nil
}
```

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
