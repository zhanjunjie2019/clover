package confs

type Enabled uint8

func (e Enabled) Bool() bool {
	return e == 1
}
