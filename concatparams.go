package piscine

func ConcatParams(args []string) string {
	res := ""
	for s := range args {
		res += args[s]
		if s != len(args)-1 {
			res += string('\n')
		}
	}
	return res
}
