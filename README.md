# CableLable
create PDF from ASCII File for Cable labels

file format is one label per line, a line containing up to 3 words,
each word is a line on the label.
Words separated by spaces (line can therefor not contain spaces).

labels of type HellermanTyton TAG132LA4

usage: CableLable <filename>

generates a PDF <filename.pdf>

Filename is printed on top of sheet.

## Formating

At any point in the file can be lines of the form
```
# font <Fontname>
# face <face>
# size <size>
```
which change the default of font, face and size.
Font can be one of "Times", "Ariel", "Helvetica", "Courier",
Face can be "b" for bold or "i" for italic, "u" for underlined, or 3rd filed empty for standard,
size can be a size in pt, like 10,11,12.

Each word can contain `\b` for bold, `\T` for Times, `\H` for Helvetica,
`\l` for large and `\L` for larger, as well as `\s` for small and `\S` for smaller. 
This affects only that line.

## Build instructions

use build.sh or build_win.sh
to do a native build or a (cross) build for windows.
