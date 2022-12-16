package model

type ExampleValueObject struct {
	// 随机数据1
	RandomValue1 string
	// 随机数据2
	RandomValue2 string
	// 随机数据3
	RandomValue3 string
}

func (e ExampleValueObject) Equats(t ExampleValueObject) bool {
	if e.RandomValue1 != t.RandomValue1 {
		return false
	}
	if e.RandomValue2 != t.RandomValue2 {
		return false
	}
	if e.RandomValue3 != t.RandomValue3 {
		return false
	}
	return true
}
