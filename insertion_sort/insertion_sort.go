package main
func insertion_sort(list []int) []int {
	for i, list_item := range list {
		if i == 0 {
			continue
		}
		position := i-1
		for position >= 0 && list_item > list[position] {
			list[position+1] = list[position]
			position--
		}
		list[position+1] = list_item
	}
	return list
}
