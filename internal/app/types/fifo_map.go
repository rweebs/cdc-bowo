package types

type KeyValue struct {
	Key   string
	Value interface{}
}

type FIFOMap struct {
	m    map[string]interface{}
	keys []string
}

func NewFIFOMap() *FIFOMap {
	return &FIFOMap{
		m:    make(map[string]interface{}),
		keys: make([]string, 0),
	}
}

func (f *FIFOMap) Set(key string, value interface{}) {
	_, ok := f.m[key]
	if !ok {
		f.keys = append(f.keys, key)
	}
	f.m[key] = value
}

func (f *FIFOMap) Get(key string) (interface{}, bool) {
	val, ok := f.m[key]
	return val, ok
}

func (f *FIFOMap) Delete(key string) {
	delete(f.m, key)
	for i, k := range f.keys {
		if k == key {
			f.keys = append(f.keys[:i], f.keys[i+1:]...)
			break
		}
	}
}

func (f *FIFOMap) Keys() []string {
	return f.keys
}

func (f *FIFOMap) Values() []KeyValue {
	values := make([]KeyValue, 0, len(f.m))
	for _, k := range f.keys {
		values = append(values, KeyValue{Key: k, Value: f.m[k]})
	}
	return values
}
