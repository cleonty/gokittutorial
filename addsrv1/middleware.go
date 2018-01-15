package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMidleware struct {
	logger log.Logger
	next   StringService
}

func (mw loggingMidleware) Uppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.next.Uppercase(ctx, s)
	return
}

func (mw loggingMidleware) Count(ctx context.Context, s string) (output int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"output", output,
			"took", time.Since(begin),
		)
	}(time.Now())
	output = mw.next.Count(ctx, s)
	return
}
