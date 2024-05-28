package cfg

import (
	"encoding/json"
	"github.com/jinzhu/configor"
	"io/ioutil"
	"otus-social-network-service_gen_swagger/cfg/models"
)

const ConfigFileName = "app.json"
const log_ext = ".log"

type Cfg struct {
	Env      string  `json:"env"`
	Services Service `json:"services"`
}

type Service struct {
	Service OtusSocialNetworkService `json:"otus-social-network-service"`
}

type OtusSocialNetworkService struct {
	WebServer models.WebServer `json:"webServer"`
	Log       models.Log       `json:"log"`
	Database  models.Database  `json:"database"`
	Auth      models.Auth      `json:"auth"`
}

var _config *Cfg = &Cfg{}

func Init() error {
	config := &Cfg{}
	err := configor.Load(config, ConfigFileName)
	if err != nil {
		return err
	}

	_config = config
	return nil
}

func Default() {
	configor.Load(_config)
	_config.Services.Service.Log.File = "service" + log_ext
}

func Save() {
	b, _ := json.MarshalIndent(_config, "", "  ")
	err := ioutil.WriteFile(_config.Services.Service.Log.File, b, 0x777)
	if err != nil {
		panic(err)
	}
}

func Config() *Cfg {
	return _config
}
