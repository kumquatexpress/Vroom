package helpers

import (
	"github.com/kumquatexpress/Vroom/logger"
	"github.com/kumquatexpress/Vroom/utils"
	"os"
	"path/filepath"
	"text/template"
)

type FileTemplate struct {
	Filename string
	Template *template.Template
}

type DirTree struct {
	FileTemplates []FileTemplate
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
				existing.FileTemplates = append(existing.FileTemplates,
					FileTemplate{Filename:fn})
				dirTreeMap[dir] = existing
			} else {
				dirTreeMap[dir] = DirTree{FileTemplates:
					[]FileTemplate{FileTemplate{Filename:fn}}}
			}
		}
		return nil
	})
	return dirTreeMap
}

func MakeAndParsePageTemplates(dirTreeMap map[string]DirTree,
	layouts []string, vo *VroomOpts) map[string]DirTree {
	for dir, tree := range dirTreeMap {
		for idx, ft := range tree.FileTemplates {
			_filepaths := append(layouts, filepath.Join(vo.PagesDirectory, dir, ft.Filename))
			temp, err := template.ParseFiles(_filepaths...)
			if err != nil {
				logger.Error(err.Error())
			}
			ft.Template = temp
			tree.FileTemplates[idx] = ft
		}
		dirTreeMap[dir] = tree
	}
	return dirTreeMap
}
