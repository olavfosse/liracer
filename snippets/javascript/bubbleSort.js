function bubble_sort(array) {
	for (i = 0; i < array.length - 1; i++) {
		let sorted = true;
		for (j = 0; j < array.length - 1; j++) {
			if (array[j] > array[j + 1]) {
				let temp = array[j + 1];
				array[j + 1] = array[j];
				array[j] = temp;
				sorted = false;
			}
		}
		if (sorted) return array;
	}
	return array;
}