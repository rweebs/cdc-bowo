package types

type KeyValue struct {
	Key   string
	Value interface{}
}

type FIFOMap struct {
	m    map[string]interface{}
	keys []string
}

// NewFIFOMap creates a new FifoMap. The caller must call Free when done with the map
func NewFIFOMap() *FIFOMap {
	return &FIFOMap{
		m:    make(map[string]interface{}),
		keys: make([]string, 0),
	}
}

// Set a value in the map. If the key doesn t exist it will be added as a key to the FIFOMap
//
// Args:
//
//	f
//	key
//	value
func (f *FIFOMap) Set(key string, value interface{}) {
	_, ok := f.m[key]
	// Add a new key to the list of keys.
	if !ok {
		f.keys = append(f.keys, key)
	}
	f.m[key] = value
}

// Get returns the value associated with the key. If there is no value associated with the key false is returned
//
// Args:
//
//	f: the map to look up
//	key: the key to look
func (f *FIFOMap) Get(key string) (interface{}, bool) {
	val, ok := f.m[key]
	return val, ok
}

// Delete removes the key from the map. If the key does not exist nothing happens. It is safe to call this multiple times with the same key
//
// Args:
//
//	f: the map to remove the key from
//	key: the key to remove from the map. This can be a string
func (f *FIFOMap) Delete(key string) {
	delete(f.m, key)
	// Remove the key from the list of keys.
	for i, k := range f.keys {
		// Remove the key from the list of keys.
		if k == key {
			f.keys = append(f.keys[:i], f.keys[i+1:]...)
			break
		}
	}
}

// Keys returns the keys that have been added to the FIFOMap. This is a copy of the keys returned by Add so you don t need to worry about it being shared between goroutines.
//
// Args:
//
//	f: the map to get the keys from. Must not be nil
func (f *FIFOMap) Keys() []string {
	return f.keys
}

// Values returns a slice of KeyValue objects that contain all the values in the map. This is useful for debugging and to avoid having to iterate over the map every time it is called.
//
// Args:
//
//	f: the map to iterate over values from. Note that it is safe to modify the map
func (f *FIFOMap) Values() []KeyValue {
	values := make([]KeyValue, 0, len(f.m))
	// Add all values to the values.
	for _, k := range f.keys {
		values = append(values, KeyValue{Key: k, Value: f.m[k]})
	}
	return values
}
