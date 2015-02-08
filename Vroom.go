package main

import (
	"fmt"
	"github.com/kumquatexpress/Vroom/helpers"
	"text/template"
	_ "io/ioutil"
	"os"
	"github.com/kumquatexpress/Vroom/logger"
	"github.com/kumquatexpress/Vroom/utils"
	"path/filepath"
)

type DirTree struct {
	Template *template.Template
	Filenames []string
}

func getDirectoryTree(vo *helpers.VroomOpts) map[string]DirTree{
	// Check for existing template directory or create new
	if !utils.Exists(vo.TemplateDirectory) {
		logger.Log(fmt.Sprintf("No template directory found, creating %s", vo.TemplateDirectory))
		os.MkdirAll(vo.TemplateDirectory, os.ModePerm)
	}

	// Go into the templates directory to find subdirectories and files
	defer os.Chdir("../")
	os.Chdir(vo.TemplateDirectory)

	// Look for files in templates
	dirTreeMap := make(map[string]DirTree)
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error(err.Error())
		}
		if !info.IsDir() {
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
	return makeAndParseTemplates(dirTreeMap)
}

func makeAndParseTemplates(dirTreeMap map[string]DirTree) map[string]DirTree {
	for dir, tree := range dirTreeMap {
		var _filepaths []string
		for _, filename := range tree.Filenames {
			_filepaths = append(_filepaths, filepath.Join(dir, filename))
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

func main() {
	// Create the options
	vo := helpers.NewVroomOpts()
	logger.Log(fmt.Sprintf("Options: %+v", vo))

	templateMap := getDirectoryTree(vo)
	logger.Log(fmt.Sprintf("Templates: %+v", templateMap))
	if len(templateMap) < 1 {
		logger.Log(fmt.Sprintf("No templates found in %s, exiting.", vo.TemplateDirectory))
		return
	}

	// Replace existing build directory if needed
	os.RemoveAll(vo.BuildDirectory)
	// Build and create new files
	var _filepath string
	var _builddir string
	for dir, temp := range templateMap {
		for _, fn := range temp.Filenames {
			_builddir = filepath.Join(vo.BuildDirectory, dir)
			_filepath = filepath.Join(_builddir, fn)

			os.MkdirAll(_builddir, os.ModePerm)
			file, err := os.Create(_filepath)
			if err != nil {
				logger.Error(err.Error())
			}
			temp.Template.ExecuteTemplate(file, fn, nil)
			logger.Log(fmt.Sprintf("File %s built", _filepath))
		}
		logger.Log(fmt.Sprintf("Directory %s built", dir))
	}
}
