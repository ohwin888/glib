package database

import (
	"github.com/hanyougame/glib/tracing"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

func registerTraceHook(tx *gorm.DB) {
	tx.Callback().Create().Before("gorm:created").Register("trace:create", func(db *gorm.DB) {
		traceSql("gorm:create", tx)
	})
	tx.Callback().Create().After("gorm:saved").Register("trace:save", func(db *gorm.DB) {
		traceSql("gorm:save", tx)
	})
	tx.Callback().Query().After("gorm:queried").Register("trace:query", func(db *gorm.DB) {
		traceSql("gorm:query", tx)
	})
	tx.Callback().Delete().After("gorm:deleted").Register("trace:delete", func(db *gorm.DB) {
		traceSql("gorm:delete", tx)
	})
	tx.Callback().Update().After("gorm:updated").Register("trace:update", func(db *gorm.DB) {
		traceSql("gorm:update", tx)
	})
	tx.Callback().Raw().After("*").Register("trace:raw", func(db *gorm.DB) {
		traceSql("gorm:raw", tx)
	})
	tx.Callback().Row().After("*").Register("trace:row", func(db *gorm.DB) {
		traceSql("gorm:row", tx)
	})
}

func traceSql(spanName string, db *gorm.DB) {
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	tracing.Inject(db.Statement.Context, spanName, func(span oteltrace.Span) oteltrace.Span {
		span.SetAttributes(attribute.KeyValue{
			Key:   "gorm.sql",
			Value: attribute.StringValue(sql),
		})
		return span
	})
}
