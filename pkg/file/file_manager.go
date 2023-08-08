package file

import (
	"os"
	"path/filepath"
)

type RelPath string
type AbsPath string

type Info struct {
	Name        string
	RelPath     RelPath
	AbsPath     AbsPath
	IsDirectory bool
}

type Manager struct {
	pwd string
}

func NewFileManager(pwd string) *Manager {
	return &Manager{pwd: pwd}
}

func (m *Manager) GetAbsPathsByInfos(infos []Info) []AbsPath {
	var result []AbsPath
	for _, info := range infos {
		result = append(result, info.AbsPath)
	}
	return result
}

func (m *Manager) GetDirRelPaths(relPath RelPath) ([]RelPath, error) {
	infos, err := m.GetFileInfos(relPath)
	if err != nil {
		return nil, err
	}

	var result []RelPath
	for _, info := range infos {
		if info.IsDirectory {
			result = append(result, info.RelPath)
		}
	}
	return result, nil
}

func (m *Manager) GetFileInfos(dirRelPath RelPath) ([]Info, error) {
	files, err := os.ReadDir(string(dirRelPath))
	if err != nil {
		return nil, err
	}

	var result []Info
	for _, elem := range files {
		rawInfo, err := elem.Info()
		if err != nil {
			return nil, err
		}

		fileRelPath := filepath.Join(string(dirRelPath), rawInfo.Name())
		fileAbsPath := filepath.Join(m.pwd, string(dirRelPath), rawInfo.Name())
		info := Info{
			Name:        rawInfo.Name(),
			RelPath:     RelPath(fileRelPath),
			AbsPath:     AbsPath(fileAbsPath),
			IsDirectory: rawInfo.IsDir(),
		}
		result = append(result, info)
	}
	return result, nil
}

func (m *Manager) ToStringsAbsPaths(absPaths []AbsPath) []string {
	var result []string
	for _, path := range absPaths {
		result = append(result, string(path))
	}
	return result
}
