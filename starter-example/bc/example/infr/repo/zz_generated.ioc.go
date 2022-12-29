//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package repo

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/starter-example/bc/example/infr/repo/po"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &exampleEntityRepo_{}
		},
	})
	exampleEntityRepoStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ExampleEntityRepo{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IRepo),
					},
				},
			},
		},
	}
	singleton.RegisterStructDescriptor(exampleEntityRepoStructDescriptor)
	allimpls.RegisterStructDescriptor(exampleEntityRepoStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &exampleValueObjectRepo_{}
		},
	})
	exampleValueObjectRepoStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ExampleValueObjectRepo{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(defs.IRepo),
					},
				},
			},
		},
	}
	singleton.RegisterStructDescriptor(exampleValueObjectRepoStructDescriptor)
	allimpls.RegisterStructDescriptor(exampleValueObjectRepoStructDescriptor)
}

type exampleEntityRepo_ struct {
	AutoMigrate_ func(ctx contextx.Context) error
	Save_        func(ctx contextx.Context, entity po.ExampleEntity) (id defs.ID, err error)
}

func (e *exampleEntityRepo_) AutoMigrate(ctx contextx.Context) error {
	return e.AutoMigrate_(ctx)
}

func (e *exampleEntityRepo_) Save(ctx contextx.Context, entity po.ExampleEntity) (id defs.ID, err error) {
	return e.Save_(ctx, entity)
}

type exampleValueObjectRepo_ struct {
	AutoMigrate_ func(ctx contextx.Context) error
	Save_        func(ctx contextx.Context, valueObject po.ExampleValueObject) (id defs.ID, err error)
}

func (e *exampleValueObjectRepo_) AutoMigrate(ctx contextx.Context) error {
	return e.AutoMigrate_(ctx)
}

func (e *exampleValueObjectRepo_) Save(ctx contextx.Context, valueObject po.ExampleValueObject) (id defs.ID, err error) {
	return e.Save_(ctx, valueObject)
}

type ExampleEntityRepoIOCInterface interface {
	AutoMigrate(ctx contextx.Context) error
	Save(ctx contextx.Context, entity po.ExampleEntity) (id defs.ID, err error)
}

type ExampleValueObjectRepoIOCInterface interface {
	AutoMigrate(ctx contextx.Context) error
	Save(ctx contextx.Context, valueObject po.ExampleValueObject) (id defs.ID, err error)
}

var _exampleEntityRepoSDID string

func GetExampleEntityRepoSingleton() (*ExampleEntityRepo, error) {
	if _exampleEntityRepoSDID == "" {
		_exampleEntityRepoSDID = util.GetSDIDByStructPtr(new(ExampleEntityRepo))
	}
	i, err := singleton.GetImpl(_exampleEntityRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ExampleEntityRepo)
	return impl, nil
}

func GetExampleEntityRepoIOCInterfaceSingleton() (ExampleEntityRepoIOCInterface, error) {
	if _exampleEntityRepoSDID == "" {
		_exampleEntityRepoSDID = util.GetSDIDByStructPtr(new(ExampleEntityRepo))
	}
	i, err := singleton.GetImplWithProxy(_exampleEntityRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ExampleEntityRepoIOCInterface)
	return impl, nil
}

type ThisExampleEntityRepo struct {
}

func (t *ThisExampleEntityRepo) This() ExampleEntityRepoIOCInterface {
	thisPtr, _ := GetExampleEntityRepoIOCInterfaceSingleton()
	return thisPtr
}

var _exampleValueObjectRepoSDID string

func GetExampleValueObjectRepoSingleton() (*ExampleValueObjectRepo, error) {
	if _exampleValueObjectRepoSDID == "" {
		_exampleValueObjectRepoSDID = util.GetSDIDByStructPtr(new(ExampleValueObjectRepo))
	}
	i, err := singleton.GetImpl(_exampleValueObjectRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ExampleValueObjectRepo)
	return impl, nil
}

func GetExampleValueObjectRepoIOCInterfaceSingleton() (ExampleValueObjectRepoIOCInterface, error) {
	if _exampleValueObjectRepoSDID == "" {
		_exampleValueObjectRepoSDID = util.GetSDIDByStructPtr(new(ExampleValueObjectRepo))
	}
	i, err := singleton.GetImplWithProxy(_exampleValueObjectRepoSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ExampleValueObjectRepoIOCInterface)
	return impl, nil
}

type ThisExampleValueObjectRepo struct {
}

func (t *ThisExampleValueObjectRepo) This() ExampleValueObjectRepoIOCInterface {
	thisPtr, _ := GetExampleValueObjectRepoIOCInterfaceSingleton()
	return thisPtr
}