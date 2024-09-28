package piscine

import (
	"strconv"
	"strings"
)

func FindInt(s string) int {
	arr := []rune(s)
	num := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			num = num*10 + int(arr[i]-'0')
		}
	}
	return num
}
func ModText(in string) string {
	input := Punct(in)
	input2 := Quot(input)
	text := strings.Fields(input2)
	for index, w := range text {
		if strings.Contains(w, "(hex)") {
			if index == 0 {
				continue
			}
			text[index-1] = ChangeHex(text[index-1])
			text[index] = Remove(text[index])
			text[index-1] = text[index-1] + text[index]
			text[index] = ""
		} else if w == "(hex," {
			if index == 0 || index >= len(text)-1 {
				continue
			}
			n := FindInt(text[index+1])
			ind := index - 1
			for j := 0; j < n; j++ {
				if ind < 0 {
					break
				}
				text[ind] = ChangeHex(text[ind])
				ind--
			}
			text[index] = ""
			text[index+1] = Remove(text[index+1])
			text[index-1] = text[index-1] + text[index+1]
			text[index+1] = ""
		}
		if strings.Contains(w, "(bin)") {
			if index == 0 {
				continue
			}
			text[index-1] = ChangeBin(text[index-1])
			text[index] = Remove(text[index])
			text[index-1] = text[index-1] + text[index]
			text[index] = ""
		} else if w == "(bin," {
			if index == 0 || index >= len(text)-1 {
				continue
			}
			n := FindInt(text[index+1])
			ind := index - 1
			for j := 0; j < n; j++ {
				if ind < 0 {
					break
				}
				text[ind] = ChangeBin(text[ind])
				ind--
			}
			text[index] = ""
			text[index+1] = Remove(text[index+1])
			text[index-1] = text[index-1] + text[index+1]
			text[index+1] = ""
		}
		if strings.Contains(w, "(up)") {
			if index == 0 {
				continue
			}
			text[index-1] = ToUp(text[index-1])
			text[index] = Remove(text[index])
			text[index-1] = text[index-1] + text[index]
			text[index] = ""
		} else if w == "(up," {
			if index == 0 || index >= len(text)-1 {
				continue
			}
			n := FindInt(text[index+1])
			ind := index - 1
			for j := 0; j < n; j++ {
				if ind < 0 {
					break
				}
				text[ind] = ToUp(text[ind])
				ind--
			}
			text[index] = ""
			text[index+1] = Remove(text[index+1])
			text[index-1] = text[index-1] + text[index+1]
			text[index+1] = ""
		}
		if strings.Contains(w, "(low)") {
			if index == 0 {
				continue
			}
			text[index-1] = ToLow(text[index-1])
			text[index] = Remove(text[index])
			text[index-1] = text[index-1] + text[index]
			text[index] = ""
		} else if w == "(low," {
			if index == 0 || index >= len(text)-1 {
				continue
			}
			n := FindInt(text[index+1])
			ind := index - 1
			for j := 0; j < n; j++ {
				if ind < 0 {
					break
				}
				text[ind] = ToLow(text[ind])
				ind--
			}
			text[index] = ""
			text[index+1] = Remove(text[index+1])
			text[index-1] = text[index-1] + text[index+1]
			text[index+1] = ""
		}
		if strings.Contains(w, "(cap)") {
			if index == 0 {
				continue
			}
			text[index-1] = Cap(text[index-1])
			text[index] = Remove(text[index])
			text[index-1] = text[index-1] + text[index]
			text[index] = ""
		} else if w == "(cap," {
			if index == 0 || index >= len(text)-1 {
				continue
			}
			n := FindInt(text[index+1])
			ind := index - 1
			for j := 0; j < n; j++ {
				if ind < 0 {
					break
				}
				text[ind] = Cap(text[ind])
				ind--
			}
			text[index] = ""
			text[index+1] = Remove(text[index+1])
			text[index-1] = text[index-1] + text[index+1]
			text[index+1] = ""
		}
	}
	out := strings.Join(text, " ")
	out = FixArticles(out)
	return out
}

func Remove(str string) string {
	results := ""
	inBrackets := false
	for i, char := range str {
		switch char {
		case ')':
			inBrackets = true
		default:
			if inBrackets && str[i] != ' ' {
				results += string(char)
			}
		}
	}
	return results
}

func Quot(s string) string {
	count := 0
	edit := strings.Fields(s)
	for r := 0; r < len(edit); r++ {
		if r < len(edit)-1 || r >= 0 {
			if edit[r] == string('\'') && count == 0 {
				edit[r] = ""
				edit[r+1] = string('\'') + edit[r+1]
				count = 1
			}
			if edit[r] == string('\'') && count == 1 {
				edit[r] = ""
				edit[r-1] = edit[r-1] + string('\'')
				count = 0
			}
			if edit[r] == string('(') {
				edit[r] = ""
				edit[r+1] = string('(') + edit[r+1]
			}
			if edit[r] == string(')') {
				edit[r] = ""
				edit[r-1] = edit[r-1] + string(')')
			}
		}
	}
	return strings.Join(edit, " ")
}
func IsPunc(r rune) bool {
	if r != '.' && r != ',' && r != '!' && r != '?' && r != ';' && r != ':' {
		return false
	}
	return true
}

func Punct(text6 string) string {
	modT := ""
	arr := []rune(text6)
	for i := 0; i < len(arr); i++ {
		if arr[i] != ' ' {
			modT += string(arr[i])
		}
		if i < len(arr)-1 {
			if !IsPunc(arr[i+1]) && arr[i] == ' ' {
				modT += " "

			} else if IsPunc(arr[i]) && !IsPunc(arr[i+1]) {
				modT += " "
			}
		}
	}
	return modT
}
func IsVowel(s string) bool {
	if strings.HasPrefix(s, "a") || strings.HasPrefix(s, "e") || strings.HasPrefix(s, "o") || strings.HasPrefix(s, "u") || strings.HasPrefix(s, "i") || strings.HasPrefix(s, "A") || strings.HasPrefix(s, "E") || strings.HasPrefix(s, "O") || strings.HasPrefix(s, "U") || strings.HasPrefix(s, "I") {
		return true
	} else {
		return false
	}
}

func FixArticles(text7 string) string {
	sub7 := strings.Fields(text7)
	for x := 0; x < len(sub7); x++ {
		if x < len(sub7)-1 {
			if sub7[x] == "a" && IsVowel(sub7[x+1]) {
				sub7[x] = "an"
			} else if sub7[x] == "A" && IsVowel(sub7[x+1]) {
				sub7[x] = "An"
			} else if sub7[x] == "an" && !IsVowel(sub7[x+1]) {
				sub7[x] = "a"
			} else if sub7[x] == "An" && !IsVowel(sub7[x+1]) {
				sub7[x] = "A"
			}
		}
	}
	return strings.Join(sub7, " ")
}

func ChangeHex(text1 string) string {
	result := ""
	dec, err := strconv.ParseInt(text1, 16, 0)
	if err == nil {
		result = strconv.Itoa(int(dec))
	}
	return result
}

func ChangeBin(text2 string) string {
	result := ""
	dec, err := strconv.ParseInt(text2, 2, 0)
	if err == nil {
		result = strconv.Itoa(int(dec))
	}
	return result
}

func ToUp(text3 string) string {
	result := ""
	result = strings.ToUpper(text3)
	return result
}

func ToLow(text4 string) string {
	result := strings.ToLower(text4)
	return result
}

func Cap(text5 string) string {
	c := 0
	arr := []rune(text5)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 65 && arr[i] <= 90 {
			c++
		}
		if arr[i] >= 97 && arr[i] <= 122 {
			arr[i] = arr[i] - 32
			c++
		}
		if c == 1 {
			break
		}
	}
	return string(arr)
}
