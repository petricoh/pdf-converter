package fileutil

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name        string
	AbsPath     string
	IsDirectory bool
	Size        int64
	Mode        fs.FileMode
	ModeTime    time.Time
}

func ReadDir(abs string) ([]*FileInfo, error) {
	entries, err := os.ReadDir(abs)
	if err != nil {
		return nil, err
	}
	result := make([]*FileInfo, 0)
	for _, entry := range entries {
		info, err := ParseDirEntry(abs, entry)
		if err != nil {
			return nil, err
		}
		result = append(result, info)
	}
	return result, nil
}

func ParseDirEntry(base string, entry os.DirEntry) (*FileInfo, error) {
	file, err := entry.Info()
	if err != nil {
		return nil, err
	}
	name := file.Name()
	return &FileInfo{
		Name:        name,
		AbsPath:     base + string(filepath.Separator) + name,
		IsDirectory: file.IsDir(),
		Size:        file.Size(),
		Mode:        file.Mode(),
		ModeTime:    file.ModTime(),
	}, nil
}
