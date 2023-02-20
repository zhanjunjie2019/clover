package utils

func AllAaccordance(args ...string) bool {
	for i := 0; i < len(args)/2; i += 2 {
		if args[i] != args[i+1] {
			return false
		}
	}
	return true
}
