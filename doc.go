// Package prometheus 实现 Goark 指标到 Prometheus 的导出适配。
//
// 本包负责 collector、histogram 和 scrape 集成，必须保持指标名称稳定，
// 并严格限制 label 基数。
package prometheus
