package main

import "github.com/go-ini/ini"

type GlobalConfig struct{
	App string `ini:"app"`
	Database DatabaseConfig `ini:"database"`
	Server ServerConfig `ini:"server"`
}
type DatabaseConfig struct {
	Dialect string `ini:"dialect"`
	Dsn	string 	`ini:"dsn"`
	Debug bool `ini:"debug"`
}

type ServerConfig struct {
	Listen string `ini:"listen"`
	Mode  string `ini:"mode"`
}

func newConfig(filename string) *GlobalConfig{
	config := &GlobalConfig{
		App: "Hello World",
		Database:DatabaseConfig{Dialect:"sqlite3",Dsn:"cache/my.db",Debug:false},
		Server:ServerConfig{Listen:":8080",Mode:"release"},
	}
	ini.MapTo(config, "conf/my.conf")
	return config
}

func (c *GlobalConfig) Save(filename ...string){
	target:="cache/my.conf"
	if len(filename) > 0{
		target = filename[0]
	}
	cfg := ini.Empty()
	err := cfg.ReflectFrom(c)
	if err != nil{
		panic(err)
	}
	cfg.SaveTo(target)
}