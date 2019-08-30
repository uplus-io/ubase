package ubase

import (
	"github.com/uplus-io/ucluster/v1"
	model2 "github.com/uplus-io/ucluster/v1/model"
)

type BaseMessageDelegate struct {
	base *UBase
}

func NewBaseMessageDelegate(base *UBase) *BaseMessageDelegate {
	return &BaseMessageDelegate{base: base}
}

func (p *BaseMessageDelegate) System(pipeline v1.Pipeline, message model2.SystemMessage) error {
	return nil
}
func (p *BaseMessageDelegate) Event(pipeline v1.Pipeline, message model2.EventMessage) error {
	return nil
}
func (p *BaseMessageDelegate) Topic(pipeline v1.Pipeline, message model2.TopicMessage) error {
	return nil
}
func (p *BaseMessageDelegate) Data(pipeline v1.Pipeline, message model2.DataMessage) error {
	return nil
}
