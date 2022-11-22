package model

import "github.com/zhanjunjie2019/clover/global/defs"

func NewEntity1(id defs.ID, value Entity1Value) Entity1 {
	return &entity1{
		id:    id,
		value: value,
	}
}

type Entity1 interface {
	ID() defs.ID
	// FullValue 核心数据
	FullValue() Entity1Value
	// SetEntity2 设置其他值
	SetEntity2(entity2 Entity2)
	// GetEntity2 其他值
	GetEntity2() Entity2
}

type entity1 struct {
	id      defs.ID
	value   Entity1Value
	entity2 Entity2
}

func (e entity1) ID() defs.ID {
	return e.id
}

func (e entity1) FullValue() Entity1Value {
	return e.value
}

func (e *entity1) SetEntity2(entity2 Entity2) {
	e.entity2 = entity2
}

func (e entity1) GetEntity2() Entity2 {
	return e.entity2
}

type Entity1Value struct {
	// FirstName 姓
	FirstName string
	// LastName 名
	LastName string
}
