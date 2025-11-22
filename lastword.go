package black

func LastWord(str string) string {

	i := len(str) - 1
	word := ""

	for i >= 0 {
		if str[i] != ',' && str[i] != ' ' {
			break
		}
		i--
	}
	end := i

	for i >= 0 {
		if str[i] == ',' || str[i] == ' ' {
			break
		}
		i--
	}
	start := i + 1

	for j := start; j <= end; j++ {
		word += string(str[j])
	}
	return word
}
