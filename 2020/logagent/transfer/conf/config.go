package conf

// Conf ...
type Conf struct {
	KafkaConf KafkaConf         `ini:"kafka"`
	EsConf    ElasticSearchConf `ini:"elasticsearch"`
}

// KafkaConf ...
type KafkaConf struct {
	Address     []string `ini:"address"`
	Topic       string   `ini:"topic"`
	DialTimeout int      `ini:"dial_timeout"`
}

type ElasticSearchConf struct {
	Address       []string `ini:"address"`
	ClientTimeOut int      `ini:"client_timeout"`
}
