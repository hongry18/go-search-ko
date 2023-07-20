package search

import (
	"bytes"
)

var (
	text    = []rune("가까나다따라마바빠사싸아자짜차카타파하")
	jamo    = []rune("ㄱㄲㄴㄷㄸㄹㅁㅂㅃㅅㅆㅇㅈㅉㅊㅋㅌㅍㅎ")
	overlay = []rune("ㄳㄵㄶㄺㄻㄼㄽㄾㄿㅀㅄ")
)

func SearchKo(s string) string {
	var (
		target rune
		search bytes.Buffer
		origin []rune
	)

	origin = []rune(s)
	target = origin[len(origin)-1]

	search.WriteByte('^')
	search.WriteString(string(origin[:len(origin)-1]))

	if !((target >= 12593 && target <= 12622) || (target >= 44032 && target <= 55203)) {
		return s
	}

	if target >= 44032 && target <= 55203 {
		var diff rune
		// 받침이 있는경우 리턴 - 강, 김
		switch (target - 44032) % 28 {
		case 0:
			// 받침이 없는 경우 규-귷
			diff = 27
		case 1, 4:
			// 받침이 ㄱ, ㄴ 인경우
			diff = 2
		case 17, 19:
			// 받침이 ㅂ, ㅅ 인경우
			diff = 1
		case 8:
			// 겹받침이 ㄹ인경우 ㅀ
			diff = 7
		default:
			return s
		}

		// 받침이 없는 경우 [가-갛], [규-귷]
		search.WriteString("[")
		search.WriteRune(target)
		search.WriteString("-")
		search.WriteRune(target + diff)
		search.WriteString("]")
		return search.String()
	}

	// 자음인 경우

	// 자음이 겹받침인 경우 - ㄳ · ㄵ · ㄶ · ㄺ · ㄻ · ㄼ · ㄽ · ㄾ · ㄿ · ㅀ · ㅄ
	for _, r := range overlay {
		if target == r {
			search.WriteRune(target)
			return search.String()
		}
	}

	// 자음인 경우 - (ㄱ|[가-깋]) 반환
	for i, r := range jamo {
		if target == r {
			search.WriteString("(")
			search.WriteRune(target)
			search.WriteString("|[")
			search.WriteRune(text[i])
			search.WriteString("-")
			search.WriteRune(text[i] + 587)
			search.WriteString("])")
			break
		}
	}

	return search.String()
}
