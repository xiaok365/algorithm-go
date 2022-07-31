package collections

type List struct {
	data  []interface{}
	fps   []func(interface{}) bool
	mapFn MapFunc
}

type MapFunc func(interface{}) interface{}

func Identity() MapFunc {
	return func(i interface{}) interface{} { return i }
}

func Int2Inf(data []int) []interface{} {
	ret := make([]interface{}, 0)
	for _, v := range data {
		ret = append(ret, v)
	}
	return ret
}

func Str2Inf(data []string) []interface{} {
	ret := make([]interface{}, 0)
	for _, v := range data {
		ret = append(ret, v)
	}
	return ret
}

func NewList(resp []interface{}) *List {
	return &List{
		data: resp,
	}
}

func (a *List) Filter(fp func(interface{}) bool) *List {
	a.fps = append(a.fps, fp)
	return a
}

func (a *List) Map(mapToList MapFunc) *List {
	a.mapFn = mapToList
	return a
}

func (a *List) CollectToList() []interface{} {
	ret := make([]interface{}, 0)
	for _, v := range a.data {
		accept := true
		for _, fp := range a.fps {
			if !fp(v) {
				accept = false
			}
		}
		if accept {
			if a.mapFn != nil {
				ret = append(ret, a.mapFn(v))
			} else {
				ret = append(ret, v)
			}
		}
	}
	return ret
}

func (a *List) CollectToMap(mapA func(interface{}) interface{}, mapB func(interface{}) interface{}) map[interface{}]interface{} {
	ret := make(map[interface{}]interface{}, 0)
	for _, v := range a.data {
		accept := true
		for _, fp := range a.fps {
			if !fp(v) {
				accept = false
			}
		}
		if accept {
			key := mapA(v)
			value := mapB(v)
			ret[key] = value
		}
	}
	return ret
}
