package main

func foreach(list []int, transform_function func(uint, int) int) []int{
	var transformed_result []int;
	for index, item := range list {
		transformed_result = append(transformed_result, transform_function(uint(index), item))
	}
	return transformed_result;
}