package ubase

import (
	model2 "github.com/uplus-io/ucluster/v1/model"
	log "github.com/uplus-io/ugo/logger"
)

type UBaseDelegate struct {
	base *UBase
}

func NewBaseDelegate(base *UBase) *UBaseDelegate {
	return &UBaseDelegate{base: base}
}

func (p *UBaseDelegate) LocalNodeStorageInfo() *model2.NodeStorageInfo {
	engine := p.base.engine
	config := p.base.config
	repository := engine.Table().Repository()
	if repository.DataCenter == 0 {
		//todo://upd cluster health
		//p.clusterHealth = proto.ClusterHealth_CH_NotInitialize
	}
	partSize := len(config.StorageConfig.Partitions)
	replicaSize := 2
	partitions := make([]*model2.Partition, partSize)
	for i := 0; i < partSize; i++ {
		part := engine.Table().PartitionOfIndex(int32(i))
		partitions[i] = &model2.Partition{Version: part.Version, Id: part.Id, Index: part.Index}
	}
	log.Infof("cluster local info loaded")
	return &model2.NodeStorageInfo{
		Repository:    &model2.Repository{DataCenter: repository.DataCenter, Area: repository.Area, Rack: repository.Rack},
		PartitionSize: int32(partSize),
		ReplicaSize:   int32(replicaSize),
		//Health:        p.getLocalNodeHealth(),
		//Status:        p.getLocalNodeStatus(),
		Partitions: partitions,
	}
	return nil
}
func (p *UBaseDelegate) LocalNodeStorageStat() *model2.NodeStorageStat {
	return nil
}
func (p *UBaseDelegate) LocalNodeHealthStat() *model2.NodeHealthStat {
	return nil
}
