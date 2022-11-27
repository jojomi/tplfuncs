package container

type StringAnyMap struct {
	data map[string]interface{}
}

func NewStringAnyMap() *StringAnyMap {
	return &StringAnyMap{
		data: make(map[string]interface{}, 0),
	}
}

func NewStringAnyMapFromMap(data map[string]interface{}) *StringAnyMap {
	return &StringAnyMap{
		data: data,
	}
}

func (x *StringAnyMap) Add(key string, value interface{}) *StringAnyMap {
	x.data[key] = value
	return x
}

func (x *StringAnyMap) AddAll(input ...interface{}) *StringAnyMap {
	for i := 0; i < len(input)-1; i = i + 2 {
		key, ok := input[i].(string)
		if !ok {
			continue
		}
		x.data[key] = input[i+1]
	}
	return x
}

func (x *StringAnyMap) SetDefaults(input ...interface{}) *StringAnyMap {
	for i := 0; i < len(input)-1; i = i + 2 {
		key, ok := input[i].(string)
		if !ok {
			continue
		}
		if _, ok := x.data[key]; ok {
			continue
		}
		x.data[key] = input[i+1]
	}
	return x
}

func (x *StringAnyMap) Clear() *StringAnyMap {
	x.data = make(map[string]interface{}, 0)
	return x
}

func (x *StringAnyMap) HasKey(query string) bool {
	for k := range x.data {
		if k == query {
			return true
		}
	}
	return false
}

func (x *StringAnyMap) AsMap() map[string]interface{} {
	return x.data
}

func (x *StringAnyMap) HasValue(query interface{}) bool {
	for _, v := range x.data {
		if v == query {
			return true
		}
	}
	return false
}

func (x *StringAnyMap) Keys() *StringList {
	r := NewStringList()
	for k := range x.data {
		r.Add(k)
	}
	return r
}

func (x *StringAnyMap) Values() *AnyList {
	r := NewAnyList()
	for _, v := range x.data {
		r.Add(v)
	}
	return r
}
