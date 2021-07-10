package util

func CheckStringInSlice(slice []string, searchedString string) bool {
	for _, s := range slice {
		if s == searchedString {
			return true
		}
	}
	return false
}
