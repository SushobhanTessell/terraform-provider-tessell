package helper

func GetStringPointer(str string) *string {
	return &str
}

func GetIntPointer(i int) *int {
	return &i
}

func GetBoolPointer(b bool) *bool {
	return &b
}

func GetFloat32Pointer(f float32) *float32 {
	return &f
}

func GetFloat64Pointer(f float64) *float64 {
	return &f
}

func GetMapPointer(m map[string]interface{}) *map[string]interface{} {
	return &m
}
