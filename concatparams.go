package piscine

func ConcatParams(args []string) string {
	res := ""
	for s := range args {
		res += args[s] + string('\n')
	}
	return res
}
