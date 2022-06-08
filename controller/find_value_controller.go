package controller

import (
	"fidtest_golang/entity"
	"fidtest_golang/model"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

func FindValue(w http.ResponseWriter, r *http.Request) {
	dataset := make(map[int]string)
	datatest := []string{"1", "X", "8", "17", "Y", "Z", "78", "113"}
	for i, val := range datatest {
		dataset[i] = val
	}

	index_dataset := make([]int, 0, len(dataset))
	var test []int
	var value []int
	var index_x int
	var value_x int
	var value_y int
	var value_z int
	var step_up int
	index0, _ := strconv.Atoi(dataset[0])
	// var exact_num int
	_ = index_x
	for k := range dataset {
		index_dataset = append(index_dataset, k)
	}

	// sort the slice by keys
	sort.Ints(index_dataset)

	// iterate by sorted keys
	count := 0
	for _, index := range index_dataset {
		if _, err := strconv.Atoi(dataset[index]); err == nil {
			count++
			if count > 0 {
				test = append(test, index)
			}
		} else {
			if count == 2 {
				count = 0
				break
			}
			if dataset[index] == "X" {
				index_x = count
			}
			count--
			test = removeIndex(test, count)
		}
	}

	for i, val := range test {
		_ = val
		a, _ := strconv.Atoi(dataset[test[i]])
		value = append(value, a)
	}

	fid_base1 := value[1] - value[0]
	fmt.Println("fid_base1 : ", fid_base1)
	for i := 2; i < value[0]; i++ {
		value = append([]int{i}, value...)
		fid_base2 := value[1] - value[0]
		exact_num := 1
		if fid_base2-value[0] > 1 && exact_num == 1 {
			if (value[0] - index0) > index0 {
				value_x = value[0]
				fmt.Println("value_x : ", value_x)
				step_up = (value[2] - value[1] - fid_base2) + exact_num
				value_y = step_up + fid_base1 + value[2]
				fmt.Println("value_y : ", value_y)
				step_up = step_up + exact_num
				aa := value_y - value[2]
				value_z = step_up + value_y + aa
				fmt.Println("value_z : ", value_z)
				break
			}
		}
		value = removeIndex(value, 0)
	}
	payload := entity.FindValueInfo{
		X: value_x,
		Y: value_y,
		Z: value_z,
	}
	model.ModelResponseData(w,http.StatusOK, http.StatusText(http.StatusOK), payload)

}
func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}