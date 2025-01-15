package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	fArr, sArr := nums1, nums2

	if len(fArr) > len(sArr) {
		fArr, sArr = sArr, fArr
	}

	total := len(fArr) + len(sArr)
	half := total / 2
	l, r := 0, len(fArr)-1

	for {
		fMid := (l + r + 1) / 2
		sMid := half - fMid

		fLeft, fRight, sLeft, sRight := math.Inf(-1), math.Inf(1), math.Inf(-1), math.Inf(1)

		if fMid-1 >= 0 {
			fLeft = float64(fArr[fMid-1])
		}
		if fMid < len(fArr) {
			fRight = float64(fArr[fMid])
		}
		if sMid-1 >= 0 {
			sLeft = float64(sArr[sMid-1])
		}
		if sMid < len(sArr) {
			sRight = float64(sArr[sMid])
		}

		if fLeft <= sRight && sLeft <= fRight {
			if total%2 == 0 {
				return (math.Max(fLeft, sLeft) + math.Min(fRight, sRight)) / 2
			} else {
				return math.Min(fRight, sRight)
			}
		} else if fLeft > sRight {
			r = fMid - 1
		} else {
			l = fMid + 1
		}
	}
}
