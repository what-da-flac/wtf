package files

import (
	"embed"
	"io/fs"
	"path/filepath"
)

// FSPath encapsulates functionality for a set of embedded files.
//
// example:
//
// //go:embed test-data/*
//
//	var filesFs embed.FS
//
//	var pathFS = FSPath{
//		filesFs,
//	}
type FSPath struct {
	embed.FS
}

// NewFSPath is shortcut to manual instantiation.
func NewFSPath(fs embed.FS) FSPath {
	return FSPath{
		FS: fs,
	}
}

func Join(paths ...FSPath) (FSResults, error) {
	var result FSResults
	for _, p := range paths {
		entries, err := p.Entries()
		if err != nil {
			return nil, err
		}
		result = append(result, entries...)
	}
	return result, nil
}

// Entries walks over all the contained files and directories,
// and returns a list of results.
func (x FSPath) Entries() (FSResults, error) {
	// we don't have a starting point, so we always assume current directory is the root
	const root = "."
	var result []*FSResult
	if err := x.Walk(root, func(root string, entry fs.DirEntry, file *embed.FS) {
		result = append(result, &FSResult{
			Path:  root,
			Entry: entry,
			File:  file,
		})
	}); err != nil {
		return nil, err
	}
	return result, nil
}

// Walk parses a embedded file system recursively.
func (x FSPath) Walk(root string, walker func(root string, entry fs.DirEntry, file *embed.FS)) error {
	entries, err := x.ReadDir(root)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		walker(root, entry, &x.FS)
		if !entry.IsDir() {
			continue
		}
		if err = x.Walk(filepath.Join(root, entry.Name()), walker); err != nil {
			return err
		}
	}
	return nil
}
