//go:build clickhouse
// +build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/dannyzhou2015/migrate/v4/database/clickhouse"
)
