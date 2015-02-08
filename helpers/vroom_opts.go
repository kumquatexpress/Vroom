package helpers

import (
	"encoding/json"
	"github.com/kumquatexpress/Vroom/logger"
	"github.com/kumquatexpress/Vroom/utils"
	"io/ioutil"
	_ "text/template"
)

const (
	TEMPLATE_DIR = "templates"
	COMPILE_DIR  = "build"
)

type VroomOpts struct {
	TemplateDirectory string
	BuildDirectory  string
	Metadata          map[string]string
}

func NewVroomOpts() *VroomOpts {
	return &VroomOpts{
		TemplateDirectory: TEMPLATE_DIR,
		BuildDirectory:  COMPILE_DIR,
	}
}

func parseOpts(data []byte) (*VroomOpts, error) {
	var vo VroomOpts
	err := json.Unmarshal(data, &vo)
	if err != nil {
		return nil, err
	}
	return &vo, nil
}

func NewVroomOptsFromFile(filename string) *VroomOpts {
	if !utils.Exists(filename) {
		logger.Log("No options file found, using default options.")
		return NewVroomOpts()
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Warn(err.Error())
		return NewVroomOpts()
	}
	opts, err := parseOpts(data)
	if err != nil {
		logger.Warn(err.Error())
		return NewVroomOpts()
	}
	return opts
}
