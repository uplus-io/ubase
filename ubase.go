package ubase

type UBaseConfig struct {
}

type UBase struct {
	config UBaseConfig
}

func NewBase(config UBaseConfig) *UBase {
	return &UBase{config: config}
}

func (p *UBase) Serving() {
}
