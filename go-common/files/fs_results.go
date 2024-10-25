package files

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// FSResults is a convenient type to attach functionality at.
type FSResults []*FSResult

// Sort walks over each of the contained files,
// and sorts them on the base filename.
func (x FSResults) Sort() FSResults {
	var result = make([]*FSResult, len(x))
	copy(result, x)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Entry.Name() < result[j].Entry.Name()
	})
	return result
}

func (x FSResults) Filter(filter func(r *FSResult) bool) FSResults {
	var result FSResults
	for _, r := range x {
		if filter(r) {
			result = append(result, r)
		}
	}
	return result
}

// WriteDir writes each of the file names to provided dir.
// Paths will remain as is.
// If dir doesn't exist, will be created.
func (x FSResults) WriteDir(dir string) error {
	const mode = os.ModePerm
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	for _, r := range x {
		if r.Entry.IsDir() {
			continue
		}
		fullPath := filepath.Join(dir, r.FullPath())
		currDir := filepath.Dir(fullPath)
		if err := os.MkdirAll(currDir, mode); err != nil {
			return err
		}
		data, err := r.File.ReadFile(r.FullPath())
		if err != nil {
			return err
		}
		if err = os.WriteFile(fullPath, data, mode); err != nil {
			return err
		}
	}
	return nil
}

// WriteUniqueFiles works like WriteDir, except that it only considers the filename
// without path to be written to dir parameter.
// If filename is duplicated, an error will return.
func (x FSResults) WriteUniqueFiles(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	for _, r := range x {
		if r.Entry.IsDir() {
			continue
		}
		filename := r.Entry.Name()
		target := filepath.Join(dir, filename)
		// check file doesn't exist
		if _, err := os.Stat(target); err == nil {
			return fmt.Errorf("file already exists: %s", filename)
		}
		data, err := r.File.ReadFile(r.FullPath())
		if err != nil {
			return err
		}
		//nolint:gosec
		if err = os.WriteFile(target, data, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
