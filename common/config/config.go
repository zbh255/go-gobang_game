package config

import (
	"github.com/BurntSushi/toml"
)

type AutoGenerated struct {
	Run struct {
		Ipaddr   string `toml:"ipaddr"`
		Port     string `toml:"port"`
		Mode     string `toml:"mode"`
		Database string `toml:"database"`
		Ouath    string `toml:"ouath"`
	} `toml:"run"`
	Mysql struct {
		Maxidle  int    `toml:"maxidle"`
		Maxopen  int    `toml:"maxopen"`
		Debug    bool   `toml:"debug"`
		Username string `toml:"username"`
		Dbname   string `toml:"dbname"`
		Password string `toml:"password"`
		Ipaddr   string `toml:"ipaddr"`
		Port     string `toml:"port"`
	} `toml:"mysql"`
	Sqllite struct {
		Username string `toml:"username"`
		Password string `toml:"password"`
		Filepath string `toml:"filepath"`
	} `toml:"sqllite"`
	Jwt struct {
		EncodeMethod     string `toml:"encodeMethod"`
		Key              string `toml:"key"`
		MaxEffectiveTime string `toml:"maxEffectiveTime"`
	} `toml:"jwt"`
	Redis struct {
		Ipaddr   string `toml:"ipaddr"`
		Port     string `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"redis"`
}


type ConFig interface {
	InitConfig()  *Config
}
type Config struct {
	ConfData *AutoGenerated
}
func (c *Config)InitConfig() *Config{
	var conf AutoGenerated
	capath := "/Users/harder/github.com-codes/go-gobang_game/conf/dev/app.conf.toml"
	if _, err := toml.DecodeFile(capath,&conf);err != nil {
		println("DecodeFile Toml error / ",err)
	}
	return &Config{ConfData: &conf}
}
