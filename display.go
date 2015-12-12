package sk

import (
	"bytes"
	"fmt"
	"io"
)

func Display(b Board) (string, error) {
	out := &bytes.Buffer{}
	_, err := out.WriteString("\n")
	if err != nil {
		return "", err
	}
	for row := 0; row < 8; row++ {
		rowData := b[row*9 : row*9+9]
		err = writeRow(out, rowData)
		if err != nil {
			return "", err
		}
		_, err = out.WriteString("\n")
		if err != nil {
			return "", err
		}
		if (row+1)%3 == 0 {
			_, err = out.WriteString("- - - + - - - + - - -\n")
			if err != nil {
				return "", err
			}
		}
	}
	err = writeRow(out, b[8*9:])
	if err != nil {
		return "", err
	}
	_, err = out.WriteString("\n")
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func writeElement(out io.Writer, element int) error {
	var err error
	if element == 0 {
		_, err = fmt.Fprintf(out, " ")
	} else {
		_, err = fmt.Fprintf(out, "%d", element)
	}
	return err
}

func writeRow(out io.Writer, row []int) error {
	for e := 0; e < 8; e++ {
		err := writeElement(out, row[e])
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(out, " ")
		if err != nil {
			return err
		}
		if (e+1)%3 == 0 {
			_, err := fmt.Fprintf(out, "| ")
			if err != nil {
				return err
			}
		}
	}
	return writeElement(out, row[8])
}
