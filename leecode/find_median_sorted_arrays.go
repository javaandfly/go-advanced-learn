package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lens1 := len(nums1)

	lens2 := len(nums2)

	ls := (lens1 + lens2) / 2
	if lens1 == 0 && lens2 != 0 {
		if (lens1+lens2)%2 == 0 {
			return float64(nums2[ls-1]+nums2[ls]) / 2
		}

		return float64(nums2[ls])
	}

	if lens1 != 0 && lens2 == 0 {
		if (lens1+lens2)%2 == 0 {
			return float64(nums1[ls-1]+nums1[ls]) / 2
		}
		return float64(nums1[ls])
	}
	nums3 := []int{}
	l1 := lens1 - 1
	l2 := lens2 - 1

	for {
		if l1 == -1 || l2 == -1 {
			break
		}

		if nums1[l1] > nums2[l2] {
			nums3 = append(nums3, nums1[l1])
			l1--

		} else {
			nums3 = append(nums3, nums2[l2])
			l2--
		}

	}

	if l1 == -1 {
		for {
			if l2 == -1 {
				break
			}
			nums3 = append(nums3, nums2[l2])
			l2--
		}

	}
	if l2 == -1 {
		for {
			if l1 == -1 {
				break
			}
			nums3 = append(nums3, nums1[l1])
			l1--
		}

	}
	if (lens1+lens2)%2 == 0 {
		v := nums3[ls] + nums3[ls-1]
		return float64(v) / 2
	}

	return float64(nums3[ls])

}
