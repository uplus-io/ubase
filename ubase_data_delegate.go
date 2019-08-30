package ubase

import (
	"github.com/uplus-io/ucluster/v1"
	model2 "github.com/uplus-io/ucluster/v1/model"
)

type BaseDataDelegate struct {
	base *UBase
}

func NewBaseDataDelegate(base *UBase) *BaseDataDelegate {
	return &BaseDataDelegate{base: base}
}

func (p *BaseDataDelegate) ForEach(iterator v1.DataIterator) {
}
func (p *BaseDataDelegate) Get(data *model2.DataBody) (bool, error) {
	return false, nil
}
func (p *BaseDataDelegate) Set(data *model2.DataBody) error {
	return nil
}
