package istring

type IString []string

func (x IString) Filter(filter func(s string) bool) IString {
	var res IString
	for _, v := range x {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func (x IString) Reverse() IString {
	var res IString
	for i := len(x) - 1; i >= 0; i-- {
		res = append(res, x[i])
	}
	return res
}

func (x IString) Len() int { return len(x) }

func (x IString) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x IString) Less(i, j int) bool { return x[i] < x[j] }
