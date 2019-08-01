package ubase

import (
	"github.com/uplus-io/ucluster"
	"github.com/uplus-io/ucluster/model"
)

type BaseMessageDelegate struct {
	base *UBase
}

func NewBaseMessageDelegate(base *UBase) *BaseMessageDelegate {
	return &BaseMessageDelegate{base: base}
}

func (p *BaseMessageDelegate) System(pipeline ucluster.Pipeline, message model.SystemMessage) error {
	return nil
}
func (p *BaseMessageDelegate) Event(pipeline ucluster.Pipeline, message model.EventMessage) error {
	return nil
}
func (p *BaseMessageDelegate) Topic(pipeline ucluster.Pipeline, message model.TopicMessage) error {
	return nil
}
func (p *BaseMessageDelegate) Data(pipeline ucluster.Pipeline, message model.DataMessage) error {
	return nil
}
