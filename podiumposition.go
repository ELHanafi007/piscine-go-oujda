package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium)/2; i++ {
		tmp := podium[i]
		podium[i] = podium[len(podium)-1-i]
		podium[len(podium)-1-i] = tmp
	}
	return podium
}
