package decorator

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type commandLoggingDecorator[C any] struct {
	base   CommandHandler[C]
	logger *logrus.Entry
}

func (d commandLoggingDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	handlerName := generateActionName(cmd)

	logger := d.logger.WithFields(logrus.Fields{
		"command":      handlerName,
		"command_body": fmt.Sprintf("%#v", cmd),
	})

	logger.Debug("executing command")
	defer func() {
		if err == nil {
			logger.Info("command executed successfully")
		} else {
			logger.WithError(err).Error("command failed")
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryLoggingDecorator[Q any, R any] struct {
	base   QueryHandler[Q, R]
	logger *logrus.Entry
}

func (d queryLoggingDecorator[Q, R]) Handle(ctx context.Context, q Q) (r R, err error) {
	handlerName := generateActionName(q)

	logger := d.logger.WithFields(logrus.Fields{
		"query":      handlerName,
		"query_body": fmt.Sprintf("%#v", q),
	})

	logger.Debug("executing query")
	defer func() {
		if err == nil {
			logger.Info("query executed successfully")
		} else {
			logger.WithError(err).Error("query failed")
		}
	}()

	return d.base.Handle(ctx, q)
}
