package app

import (
	"fmt"

	"github.com/hongry18/go-search-ko/internal/search"
)

func Run() {
	// for i := 12593; i <= 12622; i++ {
	// 	fmt.Print(string(rune(i)), ", ", rune(i), " - ")
	// 	fmt.Println(search.SearchKo(fmt.Sprintf("테스%c", rune(i))))
	// }

	// for i := 44032; i <= 55203; i++ {
	// 	fmt.Print(string(rune(i)), ", ", rune(i), " - ")
	// 	fmt.Println(search.SearchKo(fmt.Sprintf("테스%c", rune(i))))
	// }

	// for i := 62; i <= 148; i++ {
	// 	fmt.Print(string(rune(i)), ", ", rune(i), " - ")
	// 	fmt.Println(search.SearchKo(fmt.Sprintf("테스%c", rune(i))))
	// }

	// fmt.Println(search.SearchKo("유"))
	for i := 44032; i < 44332; i++ {
		fmt.Print(string(rune(i)), ", ", rune(i), " - ", (i-44032)%28, " --- ", search.SearchKo(string(rune(i))), "\n")
	}
	// fmt.Println(search.SearchKo(string(rune('ㄱ'))))

}
