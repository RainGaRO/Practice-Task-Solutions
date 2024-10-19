package main

func IsUnique(s string) bool {
	m := make(map[rune]struct{})
	if len(s) == 0 {
		return false
	}
	for _, elem := range s {
		if _, ok := m[elem]; ok {
			return false
		}
		m[elem] = struct{}{}
	}
	return true
}

func main() {

	//fmt.Println(IsUnique("asdfgh"))
}
