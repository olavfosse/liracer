def bubble_sort(array):
	for _ in range(len(array)-1):
		sorted = True
		for i in range(len(array)-1):
			if array[i] > array[i+1]:
				temp = array[i+1]
				array[i+1] = array[i]
				array[i] = temp
				sorted = False
		if sorted:
			return array

	return array