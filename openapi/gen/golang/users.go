package golang

type Users []*User

func (x Users) Len() int { return len(x) }

// ToMap returns a map which keys are user.id and values are pointers to users.
func (x Users) ToMap() map[string]*User {
	res := make(map[string]*User)
	for _, v := range x {
		key := v.Id
		if _, exists := res[key]; exists {
			continue
		}
		res[key] = v
	}
	return res
}

func (x Users) Unique() Users {
	var res Users
	seen := make(map[string]struct{})
	for _, v := range x {
		key := v.Id
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		res = append(res, v)
	}
	return res
}
