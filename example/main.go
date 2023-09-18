package main

import (
	"fmt"
	"strings"

	"github.com/orangekame3/diffy"

	"github.com/kylelemons/godebug/diff"
)

func main() {
	a := strings.TrimSpace(`
	We the People of the United States, in Order to form a more perfect Union,
	establish Justice, insure domestic Tranquility, provide for the common defence,
	promote the general Welfare, and secure the Blessings of Liberty to ourselves
	and our Posterity, do ordain and establish this Constitution for the United
	States of America.
	`)

	b := strings.TrimSpace(`
	:wq
	We the People of the United States, in Order to form a more perfect Union,
	establish Justice, insure domestic Tranquility, provide for the common defence,
	and secure the Blessings of Liberty to ourselves
	and our Posterity, do ordain and establish this Constitution for the United
	States of America.
	`)

	aLines := strings.Split(a, "\n")
	bLines := strings.Split(b, "\n")
	chunks := diff.DiffChunks(aLines, bLines)
	lines := diffy.FormatWithLineNumber(chunks)
	for _, line := range lines {
		fmt.Println(line.Ope, line.NewLineNumber, line.OldLineNumber, line.Text)
	}

}
