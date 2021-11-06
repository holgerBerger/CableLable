package main

import (
	"bufio"
	"fmt"
	"github.com/phpdave11/gofpdf"
	"io"
	"os"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: CableLable <filename>")
		fmt.Println(" file contains up to 3 words per line,")
		fmt.Println(" each word one line of label, each line a label.")
		os.Exit(0)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	pdf.SetXY(1, 3)
	pdf.CellFormat(210, 12.3/3.0, os.Args[1], "0", 2, "C", false, 0, "")

	pdf.SetFont("Arial", "", 10)

	x := 0.0
	y := 0.0
	for {
		line, err := buf.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		lines := strings.Fields(string(line))

		// NACH ANGABE HERSTELLER
		// pdf.SetXY(12.7+x*20.32, 8.47+y*38.10)

		// etwas getuned
		pdf.SetXY(11.0+x*20.32, 9.0+y*38.30)
		for i := 0; i < min(len(lines), 3); i++ {
			pdf.CellFormat(20.32, 12.3/3.0, lines[i], "0", 2, "L", false, 0, "")
		}

		x = x + 1
		if x == 9 {
			x = 0
			y = y + 1
		}
		if y == 7 {
			y = 0
			pdf.AddPage()
		}
	}

	err = pdf.OutputFileAndClose(os.Args[1] + ".pdf")
	if err != nil {
		panic(err)
	}
}