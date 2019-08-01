package ubase

import (
	"github.com/uplus-io/ucluster/model"
)

type UBaseDelegate struct {
	base *UBase
}

func NewBaseDelegate(base *UBase) *UBaseDelegate {
	return &UBaseDelegate{base: base}
}

func (p *UBaseDelegate) LocalNodeStorageInfo() *model.NodeStorageInfo {
	return nil
}
func (p *UBaseDelegate) LocalNodeStorageStat() *model.NodeStorageStat {
	return nil
}
func (p *UBaseDelegate) LocalNodeHealthStat() *model.NodeHealthStat {
	return nil
}
