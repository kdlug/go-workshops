package main

import (
	"fmt"
)

func main() {
	const wrapLimit int = 40

	text := `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.

Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
	
Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`
	var lr int // line runes
	//var count int
	for _, r := range text {
		fmt.Printf("%c", r)
		lr++
		if lr > wrapLimit && r == ' ' {
			lr = 0
			fmt.Println()
		} else if r == '\n' {
			lr = 0
		}

	}
	fmt.Println()
}
