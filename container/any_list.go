package container

type AnyList struct {
	anys []interface{}
}

func NewAnyList() *AnyList {
	return &AnyList{
		anys: make([]interface{}, 0),
	}
}

func (x *AnyList) Add(input ...interface{}) *AnyList {
	x.anys = append(x.anys, input...)
	return x
}

func (x *AnyList) Clear() *AnyList {
	x.anys = make([]interface{}, 0)
	return x
}

func (x *AnyList) Has(query interface{}) bool {
	for _, s := range x.anys {
		if s == query {
			return true
		}
	}
	return false
}

func (x *AnyList) All() []interface{} {
	return x.anys
}
