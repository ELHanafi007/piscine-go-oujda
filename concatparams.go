package piscine

func ConcatParams(args []string) string {
	res := ""             // n
	for s := range args { // args[10.20.30.40.50]
		res = res + args[s]
		if s != len(args)-1 {
			res = res + string('\n')
		}
	}
	return res
}
