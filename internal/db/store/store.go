package store

import "fmt"

type Factory interface {
}

var storePointer Factory

func GetStoreFactory() (Factory, error) {
	if storePointer == nil {
		return nil, fmt.Errorf("数据层未初始化")
	}
	return storePointer, nil
}

func SetStoreFactory(factory Factory) {
	storePointer = factory
}
