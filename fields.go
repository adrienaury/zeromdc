package zeromdc

//nolint:gochecknoglobals
var (
	_uniqueness   = make(map[string]interface{}, 0)
	_globalFields = make([]string, 0)
)

func SetGlobalFields(fields []string) {
	_globalFields = fields

	for _, field := range fields {
		_uniqueness[field] = nil
	}
}

func GetGlobalFields() []string {
	return _globalFields
}

func AddGlobalFields(field string) {
	if _, exist := _uniqueness[field]; !exist {
		_globalFields = append(_globalFields, field)
	}
}

func ClearGlobalFields() {
	_globalFields = make([]string, 0)
	_uniqueness = make(map[string]interface{}, 0)
}
