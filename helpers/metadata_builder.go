package helpers

import (
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"github.com/kumquatexpress/Vroom/logger"
	"github.com/kumquatexpress/Vroom/utils"
)

func GetTreeWithMetadata(tree map[string]DirTree, vo *VroomOpts) map[string]DirTree {
	for dir, t := range tree {
		t.Data = buildMetadata(filepath.Join(vo.PagesDirectory, dir), make(map[string]interface{}))
		tree[dir] = t
	}
	return tree
}

func buildMetadata(path string, accumulator map[string]interface{}) map[string]interface{} {
	parent := filepath.Dir(path)
	data := utils.MergeMap(extractDataFromDirectory(path), accumulator)
	if parent == "." { // empty path, base case
		return data
	}
	return buildMetadata(parent, data)
}

func extractDataFromDirectory(path string) map[string]interface{} {
	var data interface{}
	files, err := filepath.Glob(filepath.Join(path, "*.vroom.json"))
	if err != nil {
		logger.Warn(err.Error())
		return make(map[string]interface{})
	}
	for _, f := range files {
		buf, err := ioutil.ReadFile(f)
		if err != nil {
			logger.Warn(err.Error())
		} else {
			err = json.Unmarshal(buf, &data)
			if err != nil {
				logger.Warn(err.Error())
			}
		}
	}
	if data != nil {
		return data.(map[string]interface{})
	}
	return make(map[string]interface{})
}