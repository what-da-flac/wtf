package i_am

type ModelGroups []*ModelGroup

func (x ModelGroups) UnPtr() []ModelGroup {
	var res []ModelGroup
	for _, item := range x {
		res = append(res, *item)
	}
	return res
}
