package factory

import (
	"context"

	"github.com/go-xorm/xorm"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	ContextDBName     = "DB"
	ContextLoggerName = "logger"
)

func DB(ctx context.Context) *xorm.Session {
	v := ctx.Value(ContextDBName)
	if v == nil {
		panic("DB is not exist")
	}
	if db, ok := v.(*xorm.Session); ok {
		return db
	}
	panic("DB is not exist")
}

func Logger(ctx context.Context) *logrus.Entry {
	v := ctx.Value(ContextLoggerName)
	if v == nil {
		return logrus.WithFields(logrus.Fields{})
	}
	if logger, ok := v.(*logrus.Entry); ok {
		return logger
	}
	return logrus.WithFields(logrus.Fields{})
}

func Tracer(ctx context.Context) opentracing.Span {
	if s := opentracing.SpanFromContext(ctx); s != nil {
		return s
	}
	return opentracing.NoopTracer{}.StartSpan("")
}
