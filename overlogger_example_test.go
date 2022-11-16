package zeromdc_test

import (
	"os"

	"github.com/adrienaury/zeromdc"
	"github.com/rs/zerolog"
)

func ExampleNew() {
	zeromdc.New(zerolog.New(os.Stdout))
	zeromdc.MDC().Set("x-correlation-id", "1234")
	zeromdc.MDC().Set("x-agent-name", "trendyol")
	zeromdc.SetGlobalFields([]string{"x-correlation-id", "x-agent-name"})

	zeromdc.Log().Info().Msg("hello world")

	// Output: {"level":"info","x-agent-name":"trendyol","x-correlation-id":"1234","message":"hello world"}
}

func ExampleLog_trace() {
	zeromdc.New(zerolog.New(os.Stdout))
	zeromdc.MDC().Set("x-correlation-id", "1234")
	zeromdc.AddGlobalFields("x-correlation-id")

	zeromdc.Log().Trace().Msg("hello world")

	// Output: {"level":"trace","x-correlation-id":"1234","message":"hello world"}
}

func ExampleLog_debug() {
	zeromdc.New(zerolog.New(os.Stdout))
	zeromdc.MDC().Set("x-correlation-id", "1234")
	zeromdc.AddGlobalFields("x-correlation-id")

	zeromdc.Log().Debug().Msg("hello world")

	// Output: {"level":"debug","x-correlation-id":"1234","message":"hello world"}
}

func ExampleLog_info() {
	zeromdc.New(zerolog.New(os.Stdout))
	zeromdc.MDC().Set("x-correlation-id", "1234")
	zeromdc.AddGlobalFields("x-correlation-id")

	zeromdc.Log().Info().Msg("hello world")

	// Output: {"level":"info","x-correlation-id":"1234","message":"hello world"}
}

func ExampleLog_warn() {
	zeromdc.New(zerolog.New(os.Stdout))
	zeromdc.MDC().Set("x-correlation-id", "1234")
	zeromdc.AddGlobalFields("x-correlation-id")

	zeromdc.Log().Warn().Msg("hello world")

	// Output: {"level":"warn","x-correlation-id":"1234","message":"hello world"}
}

func ExampleLog_error() {
	zeromdc.New(zerolog.New(os.Stdout))
	zeromdc.MDC().Set("x-correlation-id", "1234")
	zeromdc.AddGlobalFields("x-correlation-id")

	zeromdc.Log().Error().Msg("hello world")

	// Output: {"level":"error","x-correlation-id":"1234","message":"hello world"}
}

func ExampleLog_log() {
	zeromdc.New(zerolog.New(os.Stdout))
	zeromdc.MDC().Set("x-correlation-id", "1234")
	zeromdc.AddGlobalFields("x-correlation-id")

	zeromdc.Log().Log().Msg("hello world")

	// Output: {"x-correlation-id":"1234","message":"hello world"}
}

func ExampleLog_print() {
	zeromdc.New(zerolog.New(os.Stdout))
	zeromdc.MDC().Set("x-correlation-id", "1234")
	zeromdc.AddGlobalFields("x-correlation-id")

	zeromdc.Log().Print("hello world")

	// Output: {"level":"debug","x-correlation-id":"1234","message":"hello world"}
}

func ExampleLog_printf() {
	zeromdc.New(zerolog.New(os.Stdout))
	zeromdc.MDC().Set("x-correlation-id", "1234")
	zeromdc.AddGlobalFields("x-correlation-id")

	zeromdc.Log().Printf("hello %s", "world")

	// Output: {"level":"debug","x-correlation-id":"1234","message":"hello world"}
}
