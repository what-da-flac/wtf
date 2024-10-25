package csvutils

import (
	"encoding/csv"
	"errors"
	"io"
)

func Read(reader io.Reader, cb func(index int, values []string) error) error {
	csvReader := csv.NewReader(reader)
	var index int
	for {
		index++
		values, err := csvReader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if err := cb(index-1, values); err != nil {
			return err
		}
	}
	return nil
}
