package asciiART

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var termWidth int

func init() {
	termWidth = GetTermWidth()
}

func Aliging(letters_to_be_colored, text, font, align, color string) string {
	if align == "left" || align == "right" || align == "center" {
		return RightLeft(letters_to_be_colored, text, MapFont(font), align, color)
	} else if align == "justify" {
		Justify(letters_to_be_colored, text, MapFont(font), color)
	}
	return ""
}
func RightLeft(letters_to_be_colored, text, font, align, color string) string {
	res := ""
	newres := ""
	test := ""
	size := 0
	if text == "\\n" {
		return newres
	}
	args := strings.Split(text, "\\n")
	for _, word := range args {
		string_beg_end, _ := ContainsString(letters_to_be_colored, word)
		if word == "" {
		} else {
			for i := 0; i <= 8; i++ {
				str_len := len(word)
				for idx := 0; idx < str_len; idx++ {
					char := word[idx]
					//* Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
					if char == '\\' {
						if idx < len(word)-1 {
							//* Apply tab if 't' appears
							if word[idx+1] == 't' {
								res += "\t"
								idx++
								continue
							}
						}
					}
					test += PrintFileLine(MapART(rune(char))+i, font, "")
					if i == 0 {
						size = len(test)
					}
					if IsInRange(string_beg_end, idx) || letters_to_be_colored == "" {
						// Start printing the colored letler ART
						res += PrintFileLine(MapART(rune(char))+i, (font), color)
					} else {
						// Start printing the letler ART in default color
						res += PrintFileLine(MapART(rune(char))+i, (font), "")
					}
				}
				if align == "left" {
					newres += res
					newres += "\n"
					test = ""
				} else if align == "right" {
					newres += printSpaces(termWidth-size) + res
					newres += "\n"
					test = ""
				} else if align == "center" {
					newres += printSpaces((termWidth-size)/2) + res
					newres += "\n"
					test = ""
				}
				res = ""
			}
		}
	}
	return newres
}
func Justify(letters_to_be_colored, text, font, color string) {
	lines := strings.Split(text, "\\n")

	
	for _, v := range lines {
		var sws []string
		size := 0
		test := ""
		j := 0
		container := ""
		var string_beg_end [][]int
		sws = SplitWhiteSpacesAWESOME(v)
		string_beg_end, _ = ContainsString(letters_to_be_colored, v)
		ar := make([][]string, len(sws))
		if len(sws) == 1 {
			Ascii_Print(RightLeft(letters_to_be_colored, v, font, "left", color))
		} else {
			for i := 0; i < 8; i++ {
				j = 0
				for k, letter := range v {
					if letter != ' ' {
						test += PrintFileLine(MapART((letter))+i, (font), "")
						if IsInRange(string_beg_end, k) || letters_to_be_colored == "" {
							// Start printing the colored letler ART
							container += PrintFileLine(MapART((letter))+i, (font), color)
						} else {
							// Start printing the letler ART in default color
							container += PrintFileLine(MapART((letter))+i, (font), "")
						}
						//container += PrintFileLine((MapART(letter) + i), font, "")
					} else if letter == ' ' && container != "" {
						ar[j] = append(ar[j], container)
						container = ""
						j++
					}

				}
				ar[j] = append(ar[j], container)
				container = ""
				if i == 0 {
					size = len(test)
				}
			}
			textLen := 0
			for _, arOfStr := range ar {
				textLen += len(arOfStr[0])
			}

			numSpaces := (termWidth - size) / (len(sws) - 1)
			for i := 0; i < 8; i++ {
				for k, v := range ar {
					fmt.Print(v[i])
					if k != len(ar)-1 {
						fmt.Print(printSpaces(numSpaces))
					}
				}
				fmt.Println()
			}
		}
	}
}

func printSpaces(num int) string {
	a := ""
	for i := 1; i <= num; i++ {
		a += " "
	}
	return a
}

func GetTermWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	out = out[:len(out)-1]
	tput, _ := strconv.ParseInt(string(out[3:]), 10, 32)
	tput2 := int(tput)
	return tput2
}

func SplitWhiteSpacesAWESOME(str string) []string {

	nbw := 0
	var prel rune
	prel = ' '
	word := ""

	for _, v := range str {
		if (v != ' ' && v != '\t' && v != '\n') && (prel == ' ' || prel == '\t' || prel == '\n') {
			nbw++
		}
		prel = v
	}

	ar := make([]string, nbw)
	a := 0

	for _, v := range str {
		//fmt.Println(v)
		if v != ' ' && v != '\t' && v != '\n' && word != "" {
			word = word + string(v)
		}
		if v == ' ' || v == '\t' || v == '\n' && word != "" {
			a++
			ar[a-1] = word
			word = ""
		}

	}

	if word != "" {
		ar[a] = word
	}
	return ar
}
