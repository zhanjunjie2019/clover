package model

import "github.com/zhanjunjie2019/clover/global/defs"

func NewEntity2(id defs.ID, value Entity2Value) Entity2 {
	return &entity2{
		id:    id,
		value: value,
	}
}

type Entity2 interface {
	ID() defs.ID
	// FullValue 核心数据
	FullValue() Entity2Value
}

type entity2 struct {
	id    defs.ID
	value Entity2Value
}

func (e entity2) ID() defs.ID {
	return e.id
}

func (e entity2) FullValue() Entity2Value {
	return e.value
}

type Entity2Value struct {
	// 随机数据
	RandomValue string
}
