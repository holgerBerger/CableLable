# CableLable
create PDF from ASCII File for Cable labels

file format is one label per line, a line containing up to 3 words,
each word is a line on the label.
Words separated by spaces (line can therefor not contain spaces).

labels of type HellermanTyton TAG132LA4

usage: CableLable <filename>

generates a PDF <filename.pdf>

Filename is printed on top of sheet.

# Build instructions

use build.sh or build_win.sh
to do a native build or a (cross) build for windows.
