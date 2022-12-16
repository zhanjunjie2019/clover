package model

import "github.com/zhanjunjie2019/clover/global/defs"

func NewExampleEntity(id defs.ID, value ExampleEntityValue) ExampleEntity {
	return &exampleEntity{
		id:    id,
		value: value,
	}
}

type ExampleEntity interface {
	ID() defs.ID
	// FullValue 核心数据
	FullValue() ExampleEntityValue
	// SetValueObject 设置其他值对象
	SetValueObject(exampleValueObject ExampleValueObject)
	// GetValueObject 其他值对象
	GetValueObject() ExampleValueObject
}

type exampleEntity struct {
	id                 defs.ID
	value              ExampleEntityValue
	exampleValueObject ExampleValueObject
}

func (e exampleEntity) ID() defs.ID {
	return e.id
}

func (e exampleEntity) FullValue() ExampleEntityValue {
	return e.value
}

func (e *exampleEntity) SetValueObject(exampleValueObject ExampleValueObject) {
	e.exampleValueObject = exampleValueObject
}

func (e exampleEntity) GetValueObject() ExampleValueObject {
	return e.exampleValueObject
}

type ExampleEntityValue struct {
	// FirstName 姓
	FirstName string
	// LastName 名
	LastName string
}
