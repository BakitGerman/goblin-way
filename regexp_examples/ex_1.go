package main

import (
	"fmt"
	"regexp"
)

func errorExample() {
	// можно сразу передавать шаблон в match для простых регулярных выражений
	// чтобы быстро проверить на соответствие
	ok, err := regexp.Match(")(", []byte("Bakit"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println(ok)
	}
}

// в golang регулярные выражения компилируются один раз
func main() {

	// начнем

	// можно сразу передавать шаблон в match для простых регулярных выражений и срез байт
	// чтобы быстро проверить на соответствие
	ok, err := regexp.Match(". . .", []byte("Bakit"))
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println(ok)
	}
	//errorExample()

	// можно так же и со строкой
	if ok, err := regexp.MatchString(". . .", "l l l"); err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Println(ok)
	}

	// экранирует все метасимволы регулярных выражений внутри текста аргумента
	fmt.Println(regexp.QuoteMeta(`()[0-9]`))

	// для более сложныж регулярных выражений нужно скомпилировать через Compile или MustCompile чтобы
	// использовать различные методы для работы с регулярными выражениями
	reg, err := regexp.Compile(`^[0-9][a-z]%`)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(reg.Match([]byte("999sss")))

	// или так, с паникой
	regM := regexp.MustCompile(`^[0-9][a-z]%`)
	fmt.Println(regM.Match([]byte("sss999")))

}
