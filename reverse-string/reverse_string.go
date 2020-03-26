package reverse

func Reverse(text string) string {
	var reverseRunes []rune  
	runes := []rune(text)
	for i := len(runes)-1; i >= 0 ; i-- {
		reverseRunes = append(reverseRunes, runes[i])
	}
	return string(reverseRunes)
}   