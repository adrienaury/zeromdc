package zeromdc

import "github.com/rs/zerolog"

type MDCHook struct{}

func (m MDCHook) Run(event *zerolog.Event, level zerolog.Level, message string) {
	if len(_globalFields) == 0 {
		return
	}

	fields := make(map[string]interface{})
	for _, field := range _globalFields {
		fields[field] = MDC().GetString(field)
	}

	event.Fields(fields)
}
