package ubase

import (
	"github.com/uplus-io/ucluster/v1"
	"github.com/uplus-io/uengine"
)

type UBaseConfig struct {
	Cluster       v1.UClusterConfig `json:"cluster" yaml:"cluster"`
	StorageConfig StorageConfig     `json:"storage" yaml:"storage"`
	LogConfig     LogConfig         `json:"log" yaml:"log"`
}

type StorageConfig struct {
	Engine     string   `json:"engine" yaml:"engine"`
	Meta       string   `json:"meta" yaml:"meta"`
	Wal        string   `json:"wal" yaml:"wal"`
	Partitions []string `json:"partitions" yaml:"partitions"`
}

type LogConfig struct {
	Path     string `json:"path" yaml:"path"`
	Filename string `json:"filename" yaml:"filename"`
	Level    uint32 `json:"level" yaml:"level"`
}

type UBase struct {
	config  *UBaseConfig
	cluster *v1.UCluster
	engine  *uengine.Engine
}

func NewBase(config *UBaseConfig) *UBase {
	engine := startEngine(config)
	base := &UBase{
		config:  config,
		engine:  engine,
		cluster: v1.NewUCluster(&config.Cluster),
	}
	configure(config, base)
	return base
}

func configure(config *UBaseConfig, base *UBase) {
	config.Cluster.Delegate = NewBaseDelegate(base)
	config.Cluster.MessageDelegate = NewBaseMessageDelegate(base)
	config.Cluster.DataDelegate = NewBaseDataDelegate(base)
}

func startEngine(config *UBaseConfig) *uengine.Engine {
	storageConfig := config.StorageConfig
	engineConfig := uengine.UEngineConfig{
		Engine:     storageConfig.Engine,
		Wal:        storageConfig.Wal,
		Meta:       storageConfig.Meta,
		Partitions: storageConfig.Partitions,
	}
	engine := uengine.NewEngine(engineConfig)
	return engine
}

func (p *UBase) Serving() {
	err := p.cluster.Serving()
	if err != nil {
		panic(err)
	}
}
