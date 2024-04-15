package problems

func SumZero(arr []int, n int) bool {
	mp := map[int]interface{}{}
	sum := 0
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			return true
		}
		sum += arr[i]
		if _, ok := mp[sum]; ok && mp[sum] != i-1 || sum == 0 {
			return true
		} else {
			mp[sum] = i
		}
	}
	return false
}
