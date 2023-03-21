package utils

// AllAaccordance 检查所有给定的字符串是否等于它们对应的字符串。
// 如果所有字符串都相等，则返回true，否则返回false。
func AllAaccordance(args ...string) bool {
	if len(args)%2 != 0 {
		return false
	}
	for i := 0; i < len(args); i += 2 {
		if args[i] != args[i+1] {
			return false
		}
	}
	return true
}
