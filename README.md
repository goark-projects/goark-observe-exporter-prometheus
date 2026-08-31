# Goark Observe Exporter Prometheus

[中文说明](README.zh-CN.md)

`goark.dev/observe-exporter-prometheus` exports Goark metrics to Prometheus.

This repository is the Prometheus adapter layer. It owns metric registration, collectors, histograms, and scrape integration while keeping Prometheus dependencies out of `goark.dev/observe`.

## Module

```bash
go get goark.dev/observe-exporter-prometheus
```

GitHub repository:

```text
github.com/goark-projects/goark-observe-exporter-prometheus
```

## Scope

- Bridge Goark metric contracts to Prometheus collectors.
- Provide production-oriented histogram bucket defaults.
- Expose scrape handlers for Goark Boot integration.
- Enforce bounded labels and stable metric names.

## Boundary

This module exports metrics. It does not create HTTP, ORM, or Gnalloy observation points, and it must not accept unbounded labels such as user IDs, request IDs, raw URLs, full SQL, SQL arguments, or full error messages.

## License

Apache-2.0
