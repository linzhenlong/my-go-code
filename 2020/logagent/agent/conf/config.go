package conf

import "time"

// AppConf .
type AppConf struct {
	KafkaConf `ini:"kafka"` // 与ini中的[kafka]对应
	//TailLogConf `ini:"taillog"` // 与ini中的[taillog]对应
	EtcdConf `ini:"etcd"` // 与ini中的[etcd]对应.
}

// KafkaConf .
type KafkaConf struct {
	Address []string `ini:"address"`
	//Topic   string   `ini:"topic"`
	ChanMaxSize int `ini:"chan_max_size"`
}

// EtcdConf .
type EtcdConf struct {
	Address       []string      `ini:"address"`
	TimeOut       time.Duration `ini:"timeout"`
	CollectLogkey string        `ini:"collect_log_key"`
}

// TailLogConf .
type TailLogConf struct {
	LogFile string `ini:"logfile"`
}
