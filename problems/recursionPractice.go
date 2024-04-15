package problems

import (
	"fmt"
	"sort"
)

func DistSubSeq(s string) (int, []string) {
	mp := map[string]interface{}{}
	res := []string{}
	cnt := sequences(s, 0, 0, "", &res, &mp)
	return cnt, res
}
func sequences(s string, i, j int, str string, res *[]string, mp *map[string]interface{}) int {
	if i == len(s) {
		hmap := *mp
		if hmap[str] == nil {
			*res = append(*res, str)
			hmap[str] = true
			return 1
		}
		return 0
	}
	str += string(s[i])
	pick := sequences(s, i+1, j, str, res, mp)
	str = str[0 : len(str)-1]
	notPick := sequences(s, i+1, j, str, res, mp)
	return pick + notPick
}
func BinaryString(n int) ([]string, int) {
	res := []string{}
	cnt := BinaryStringRecur(n, "", false, &res)
	return res, cnt
}
func BinaryStringRecur(n int, str string, one bool, res *[]string) int {
	if n == 0 {
		*res = append(*res, str)
		return 1
	}
	pick0 := BinaryStringRecur(n-1, "0"+str, false, res)
	pick1 := 0
	if !one {
		pick1 = BinaryStringRecur(n-1, "1"+str, true, res)
	}
	return pick1 + pick0
}
func GenerateParanthesis(n int) (int, []string) {
	result := []string{}
	cnt := GenerateParanthesisRecur("", n, n, &result)
	return cnt, result
}
func GenerateParanthesisRecur(str string, open, close int, res *[]string) int {
	if open == 0 && close == 0 {
		*res = append(*res, str)
		return 1
	}
	pickOpen := 0
	if open > 0 {
		pickOpen = GenerateParanthesisRecur(str+"(", open-1, close, res)
	}
	pickClose := 0
	if close > open {
		pickClose = GenerateParanthesisRecur(str+")", open, close-1, res)
	}
	return pickOpen + pickClose
}
func SubSeqWithSumK(arr []int, k int) (int, [][]int) {
	result := [][]int{}
	// ds := []int{}
	cnt := subSeqWithSumKRecur(arr, k, 0, []int{}, &result, 0)
	return cnt, result
}
func subSeqWithSumKRecur(arr []int, k, index int, ds []int, res *[][]int, currSum int) int {
	if index == len(arr) {
		if k == currSum {
			// fmt.Println(k, currSum, ds)
			*res = append(*res, ds)
			return 1
		}
		return 0
	}
	ds = append(ds, arr[index])
	pick := subSeqWithSumKRecur(arr, k, index+1, ds, res, arr[index]+currSum)
	ds = (ds)[:len(ds)-1]
	notPick := subSeqWithSumKRecur(arr, k, index+1, ds, res, currSum)
	return pick + notPick
}
func CombinationSum(candidates []int, target int) [][]int {
	mp := map[string]bool{}
	result := [][]int{}
	combinationUniqueSumRecur(candidates, 0, len(candidates), target, []int{}, 0, &result, &mp)
	return result
}
func combinationSumRecur(arr []int, index, n, k int, ds []int, currSum int, res *[][]int) {
	if index == n {
		if k == currSum {
			temp := make([]int, len(ds))
			copy(temp, ds)
			*res = append(*res, temp)
		}
		return
	}

	if k-currSum >= arr[index] {
		ds = append(ds, arr[index])
		combinationSumRecur(arr, index, n, k, ds, currSum+arr[index], res)
		ds = ds[:len(ds)-1]
	}
	combinationSumRecur(arr, index+1, n, k, ds, currSum, res)
}
func combinationUniqueSumRecur(arr []int, index, n, k int, ds []int, currSum int, res *[][]int, mp *map[string]bool) {
	if index == n {
		str := fmt.Sprintf("%v", ds)
		fmt.Println(str)
		if k == currSum && (*mp)[str] == false {
			temp := make([]int, len(ds))
			copy(temp, ds)
			*res = append(*res, temp)
		}
		return
	}

	if k-currSum >= arr[index] {
		ds = append(ds, arr[index])
		combinationUniqueSumRecur(arr, index+1, n, k, ds, currSum+arr[index], res, mp)
		ds = ds[:len(ds)-1]
	}
	combinationUniqueSumRecur(arr, index+1, n, k, ds, currSum, res, mp)
}

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	findComb(candidates, &result, target, []int{}, 0)
	return result
}

func findComb(arr []int, res *[][]int, k int, ds []int, index int) {
	if k == 0 {
		temp := make([]int, len(ds))
		copy(temp, ds)
		*res = append(*res, temp)
		return
	}

	for i := index; i < len(arr); i++ {
		if i > index && arr[i] == arr[i-1] {
			continue
		}
		if arr[i] > k {
			break
		}
		ds = append(ds, arr[i])
		findComb(arr, res, k-arr[i], ds, i+1)
		ds = ds[:len(ds)-1]
	}
}
func combinationSum3(k, n int) [][]int {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := [][]int{}
	findComb3(arr, &result, k, n, 0, []int{}, 0)
	return result
}
func findComb3(arr []int, res *[][]int, k, n int, currSum int, ds []int, index int) {
	if index == n {
		if k == 0 && n == currSum {
			fmt.Println(len(ds))
			temp := make([]int, len(ds))
			copy(temp, ds)
			*res = append(*res, temp)
		}
		return
	}

	ds = append(ds, arr[index])
	findComb3(arr, res, k-1, n, currSum+arr[index], ds, index+1)
	ds = ds[:len(ds)-1]
	findComb3(arr, res, k, n, currSum, ds, index+1)
}
func CombinationSum3(k int, n int) (combs [][]int) {
	helper(k, n, 1, []int{}, &combs)
	return
}

func helper(k, n, i int, cur []int, res *[][]int) {
	if k < 0 || n < 0 {
		return
	}
	if k == 0 && n == 0 {
		*res = append(*res, cur)
		return
	}
	for j := i; j <= 9; j++ {
		helper(k-1, n-j, j+1, append(cur, j), res)
	}
}
