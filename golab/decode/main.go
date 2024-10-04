package main

import (
	"fmt"
	"strings"
)

func main() {

	dataFromPdf := "<20>5<150B090507>1<0804>-6<07>1<0905150A07>1<0E090507>1<0B>-8<0903>1<1102>2<07>1<0B>-8<091502>2<0507>1<0B0919081517>-2<08090E07>1<1107>1<0B09050D>2<0C>-2<160921>-1<150B0904>1<15120802>2<14>-7<090815>8<17>-2<090C>-2<0D>2<0C>-2<02>2<07>1<0B>-8<0907>1<081003>1<1007>1<090503>1<0E>-7<090C>6<07>1<03>1<0C>-2<1508091307>1<02>-5<02>2<07>1<0B>-8<09>"
	data := []byte{}
	removeABCD := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(strings.ReplaceAll(dataFromPdf, "A", "10"), "B", "11"), "C", "12"), "D", "13"), "E", "14"), "F", "15")
	for _, i := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(removeABCD, "<", ""), ">", ""), " ") {
		println(i)
	}
	fmt.Println(data)
}

// func DecmialToBinary(number string){

// }
