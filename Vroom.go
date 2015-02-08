package main

import (
	"fmt"
	"github.com/kumquatexpress/Vroom/helpers"
	"github.com/kumquatexpress/Vroom/utils"
	"github.com/kumquatexpress/Vroom/logger"
	"os"
	"path/filepath"
)

func generateAndWriteFiles(pagesMap map[string]helpers.DirTree,
	vo *helpers.VroomOpts) {
	// Build and create new files
	var _filepath string
	var _builddir string
	for dir, tree := range pagesMap {
		dirData := utils.MergeMap(vo.Metadata, tree.Data)

		for _, ft := range tree.FileTemplates {
			_builddir = filepath.Join(vo.BuildDirectory, dir)
			_filepath = filepath.Join(_builddir, ft.Filename)

			os.MkdirAll(_builddir, os.ModePerm)
			file, err := os.Create(_filepath)
			if err != nil {
				logger.Error(err.Error())
			}
			ft.Template.ExecuteTemplate(file, ft.Filename, dirData)
			logger.Log(fmt.Sprintf("File %s built", _filepath))
		}
		logger.Log(fmt.Sprintf("Directory %s built with data %+v", _builddir, dirData))
	}
}

func main() {
	// Create the options
	vo := helpers.NewVroomOptsFromFile("config.vroom.json")
	logger.Log(fmt.Sprintf("Options: %+v", vo))

	pagesMap := helpers.GetPagesDirectoryTree(vo)
	layoutFiles := helpers.FindLayoutFiles(vo)
	pagesMap = helpers.MakeAndParsePageTemplates(pagesMap, layoutFiles, vo)
	if len(pagesMap) < 1 {
		logger.Log(fmt.Sprintf("Found no pages to render in %s, exiting.", vo.PagesDirectory))
		return
	}
	pagesMap = helpers.GetTreeWithMetadata(pagesMap, vo)

	// Replace existing build directory if needed, then generate
	os.RemoveAll(vo.BuildDirectory)
	generateAndWriteFiles(pagesMap, vo)
}
