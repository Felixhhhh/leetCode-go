package main

func reverseString(s []byte) {

	left, right := 0, len(s)
	for left < right {
		temp := s[right]
		s[right] = s[left]
		s[left] = temp
	}

}

func main() {
	s := "abacada"

	println("s[]: ", s[3:5])
	println("s[]: ", s[2:5])
}
