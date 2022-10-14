package helper

func ReverseString(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}
	reversedStr := ReverseString(str[1:])
	return reversedStr + string(str[0])
}
