package util

import (
	"encoding/csv"
	"io"
)

type CsvScanner struct {
	Reader *csv.Reader
	Head   map[string]int
	Row    []string
}

func NewCsvScanner(o io.Reader) CsvScanner {
	csv_o := csv.NewReader(o)
	a, e := csv_o.Read()
	if e != nil {
		return CsvScanner{}
	}
	m := map[string]int{}
	for n, s := range a {
		m[s] = n
	}
	return CsvScanner{Reader: csv_o, Head: m}
}

func (o *CsvScanner) Scan() bool {
	a, e := o.Reader.Read()
	o.Row = a
	return e == nil
}

func (o CsvScanner) Text(s string) string {
	i, ok := o.Head[s]

	if !ok {
		return ""
	}

	return o.Row[i]
}
