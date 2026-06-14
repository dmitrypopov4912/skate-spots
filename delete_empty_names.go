package main

func FilterEmptyStrings(slice []string) []string {
	i := 0
	for j, _ := range slice {
		if slice[j] != "" {
			slice[i] = slice[j]
			i++
		}
	}
	return slice[:i]
}
