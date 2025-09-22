package maps

type MapsLib struct{}

func (s *MapsLib) AddIfStrNotEmpty(m map[string]any, key string, val string) {
	if val != "" {
		m[key] = val
	}
}

func (s *MapsLib) AddIfStrPtrNotEmpty(m map[string]any, key string, ptr *string) {
	if ptr != nil && *ptr != "" {
		m[key] = ptr
	}
}

func (s *MapsLib) AddIfSliceNotEmpty(m map[string]any, key string, vals []string) {
	if len(vals) > 0 {
		m[key] = vals
	}
}
