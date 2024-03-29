package main

import (
	"bufio"
	"fmt"
	"github.com/phpdave11/gofpdf"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	defaultfont = "Helvetica"
	defaultface = ""
	defaultsize = 10.0
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var (
		large = false
		filename = ""
	)

	if len(os.Args) < 2 {
		fmt.Println("usage: CableLable [--large] <filename>")
		fmt.Println(" file contains up to 3 words per line,")
		fmt.Println(" each word one line of label, each line a label.")
		os.Exit(0)
	}


	if os.Args[1] == "--large" {
		large=true	
		filename=os.Args[2]	
		fmt.Println("HellermanTyton TAG02LA4")
	} else {
		filename=os.Args[1]	
		fmt.Println("HellermanTyton TAG132LA4")
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	pdf.SetXY(1, 3)
	pdf.CellFormat(210, 12.3/3.0, filename , "0", 2, "C", false, 0, "")

	pdf.SetFont("Arial", "", 10)

	x := 0.0
	y := 0.0
	pages := 1
	newpage := false
	for {
		line, err := buf.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		lines := strings.Fields(string(line))

		if len(lines) == 0 {
			continue
		}

		if lines[0] == "#" {
			if lines[1] == "font" {
				defaultfont = lines[2]
				fmt.Println("Font set to", defaultfont)
			}
			if lines[1] == "face" {
				if len(lines) == 2 {
					defaultface = ""
					fmt.Println("Face set to standard")
				} else {
					defaultface = lines[2]
					fmt.Println("Face set to", defaultface)
				}
			}
			if lines[1] == "size" {
				s, err := strconv.Atoi(lines[2])
				if err == nil {
					defaultsize = float64(s)
					fmt.Println("Size set to", defaultsize)
				}
			}
			continue
		}

		if newpage {
			pdf.AddPage()
			fmt.Println("New page")
			pages += 1
			newpage = false
		}

		// NACH ANGABE HERSTELLER
		// pdf.SetXY(12.7+x*20.32, 8.47+y*38.10)

		// etwas getuned
		// pdf.SetXY(11.0+x*20.32, 9.0+y*38.30)
		if large {
			pdf.SetXY(12.0+x*25.4, 9.0+y*38.30)
		} else {
			pdf.SetXY(12.0+x*20.32, 9.0+y*38.30)
		}

		for i := 0; i < min(len(lines), 3); i++ {
			font := defaultfont
			face := defaultface
			size := defaultsize
			if strings.Contains(lines[i], `\b`) {
				face = "b"
				lines[i] = strings.Replace(lines[i], `\b`, "", 1)
			}
			if strings.Contains(lines[i], `\l`) {
				size = defaultsize + 1
				lines[i] = strings.Replace(lines[i], `\l`, "", 1)
			}
			if strings.Contains(lines[i], `\L`) {
				size = defaultsize + 1
				lines[i] = strings.Replace(lines[i], `\L`, "", 1)
			}
			if strings.Contains(lines[i], `\s`) {
				size = defaultsize - 1
				lines[i] = strings.Replace(lines[i], `\s`, "", 1)
			}
			if strings.Contains(lines[i], `\S`) {
				size = defaultsize - 2
				lines[i] = strings.Replace(lines[i], `\S`, "", 1)
			}
			if strings.Contains(lines[i], `\T`) {
				font = "Times"
				lines[i] = strings.Replace(lines[i], `\T`, "", 1)
			}
			if strings.Contains(lines[i], `\H`) {
				font = "Helvetica"
				lines[i] = strings.Replace(lines[i], `\H`, "", 1)
			}
			pdf.SetFont(font, face, size)
			// pdf.CellFormat(20.32, 12.3/3.0, lines[i], "0", 2, "L", false, 0, "")   // fixed line distance
			pdf.CellFormat(20.32, size*0.38, lines[i], "0", 2, "L", false, 0, "") // font size relative distance
			if large {
				pdf.TransformBegin()
				// leichte korrektur des rotationspunkts nach links -2mm
				pdf.TransformRotate(180, -2+(25.4/2.0)+12.0+x*25.4, (36.5/2.0)+9.0+y*38.30)
				pdf.CellFormat(20.32, size*0.38, lines[i], "0", 2, "L", false, 0, "") // font size relative distance
				pdf.TransformEnd()
			}
		}

		x = x + 1
		if large {
			if x == 7 {
				x = 0
				y = y + 1
			}
		} else {
			if x == 9 {
				x = 0
				y = y + 1
			}
		}
		if y == 7 {
			y = 0
			newpage = true
		}
	}

	fmt.Println(pages, "pages")
	err = pdf.OutputFileAndClose(filename + ".pdf")
	if err != nil {
		panic(err)
	}
}
