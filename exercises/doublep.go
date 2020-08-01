package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func threeSum(nums []int, target int) [][]int {
	var result [][]int  //// 记录返回结果
	length := len(nums) //// 获取数组长度
	if nums == nil || length < 3 {
		return result
	}
	sort.Ints(nums) //// 排序
	for i := 0; i < length; i++ {
		//// 如果当前数字大于0，则三数之和一定大于0，所以结束循环
		if nums[i] > target {
			break
		}
		//// 去重，下标 i 的数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := length - 1
		for {
			if l >= r {
				break
			}
			sum := nums[i] + nums[l] + nums[r]
			if sum > target {
				r--
				continue
			}
			if sum < target {
				l++
				continue
			}
			if sum == target {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				//// 去重，下标 l
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				//// 去重，下标 r
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return result
}

func GetArr(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

func main() {
	target := 20
	aa := GetArr(100)
	fmt.Println(threeSum(aa, target))
}
