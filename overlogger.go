package zeromdc

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var overLogger *Overlog //nolint:gochecknoglobals

type Overlog struct {
	*zerolog.Logger
}

func New(logger zerolog.Logger) *Overlog {
	log.Logger = logger.Hook(MDCHook{})

	ResetGlobalMdcAdapter()
	ClearGlobalFields()

	overLogger = &Overlog{
		Logger: &log.Logger,
	}

	return overLogger
}

func Log() *Overlog {
	return overLogger
}
