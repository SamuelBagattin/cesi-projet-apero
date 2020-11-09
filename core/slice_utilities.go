package core

func ArrayContainsString(a *[]string, el string) bool {
	for _, b := range *a {
		if el == b {
			return true
		}
	}
	return false
}
