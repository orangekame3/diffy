package diffy

import (
	"fmt"

	"github.com/kylelemons/godebug/diff"
)

// Ope defines the operation of a diff item.
type Ope int8

func (o Ope) String() string {
	switch o {
	case Delete:
		return "@Delete"
	case Add:
		return "@Add"
	case Equal:
		return "@Equal"
	}
	panic(fmt.Sprintf("unknown value of Ope: %d", int(o)))
}

const (
	// Delete item represents a delete diff.
	Delete Ope = -1
	// Add item represents an add diff.
	Add Ope = 1
	// Equal item represents an equal diff.
	Equal Ope = 0
)

type Line struct {
	Ope           Ope
	NewLineNumber int
	OldLineNumber int
	Text          string
}

func FormatWithLineNumber(chunks []diff.Chunk) []Line {
	var lines []Line
	oldLineNumber, newLineNumber := 1, 1

	for _, chunk := range chunks {
		for _, text := range chunk.Deleted {
			lines = append(lines, Line{
				Ope:           Delete,
				NewLineNumber: -1,
				OldLineNumber: oldLineNumber,
				Text:          text,
			})
			oldLineNumber++
		}

		for _, text := range chunk.Added {
			lines = append(lines, Line{
				Ope:           Add,
				NewLineNumber: newLineNumber,
				OldLineNumber: -1,
				Text:          text,
			})
			newLineNumber++
		}

		for _, text := range chunk.Equal {
			lines = append(lines, Line{
				Ope:           Equal,
				NewLineNumber: newLineNumber,
				OldLineNumber: oldLineNumber,
				Text:          text,
			})
			newLineNumber++
			oldLineNumber++
		}
	}

	return lines
}
