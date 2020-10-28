package conf

import "flag"

import "github.com/go-ini/ini"

var (
	configPath *string
	// FrameConf 框架配置...
	FrameConf = &FrameConfig{}
)

// FrameConfig 框架配置.
type FrameConfig struct {
	Debug   int      `ini:"debug"`
	AppConf *APPconf `ini:"app"`
}

func defaultConfig() *FrameConfig {
	return &FrameConfig{
		Debug:   0,
		AppConf: defaultAppConfig(),
	}
}

// InitConf 初始配置文件.
func InitConf() error {
	defaultConfig()
	err := ini.MapTo(FrameConf, *configPath)
	return err
}

func init() {
	configPath = flag.String("conf", "./config.ini", "加载配置文件")
}
