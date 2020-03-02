package stringutil

func NullIfBlank(str *string) *string {
	if str != nil && *str == "" {
		return nil
	}
	return str
}