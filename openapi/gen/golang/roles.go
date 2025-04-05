package golang

type Roles []*Role

func (x Roles) Names() []string {
	var res []string
	for _, role := range x {
		res = append(res, role.Name)
	}
	return res
}
