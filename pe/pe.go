package pe

import (
	"bytes"
	"debug/pe"
)

func Fuzz(data []byte) int {
	f, err := pe.NewFile(bytes.NewReader(data))
	if err != nil {
		if f != nil {
			panic("file is not nil on error")
		}
		return 0
	}
	defer f.Close()
	f.ImportedLibraries()
	f.ImportedSymbols()
	f.Section(".data")
	f.Section(".text")
	for _, c := range f.COFFSymbols {
		_, err := c.FullName(f.StringTable)
		if err != nil {
			return 1
		}
	}
	dw, err := f.DWARF()
	if err != nil {
		if dw != nil {
			panic("dwarf is not nil on error")
		}
		return 1
	}
	dr := dw.Reader()
	for {
		e, _ := dr.Next()
		if e == nil {
			break
		}
	}
	return 2
}
