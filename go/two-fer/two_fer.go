package twofer

// ShareWith based on input resolve results
func ShareWith(name string) string {
	var result string
	if name == "" {
		result = "One for you, one for me."
	} else {
		result = "One for " + name + ", one for me."
	}
	return result
}
