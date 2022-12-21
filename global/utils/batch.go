package utils

import (
	"github.com/samber/lo"
	"github.com/zhanjunjie2019/clover/global/defs"
)

// ConvIDsToUint64 批量将id转为uint64类型
func ConvIDsToUint64(ids []defs.ID) []uint64 {
	return lo.Map(ids, func(item defs.ID, index int) uint64 {
		return item.UInt64()
	})
}

// LoadChangeByArrays 传入2个集合进行比较，最后对比出集合变更项，顺序不敏感。
// newArray:新结果集合;oldArray:原集合;compare:比较方法
func LoadChangeByArrays[T any](newArray, oldArray []T, compare func(newObject, oldObject *T) bool) (inserts, updates, deletes []T) {
	// 从新集合中找出新增/修改的对象
	for x := range newArray {
		newObject := newArray[x]
		var isNew = true
		for y := range oldArray {
			oldObject := oldArray[y]
			if compare(&newObject, &oldObject) {
				isNew = false
				break
			}
		}
		if isNew {
			inserts = append(inserts, newObject)
		} else {
			updates = append(updates, newObject)
		}
	}
	// 从旧集合中找出删除的对象
	for y := range oldArray {
		oldObject := oldArray[y]
		var isDel = true
		for x := range newArray {
			newObject := newArray[x]
			if compare(&newObject, &oldObject) {
				isDel = false
				break
			}
		}
		if isDel {
			deletes = append(deletes, oldObject)
		}
	}
	return
}
