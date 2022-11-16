package zeromdc

import (
	"fmt"
	"sync"
)

var (
	_globalMdc  = InitGlobalMdcAdapter() //nolint:gochecknoglobals
	_MdcAdapter *MdcAdapter              //nolint:gochecknoglobals
)

type MdcAdapter struct {
	items map[string]interface{}
	sync.RWMutex
}

func InitGlobalMdcAdapter() *MdcAdapter {
	if _MdcAdapter == nil {
		_MdcAdapter = &MdcAdapter{
			items:   make(map[string]interface{}),
			RWMutex: sync.RWMutex{},
		}
	}

	return _MdcAdapter
}

func ResetGlobalMdcAdapter() {
	_MdcAdapter.RLock()
	_MdcAdapter.items = make(map[string]interface{})
	_MdcAdapter.RUnlock()
}

func MDC() *MdcAdapter {
	_MdcAdapter.RLock()
	s := _globalMdc
	_MdcAdapter.RUnlock()

	return s
}

func (m *MdcAdapter) Set(key string, value interface{}) {
	m.Lock()
	m.items[key] = value
	m.Unlock()
}

func (m *MdcAdapter) Get(key string) (interface{}, bool) {
	m.RLock()
	v, ok := m.items[key]
	m.RUnlock()

	return v, ok
}

func (m *MdcAdapter) GetString(key string) string {
	m.RLock()
	value, ok := m.items[key]
	m.RUnlock()

	if !ok {
		value = ""
	}

	return fmt.Sprintf("%v", value)
}

func (m *MdcAdapter) Remove(key string) {
	m.Lock()
	delete(m.items, key)
	m.Unlock()
}
