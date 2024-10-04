package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"log"
	"strings"
)

func main() {

	dataFromPdf := "<20>5<150B090507>1<0804>-6<07>1<0905150A07>1<0E090507>1<0B>-8<0903>1<1102>2<07>1<0B>-8<091502>2<0507>1<0B0919081517>-2<08090E07>1<1107>1<0B09050D>2<0C>-2<160921>-1<150B0904>1<15120802>2<14>-7<090815>8<17>-2<090C>-2<0D>2<0C>-2<02>2<07>1<0B>-8<0907>1<081003>1<1007>1<090503>1<0E>-7<090C>6<07>1<03>1<0C>-2<1508091307>1<02>-5<02>2<07>1<0B>-8<09>"
	data := [][]byte{}
	for _, i := range strings.Split(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(dataFromPdf, "<", ""), ">", ""), "-", " "), " ") {
		data = append(data, HexaToBinary(i))
	}
	result := ReduceByteArray(data)
	r, err := zlib.NewReader(bytes.NewReader(result))
	if err != nil {
		log.Println("first log", err)
	}

	for {
		n, err := r.Read(result)
		fmt.Println(result[:n])
		if err != nil {
			log.Println("second log", err)
		}
	}

}

func HexaToBinary(number string) []byte {
	hexatobinary := map[string][]byte{}
	//                       8,4,2,1
	hexatobinary["0"] = []byte{0, 0, 0, 0}
	hexatobinary["1"] = []byte{0, 0, 0, 1}
	hexatobinary["2"] = []byte{0, 0, 1, 0}
	hexatobinary["3"] = []byte{0, 0, 1, 1}
	hexatobinary["4"] = []byte{0, 1, 0, 0}
	hexatobinary["5"] = []byte{0, 1, 0, 1}
	hexatobinary["6"] = []byte{0, 1, 1, 0}
	hexatobinary["7"] = []byte{0, 1, 1, 1}
	hexatobinary["8"] = []byte{1, 0, 0, 0}
	hexatobinary["9"] = []byte{1, 0, 0, 1}
	hexatobinary["A"] = []byte{1, 0, 1, 0}
	hexatobinary["B"] = []byte{1, 0, 1, 1}
	hexatobinary["C"] = []byte{1, 1, 0, 0}
	hexatobinary["D"] = []byte{1, 1, 0, 1}
	hexatobinary["E"] = []byte{1, 1, 1, 0}
	hexatobinary["F"] = []byte{1, 1, 1, 1}
	numberbinary := [][]byte{}
	for _, i := range number {
		numberbinary = append(numberbinary, hexatobinary[string(i)])
	}
	return ReduceByteArray(numberbinary)
}

func ReduceByteArray(array [][]byte) []byte {
	result := []byte{}
	for _, i := range array {
		for _, j := range i {
			result = append(result, j)
		}
	}
	return result
}
