package asciiART

import (
	"fmt"
	"io/ioutil"
)

type Banner struct {
	Id          int
	AsciiSymbol [8][]rune
}

func InputFile(fileName, bannerFile string) {
	var bannerTemplateList []Banner = ReadTemplate(bannerFile)
	file, _ := ioutil.ReadFile(fileName)
	text := MakeRunes(file)
	var charCounterInTheRow int = 1
	for i := 0; i < len(text); i++ {
		if text[i] != 10 {
			charCounterInTheRow++
		} else {
			break
		}
	}
	var symbolFound bool = false
	var resultString string
	var saveIndex int = 0
	var startIndex int = 0

	for i := 0; i < len(bannerTemplateList); i++ {
		symbolFound = false
		for k := 0; k < len(bannerTemplateList[i].AsciiSymbol[0]); k++ {
			if text[k+startIndex] == bannerTemplateList[i].AsciiSymbol[0][k] && text[k+startIndex+charCounterInTheRow] == bannerTemplateList[i].AsciiSymbol[1][k] &&
				text[k+startIndex+charCounterInTheRow*2] == bannerTemplateList[i].AsciiSymbol[2][k] && text[k+startIndex+charCounterInTheRow*3] == bannerTemplateList[i].AsciiSymbol[3][k] &&
				text[k+startIndex+charCounterInTheRow*4] == bannerTemplateList[i].AsciiSymbol[4][k] && text[k+startIndex+charCounterInTheRow*5] == bannerTemplateList[i].AsciiSymbol[5][k] &&
				text[k+startIndex+charCounterInTheRow*6] == bannerTemplateList[i].AsciiSymbol[6][k] && text[k+startIndex+charCounterInTheRow*7] == bannerTemplateList[i].AsciiSymbol[7][k] {
				symbolFound = true
				saveIndex = k + 1
			} else {
				saveIndex = startIndex
				symbolFound = false
				break
			}
		}

		if symbolFound {
			startIndex = startIndex + saveIndex
			resultString = resultString + string(bannerTemplateList[i].Id)
			i = -1
			continue
		}
	}

	fmt.Println(resultString)
}

func MakeRunes(bytes []byte) []rune {
	var text []rune
	for i := range bytes {
		text = append(text, rune(bytes[i]))
	}
	return text
}

func ReadTemplate(bannerFile string) []Banner {
	var result []Banner
	var file []byte
	var text []rune

	file, _ = ioutil.ReadFile(MapFont(bannerFile))
	text = MakeRunes(file)
	var bannerToApply Banner
	var textIndex int

	if text[0] == 10 {
		textIndex = 1
	} else {
		textIndex = 2
	}

	for i := 32; i < 127; i++ {
		var tempArr [8][]rune

		for k := 0; k < 8; k++ {
			for l := 0; l < 32; l++ {

				if text[textIndex] == 13 {
					if text[textIndex+1] == 10 {
						textIndex = textIndex + 2
						break
					}
				} else if text[textIndex] == 10 {
					textIndex++
					break
				}

				tempArr[k] = append(tempArr[k], text[textIndex])
				textIndex++
			}
		}

		bannerToApply.Id = i
		bannerToApply.AsciiSymbol = tempArr
		result = append(result, bannerToApply)

		if i != 126 && text[textIndex] == 13 {
			textIndex = textIndex + 2
		} else {
			textIndex++
		}
	}

	return result
}
