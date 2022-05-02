package lib

// BoolValue returns the value of the bool pointer or false if the pointer is nil
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// Int64Value returns the value of the int64 pointer or 0 if the pointer is nil
func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}
