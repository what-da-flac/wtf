package converters

func InterfaceToString(i interface{}) string {
	if val, ok := i.(string); ok {
		return val
	}
	return ""
}

func InterfaceToInt(i interface{}) int {
	if val, ok := i.(int); ok {
		return val
	}
	return 0
}

func InterfaceToBool(i interface{}) bool {
	if val, ok := i.(bool); ok {
		return val
	}
	return false
}
