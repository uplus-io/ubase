package ubase

import (
	"github.com/uplus-io/ucluster"
	"github.com/uplus-io/ucluster/model"
)

type BaseDataDelegate struct {
	base *UBase
}

func NewBaseDataDelegate(base *UBase) *BaseDataDelegate {
	return &BaseDataDelegate{base: base}
}

func (p *BaseDataDelegate) ForEach(iterator ucluster.DataIterator) {
}
func (p *BaseDataDelegate) Get(data *model.DataBody) (bool, error) {
	return false, nil
}
func (p *BaseDataDelegate) Set(data *model.DataBody) error {
	return nil
}
