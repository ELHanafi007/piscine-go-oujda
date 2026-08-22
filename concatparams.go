package piscine

func ConcatParams(args []string) string {
	res := ""
	for s := range args {
		if s != len(args)-1 {
			res += args[s] + string('\n')
		}
	}
	return res
}
