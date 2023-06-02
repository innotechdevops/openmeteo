package openmeteo

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Float32(f float32) *float32 {
	return &f
}

func Bool(b bool) *bool {
	return &b
}
