function linear_search(array, target) {
	for (i = 0; i < array.length; i++) {
		if (array[i] == target) {
			return i;
		}
	}
}
