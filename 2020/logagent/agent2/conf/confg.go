package conf

import "time"

// AppConf ...
type AppConf struct {
	KafkaConf KafkaConf `ini:"kafka"`
	EtcdConf  EtcdConf  `ini:"etcd"`
}

// KafkaConf ...
type KafkaConf struct {
	Address []string `ini:"address"`
}

// EtcdConf ...
type EtcdConf struct {
	Address       []string      `ini:"address"`
	Timeout       time.Duration `ini:"timeout"`
	CollectLogKey string        `ini:"collect_log_key"`
}
