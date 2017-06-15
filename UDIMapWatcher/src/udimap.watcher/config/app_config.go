package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type AppConfiguration struct {
	// Base Setting
	ServerListenEndpoint        string `yaml:"server_listen_endpoint"`
	DBConnection                string `yaml:"db_connection"`
	MaxIdleConns                int    `yaml:"max_idle_conns"`
	MaxOpenConns                int    `yaml:"max_open_conns"`
	ConnMaxLifetime             int    `yaml:"conn_max_life_time"`
	IntervalTime                int    `yaml:"interval_time"`
	UDIMapEndpoint              string `yaml:"udi_map_endpoint"`
	UDIMapKey                   string `yaml:"udi_map_key"`
	GetCameraSoapAction         string `yaml:"get_camera_soap_action"`
	GetTramDoMuaSoapAction      string `yaml:"get_tram_do_mua_soap_action"`
	GetTramDoTrieuSoapAction    string `yaml:"get_tram_do_trieu_soap_action"`
	GetDSQuanTracNgapSoapAction string `yaml:"get_ds_quan_trac_ngap_soap_action"`
	GetDSDiemNgapSoapAction     string `yaml:"get_ds_diem_ngap_soap_action"`
}

// AppConfig is global instance of AppConfiguration
var AppConfig *AppConfiguration

// InitFromYAML reads config from a configuration.yaml file and set value into AppConfig instance
func InitFromYAML() bool {
	filename := os.Args[0] + ".yaml"
	fmt.Println("[UDIMapWatcher] Start loading configurations from", filename)
	AppConfigAbsPath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("[UDIMapWatcher] Loading configurations fail", err)
		return false
	}

	// read the raw contents of the file
	data, err := ioutil.ReadFile(AppConfigAbsPath)
	if err != nil {
		fmt.Println("[UDIMapWatcher] Read file error", err)
		return false
	}

	c := AppConfiguration{}
	// put the file's contents as yaml to the default configuration(c)
	if err := yaml.Unmarshal(data, &c); err != nil {
		fmt.Println("[UDIMapWatcher] Read file error", err)
		return false
	}

	AppConfig = &c

	fmt.Println("[UDIMapWatcher] Load configurations successful.")
	return true
}
