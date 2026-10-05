func isAnagram(s string, t string) bool {
	mp := make(map[rune]int)
	flag := true
	s2 := []rune(s)
	t2 := []rune(t)
	len1 := len(s)
	len2 := len(t)
	if len1 != len2 {
		return false
	}
	
	for i:= 0; i < len1; i++ {
		elems := s2[i]
		mp[elems]++
	}

	for i:= 0; i < len2; i++ {
		elems := t2[i]
		if mp[elems] == 0 {
			return false
		} else {
			mp[elems]--
		}
		
	}

	return flag
}
