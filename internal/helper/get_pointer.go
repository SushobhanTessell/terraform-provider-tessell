package helper

func GetStringPointer(str interface{}) *string {
	if str != nil {
		str2 := str.(string)
		return &str2
	}
	var str2 string
	return &str2
}

func GetIntPointer(i interface{}) *int {
	if i != nil {
		i2 := i.(int)
		return &i2
	}
	var i2 int
	return &i2
}

func GetBoolPointer(b interface{}) *bool {
	if b != nil {
		b2 := b.(bool)
		return &b2
	}
	var b2 bool
	return &b2
}

func GetFloat32Pointer(f interface{}) *float32 {
	if f != nil {
		f2 := f.(float32)
		return &f2
	}
	var f2 float32
	return &f2
}

func GetFloat64Pointer(f interface{}) *float64 {
	if f != nil {
		f2 := f.(float64)
		return &f2
	}
	var f2 float64
	return &f2
}

func GetMapPointer(m map[string]interface{}) *map[string]interface{} {
	return &m
}
