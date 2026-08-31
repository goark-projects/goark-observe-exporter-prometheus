# Goark Observe Exporter Prometheus

`goark.dev/observe-exporter-prometheus` 负责把 Goark 指标导出到 Prometheus。

本仓库是 Prometheus 适配层。它负责 metric 注册、collector、histogram 和 scrape 集成，同时让 `goark.dev/observe` 不依赖 Prometheus。

## 模块路径

```bash
go get goark.dev/observe-exporter-prometheus
```

GitHub 仓库：

```text
github.com/goark-projects/goark-observe-exporter-prometheus
```

## 职责

- 将 Goark metric 契约桥接到 Prometheus collector。
- 提供面向生产的 histogram bucket 默认值。
- 为 Goark Boot 集成暴露 scrape handler。
- 强制使用有界 label 和稳定 metric 名称。

## 边界

本模块只导出指标。它不创建 HTTP、ORM 或 Gnalloy 观测点，并且不能接受 user ID、request ID、raw URL、完整 SQL、SQL 参数、完整错误消息等无界 label。

## 许可证

Apache-2.0
