def linear_search(array, target_value):
	for i in range(0, len(array)):
		if array[i] == target_value:
			return i+1