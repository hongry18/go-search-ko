package search

import (
	"fmt"
	"strings"
)

var (
	text    = []rune("가깋닣딯띻맇밓빟삫싷앃잏짛찧칳킿팋핗힣")
	jamo    = []rune("ㄱㄲㄴㄷㄸㄹㅁㅂㅃㅅㅆㅇㅈㅉㅊㅋㅌㅍㅎ")
	overlay = []rune("ㄳㄵㄶㄺㄻㄼㄽㄾㄿㅀㅄ")
)

const (
	diff rune = 588
)

func SearchKo(s string) string {
	var (
		target rune
		search string
		origin []rune
	)

	origin = []rune(s)
	target = origin[len(origin)-1]
	origin = origin[:len(origin)-1]

	if !((target >= 12593 && target <= 12622) || (target >= 44032 && target <= 55203)) {
		return s
	}

	if target <= 12622 {
		var isOverlay bool
		for _, r := range overlay {
			if target == r {
				isOverlay = true
				search = string(target)
			}
		}

		if !isOverlay {
			for i, r := range jamo {
				if target == r {
					search = fmt.Sprintf("(%c|[%c-%c])", target, text[i]-diff+1, text[i])
				}
			}
		}
	} else {
		for i, r := range text {
			if target > r && target <= text[i+1] {
				search = fmt.Sprintf("[%c-%c]", target, text[i+1])
			}
		}
	}

	return strings.Join([]string{string(origin), search}, "")
}
