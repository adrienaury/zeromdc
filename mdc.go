package zeromdc

import (
	"strconv"
	"sync"
)

// value, exists := over.MDC().Get("stats")
// over.MDC().Set("stats", &stats{CreatedLinesCount: map[string]int{}, DeletedLinesCount: map[string]int{}})
// over.MDC().Set("action", "pull")
// over.MDC().Set("action", "push")
// over.SetGlobalFields([]string{"action"})
// over.MDC().Set("workerid", id)
// stats, ok := over.MDC().Get("stats")

var (
	_globalMdc  = InitGlobalMdcAdapter()
	_MdcAdapter *MdcAdapter
)

type MdcAdapter struct {
	items map[string]any
	sync.RWMutex
}

func InitGlobalMdcAdapter() *MdcAdapter {
	if _MdcAdapter == nil {
		_MdcAdapter = &MdcAdapter{
			items:   make(map[string]any),
			RWMutex: sync.RWMutex{},
		}
	}

	return _MdcAdapter
}

func MDC() *MdcAdapter {
	_MdcAdapter.RLock()
	s := _globalMdc
	_MdcAdapter.RUnlock()

	return s
}

func (m *MdcAdapter) Get(key string) (any, bool) {
	uniqueKey := m.getUniqueKey(key)
	m.RLock()
	v, ok := m.items[uniqueKey]
	m.RUnlock()
	return v, ok
}

func (m *MdcAdapter) getUniqueKey(key string) string {
	if key == "" {
		panic("MDC key cannot be empty")
	}
	return key + "-" + strconv.FormatUint(GetGoroutineID(), 10)
}
