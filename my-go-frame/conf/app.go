package conf

// APPconf ...
type APPconf struct {
	Domain       string `ini:"domain"`
	Addr         string `ini:"listen"`
	ReadTimeout  int    `ini:"read_timeout"`  // 请求超时时间
	WriteTimeout int    `ini:"write_timeout"` // 请求超时时间
}

func defaultAppConfig() *APPconf {
	return &APPconf{
		Domain:       "lzl-lo",
		ReadTimeout:  5000,
		WriteTimeout: 5000,
	}
}
