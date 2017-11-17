package stringutil

func Reverse(s string)  string{
	stringArr := []rune(s)
	for i, j := 0, len(stringArr) -1; i < len(stringArr)/2 ; i, j= i+1 , j-1 {
		stringArr[i] , stringArr[j] = stringArr[j], stringArr[i]
	}
	return string(stringArr)
}