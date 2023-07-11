package main

import (
	"asciiART"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var ascii_art struct {
		text         string
		size         int
		text_nocolor string
	}

	// Flags
	colorname := flag.String("color", "", "coloring")
	outputfile := flag.String("output", "", "output")
	justify := flag.String("align", "", "aliging")
	reverse := flag.String("reverse", "", "reverse")
	flag.Parse()

	if *reverse != "" && *outputfile == "" && *justify == "" && *colorname == "" {
		asciiART.InputFile(*reverse, os.Args[2])
		return
	} else if *colorname != "" && *outputfile == "" && *justify == "" {
		asciiART.ColorFlag(os.Args, *colorname)
		return
	} else if *colorname != "" && *outputfile != "" && *justify == "" && (strings.Contains(os.Args[1], "--color=") && strings.Contains(os.Args[2], "--output=") || strings.Contains(os.Args[2], "--color=") && strings.Contains(os.Args[1], "--output=")) {
		if len(os.Args) == 4 {
			ascii_art.text = asciiART.PrintART("", os.Args[3], "standard", *colorname)
			asciiART.Ascii_Print(ascii_art.text)
			ascii_art.text_nocolor = asciiART.PrintART("", os.Args[3], "standard", "")
			asciiART.ExportFile(*outputfile, ascii_art.text_nocolor)
		} else if len(os.Args) == 5 {
			if os.Args[4] == "shadow" || os.Args[4] == "standard" || os.Args[4] == "thinkertoy" || os.Args[4] == "weird" || os.Args[4] == "refined" {
				ascii_art.text = asciiART.PrintART("", os.Args[3], os.Args[4], *colorname)
				asciiART.Ascii_Print(ascii_art.text)
				ascii_art.text_nocolor = asciiART.PrintART("", os.Args[3], os.Args[4], "")
				asciiART.ExportFile(*outputfile, ascii_art.text_nocolor)
			} else {
				ascii_art.text = asciiART.PrintART(os.Args[3], os.Args[4], "standard", *colorname)
				asciiART.Ascii_Print(ascii_art.text)
				ascii_art.text_nocolor = asciiART.PrintART("", os.Args[4], "standard", "")
				asciiART.ExportFile(*outputfile, ascii_art.text_nocolor)
			}
		} else if len(os.Args) == 6 {
			ascii_art.text = asciiART.PrintART(os.Args[3], os.Args[4], os.Args[5], *colorname)
			asciiART.Ascii_Print(ascii_art.text)
			ascii_art.text_nocolor = asciiART.PrintART(os.Args[3], os.Args[4], os.Args[5], "")
			asciiART.ExportFile(*outputfile, ascii_art.text_nocolor)
		}
		return
	} else if *colorname == "" && *justify == "" && *outputfile != "" && strings.Contains(os.Args[1], "--output=") {
		if len(os.Args) == 3 {
			ascii_art.text = asciiART.PrintART("", os.Args[2], "standard", "")
			asciiART.Ascii_Print(ascii_art.text)
			asciiART.ExportFile(*outputfile, ascii_art.text)
		} else if len(os.Args) == 4 {
			ascii_art.text = asciiART.PrintART("", os.Args[2], os.Args[3], "")
			asciiART.Ascii_Print(ascii_art.text)
			asciiART.ExportFile(*outputfile, ascii_art.text)
		}
		return
	} else if *colorname == "" && *justify != "" && *outputfile == "" && strings.Contains(os.Args[1], "--align=") {
		if *justify != "right" && *justify != "left" && *justify != "justify" && *justify != "center" {
			fmt.Println("Error: align must be right or left or center or justify")
		}
		if len(os.Args) == 3 {
			asciiART.Ascii_Print(asciiART.Aliging("", os.Args[2], "standard", *justify, ""))
		} else if len(os.Args) == 4 {
			asciiART.Ascii_Print(asciiART.Aliging("", os.Args[2], os.Args[3], *justify, ""))
		} else {
			asciiART.PrintError(1)
		}
		return
	} else if *colorname != "" && *justify != "" && *outputfile == "" && (strings.Contains(os.Args[1], "--color=") && strings.Contains(os.Args[2], "--align=") || strings.Contains(os.Args[2], "--color=") && strings.Contains(os.Args[1], "--align=")) {
		if *justify != "right" && *justify != "left" && *justify != "justify" && *justify != "center" {
			fmt.Println("Error: align must be right or left or center or justify")
		}
		if len(os.Args) == 4 {
			asciiART.Ascii_Print(asciiART.Aliging("", os.Args[3], "standard", *justify, *colorname))
		} else if len(os.Args) == 5 {
			if os.Args[4] == "shadow" || os.Args[4] == "standard" || os.Args[4] == "thinkertoy" || os.Args[4] == "weird" || os.Args[4] == "refined" {
				asciiART.Ascii_Print(asciiART.Aliging("", os.Args[3], os.Args[4], *justify, *colorname))
			} else {
				asciiART.Ascii_Print(asciiART.Aliging(os.Args[3], os.Args[4], "standard", *justify, *colorname))
			}
		} else if len(os.Args) == 6 {
			asciiART.Ascii_Print(asciiART.Aliging(os.Args[3], os.Args[4], os.Args[5], *justify, *colorname))
		}
		return
	} else if *justify != "" && *outputfile != "" && *colorname != "" {
		terminalargs := os.Args[4:]
		switch len(terminalargs) {
		case 0:
			asciiART.PrintError(0)
		case 1:
			asciiART.Ascii_Print(asciiART.Aliging("", terminalargs[0], "standard", *justify, *colorname))
			asciiART.ExportFile(*outputfile, asciiART.Aliging("", terminalargs[0], "standard", *justify, ""))
			return
		case 2:
			switch terminalargs[1] {
			case "shadow", "standard", "thinkertoy", "weird", "refined":
				asciiART.Ascii_Print(asciiART.Aliging("", terminalargs[0], terminalargs[1], *justify, *colorname))
				asciiART.ExportFile(*outputfile, asciiART.Aliging("", terminalargs[0], terminalargs[1], *justify, ""))
				return
			default:
				asciiART.Ascii_Print(asciiART.Aliging(terminalargs[0], terminalargs[1], "standard", *justify, *colorname))
				return
			}
		case 3:
			asciiART.Ascii_Print(asciiART.Aliging(terminalargs[0], terminalargs[1], terminalargs[2], *justify, *colorname))
			asciiART.ExportFile(*outputfile, asciiART.Aliging("", terminalargs[1], terminalargs[2], *justify, ""))
			return
		}
	} else {
		if len(os.Args) == 2 {
			ascii_art.text = asciiART.PrintART("", os.Args[1], "standard", "")
			asciiART.Ascii_Print(ascii_art.text)
			return
		} else if len(os.Args) == 3 {
			ascii_art.text = asciiART.PrintART("", os.Args[1], os.Args[2], "")
			asciiART.Ascii_Print(ascii_art.text)
			return
		}
	}
	asciiART.PrintError(0)
}
