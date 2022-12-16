package utils

import "encoding/json"

// CopyObject 通过json序列化方式，进行对象复制。一定要是指针
func CopyObject(src, target any) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, target)
}
