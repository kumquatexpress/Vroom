package helpers

import (
	"github.com/kumquatexpress/Vroom/logger"
	"github.com/kumquatexpress/Vroom/utils"
	"os"
	"path/filepath"
	"text/template"
)

type DirTree struct {
	Template  *template.Template
	Filenames []string
	Data map[string]interface{}
}

func isCorrectExtension(info os.FileInfo) bool {
	matches, err := filepath.Match("*.vroom.html", info.Name())
	return !info.IsDir() && err == nil && matches
}

func FindLayoutFiles(vo *VroomOpts) []string {
	utils.FindOrCreateDir(vo.LayoutDirectory)
	var files []string
	filepath.Walk(vo.LayoutDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error(err.Error())
		}
		if isCorrectExtension(info) {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func GetPagesDirectoryTree(vo *VroomOpts) map[string]DirTree {
	utils.FindOrCreateDir(vo.PagesDirectory)
	// Go into the pages directory to find subdirectories and files
	defer os.Chdir("../")
	os.Chdir(vo.PagesDirectory)

	// Look for files in pages
	dirTreeMap := make(map[string]DirTree)
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error(err.Error())
		}
		if isCorrectExtension(info) {
			dir, fn := filepath.Split(path)
			if existing, ok := dirTreeMap[dir]; ok {
				existing.Filenames = append(existing.Filenames, fn)
				dirTreeMap[dir] = existing
			} else {
				dirTreeMap[dir] = DirTree{Filenames: []string{fn}}
			}
		}
		return nil
	})
	return dirTreeMap
}

func MakeAndParsePageTemplates(dirTreeMap map[string]DirTree,
	layouts []string, vo *VroomOpts) map[string]DirTree {
	for dir, tree := range dirTreeMap {
		var _filepaths []string
		for _, filename := range tree.Filenames {
			_filepaths = append(layouts, filepath.Join(vo.PagesDirectory, dir, filename))
		}

		temp, err := template.ParseFiles(_filepaths...)
		if err != nil {
			logger.Error(err.Error())
		}
		tree.Template = temp
		dirTreeMap[dir] = tree
	}
	return dirTreeMap
}
