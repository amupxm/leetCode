package containerwiththemostwater

func maxArea(height []int) int {
	var left, right int
	left = 0
	right = len(height) - 1

	maxAmmountOfWater := 0

	for left != right {
		ammountOfWater := calculateAmount(height[left], height[right], left, right)
		if ammountOfWater > maxAmmountOfWater {
			maxAmmountOfWater = ammountOfWater
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxAmmountOfWater
}

func maxOf(int1, int2 int) int {
	if int1 > int2 {
		return int1
	} else {
		return int2
	}
}

func minOf(int1, int2 int) int {
	if int1 < int2 {
		return int1
	} else {
		return int2
	}
}
func calculateAmount(left, right, leftIndex, rightIndex int) int {
	return (rightIndex - leftIndex) * minOf(left, right)
}
