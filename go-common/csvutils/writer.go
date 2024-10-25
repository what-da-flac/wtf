package csvutils

import (
	"encoding/csv"
	"io"
)

func Write(w io.Writer, headers []string, rowFn func(index int) ([]string, error)) error {
	var (
		index    int
		errWrite error
	)
	writer := csv.NewWriter(w)
	if err := writer.Write(headers); err != nil {
		return err
	}
	for {
		values, err := rowFn(index)
		if err != nil {
			break
		}
		if err := writer.Write(values); err != nil {
			errWrite = err
			break
		}
		index++
	}
	if errWrite != nil {
		return errWrite
	}
	writer.Flush()
	return nil
}
