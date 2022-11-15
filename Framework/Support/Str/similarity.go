package Str

// SimilarText2 calculate the similarity between two strings
// returns the number of matching chars in both strings and the similarity percent
// note that there is a little difference between the original php function that
// a multi-byte character is counted as 1 while it would be 2 in php
//
// This function is case sensitive
func SimilarText2(first, second string) (count int, percent float32) {
	txt1 := []rune(first)
	len1 := len(txt1)
	txt2 := []rune(second)
	len2 := len(txt2)

	if len1 == 0 || len2 == 0 {
		return 0, 0
	}

	sim := similarChar(txt1, len1, txt2, len2)

	return sim, float32(sim*200) / float32(len1+len2)
}

func similarChar(txt1 []rune, len1 int, txt2 []rune, len2 int) int {
	var sum int

	pos1, pos2, max := similarStr(txt1, len1, txt2, len2)

	if sum = max; sum > 0 {
		if pos1 > 0 && pos2 > 0 {
			sum += similarChar(txt1, pos1, txt2, pos2)
		}
		if (pos1+max < len1) && (pos2+max < len2) {
			sum += similarChar(txt1[(pos1+max):], len1-pos1-max, txt2[(pos2+max):], len2-pos2-max)
		}
	}

	return sum
}

// SimilarText similar_text()
func SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}

func similarStr(txt1 []rune, len1 int, txt2 []rune, len2 int) (pos1, pos2, max int) {
	var l int
	pos1, pos2, max = 0, 0, 0
	for i := range txt1 {
		for j := range txt2 {
			for l = 0; i+l < len1 && j+l < len2 && txt1[i+l] == txt2[j+l]; l++ {
			}
			if l > max {
				max = l
				pos1 = i
				pos2 = j
			}
		}
	}
	return pos1, pos2, max
}
