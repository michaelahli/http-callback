package slice

func RemoveItemFromStringArray(arr []string, item int) (new_arr []string) {
	for i := 0; i < len(arr); i++ {
		if i == item {
			continue
		}
		new_arr = append(new_arr, arr[i])
	}
	return new_arr
}

func RemoveItemFrom2DStringArray(arr [][]string, item int) (new_arr [][]string) {
	for i := 0; i < len(arr); i++ {
		if i == item {
			continue
		}
		new_arr = append(new_arr, arr[i])
	}

	return new_arr
}

func RemoveItemFromIntArray(arr []int, item int) (new_arr []int) {
	for i := 0; i < len(arr); i++ {
		if i == item {
			continue
		}
		new_arr = append(new_arr, arr[i])
	}
	return new_arr
}

func RemoveElementFromStringArray(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func RemoveItemFrom2DIntArray(arr [][]int, item int) (new_arr [][]int) {
	for i := 0; i < len(arr); i++ {
		if i == item {
			continue
		}
		new_arr = append(new_arr, arr[i])
	}
	return new_arr
}
