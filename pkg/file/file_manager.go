package file

import (
	"os"
	"path/filepath"
	"strings"
)

type Info struct {
	Name        string
	Path        string
	IsDirectory bool
}

type Manager struct {
	pwd string
}

func NewFileManager(pwd string) *Manager {
	return &Manager{pwd: pwd}
}

func (r *Manager) GetPathsByInfos(infos []Info) []string {
	var result []string
	for _, info := range infos {
		result = append(result, info.Path)
	}
	return result
}

func (r *Manager) GetDirPaths(path string) ([]string, error) {
	infos, err := r.GetFileInfos(path)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, info := range infos {
		if info.IsDirectory {
			result = append(result, info.Path)
		}
	}
	return result, nil
}

func (r *Manager) GetFileInfos(path string) ([]Info, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var result []Info
	for _, elem := range files {
		rawInfo, err := elem.Info()
		if err != nil {
			return nil, err
		}

		middlePath := strings.ReplaceAll(path, r.pwd, "")
		path := filepath.Join(r.pwd, middlePath, rawInfo.Name())
		info := Info{
			Name:        rawInfo.Name(),
			Path:        path,
			IsDirectory: rawInfo.IsDir(),
		}
		result = append(result, info)
	}
	return result, nil
}
