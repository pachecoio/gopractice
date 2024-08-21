package exercices

func MoreThanFiveLetters(originalSlice []string) []string {
	filteredSlice := []string{}
	for _, value := range originalSlice {
		if len(value) > 5 {
			filteredSlice = append(filteredSlice, value)
		}
	}
	return filteredSlice

}
