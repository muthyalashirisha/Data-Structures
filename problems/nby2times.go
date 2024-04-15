package problems

func Nby2Time(arr []int) int {
	cnt := 0
	ele := arr[0]

	for i := 0; i < len(arr); i++ {
		if cnt == 0 {
			ele = arr[i]
		}
		if ele == arr[i] {
			cnt++
		} else {
			cnt--
		}
	}
	return ele
}
